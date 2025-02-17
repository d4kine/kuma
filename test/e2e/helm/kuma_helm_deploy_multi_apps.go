package helm

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/pkg/config/core"
	. "github.com/kumahq/kuma/test/framework"
	"github.com/kumahq/kuma/test/framework/client"
	"github.com/kumahq/kuma/test/framework/deployments/democlient"
	"github.com/kumahq/kuma/test/framework/deployments/testserver"
)

func AppDeploymentWithHelmChart() {
	var cluster Cluster
	var skip bool

	E2EAfterEach(func() {
		if !skip {
			Expect(cluster.DeleteNamespace(TestNamespace)).To(Succeed())
			Expect(cluster.DeleteKuma()).To(Succeed())
		}

		Expect(cluster.DismissCluster()).To(Succeed())
	})

	DescribeTable(
		"Should deploy two apps",
		func(cniVersion CNIVersion) {
			cluster = NewK8sCluster(NewTestingT(), Kuma1, Silent).
				WithTimeout(6 * time.Second).
				WithRetries(60)

			annotations := map[string]string{}
			if cniVersion == CNIVersion1 {
				if c, msg := cluster.(*K8sCluster).K8sVersionCompare("1.22.0", "k8s cluster doesn't support legacy CNI"); c >= 0 {
					// k3s from version 1.22 comes with flannel CNI plugin in version 1,
					// which is not supported with our default/legacy kuma-cni plugin
					// (max supported version is 0.4)
					Skip(msg)
				}
				annotations["kuma.io/builtindns"] = "enabled"
				annotations["kuma.io/builtindnsport"] = "15053"
			}

			minReplicas := 3
			err := NewClusterSetup().
				Install(Kuma(core.Standalone,
					WithInstallationMode(HelmInstallationMode),
					WithHelmReleaseName(fmt.Sprintf("kuma-%s", strings.ToLower(random.UniqueId()))),
					WithSkipDefaultMesh(true), // it's common case for HELM deployments that Mesh is also managed by HELM therefore it's not created by default
					WithHelmOpt("controlPlane.autoscaling.enabled", "true"),
					WithHelmOpt("controlPlane.autoscaling.minReplicas", strconv.Itoa(minReplicas)),
					WithCNI(cniVersion),
				)).
				Install(MeshKubernetes("default")).
				Install(NamespaceWithSidecarInjection(TestNamespace)).
				Install(democlient.Install(democlient.WithNamespace(TestNamespace), democlient.WithPodAnnotations(annotations))).
				Install(testserver.Install(testserver.WithPodAnnotations(annotations))).
				Setup(cluster)
			Expect(err).ToNot(HaveOccurred())

			Expect(cluster.(*K8sCluster).WaitApp(Config.KumaServiceName, Config.KumaNamespace, minReplicas)).To(Succeed())

			Eventually(func(g Gomega) {
				_, err := client.CollectEchoResponse(
					cluster, "demo-client", "test-server",
					client.FromKubernetesPod(TestNamespace, "demo-client"),
				)
				g.Expect(err).ToNot(HaveOccurred())
			}, "30s", "1s").Should(Succeed())

			Eventually(func(g Gomega) {
				_, err := client.CollectEchoResponse(
					cluster, "demo-client", "test-server_kuma-test_svc_80.mesh",
					client.FromKubernetesPod(TestNamespace, "demo-client"),
				)
				g.Expect(err).ToNot(HaveOccurred())
			}, "30s", "1s").Should(Succeed())

			Eventually(func(g Gomega) {
				_, err := client.CollectEchoResponse(
					cluster, "demo-client", "test-server.kuma-test.svc.80.mesh",
					client.FromKubernetesPod(TestNamespace, "demo-client"),
				)
				g.Expect(err).ToNot(HaveOccurred())
			}, "30s", "1s").Should(Succeed())
		},
		Entry("with cni v1 (legacy)", CNIVersion1),
		Entry("with cni v2 (default)", CNIVersion2),
	)
}
