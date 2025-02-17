package v1alpha1_test

import (
	"path/filepath"
	"strings"

	envoy_resource "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	core_plugins "github.com/kumahq/kuma/pkg/core/plugins"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
	"github.com/kumahq/kuma/pkg/plugins/policies/meshloadbalancingstrategy/api/v1alpha1"
	plugin "github.com/kumahq/kuma/pkg/plugins/policies/meshloadbalancingstrategy/plugin/v1alpha1"
	policies_xds "github.com/kumahq/kuma/pkg/plugins/policies/xds"
	"github.com/kumahq/kuma/pkg/test/matchers"
	"github.com/kumahq/kuma/pkg/test/resources/builders"
	"github.com/kumahq/kuma/pkg/test/resources/samples"
	"github.com/kumahq/kuma/pkg/util/pointer"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/envoy/clusters"
	. "github.com/kumahq/kuma/pkg/xds/envoy/listeners"
	"github.com/kumahq/kuma/pkg/xds/generator"
)

func getResource(resourceSet *core_xds.ResourceSet, typ envoy_resource.Type) []byte {
	resources, err := resourceSet.ListOf(typ).ToDeltaDiscoveryResponse()
	Expect(err).ToNot(HaveOccurred())
	actual, err := util_proto.ToYAML(resources)
	Expect(err).ToNot(HaveOccurred())

	return actual
}

var _ = Describe("MeshLoadBalancingStrategy", func() {
	type testCase struct {
		resources []core_xds.Resource
		toRules   core_xds.ToRules
	}
	DescribeTable("Apply",
		func(given testCase) {
			resources := core_xds.NewResourceSet()
			for _, res := range given.resources {
				r := res
				resources.Add(&r)
			}

			context := xds_context.Context{}
			proxy := core_xds.Proxy{
				APIVersion: envoy_common.APIV3,
				Dataplane: samples.DataplaneBackendBuilder().
					AddOutbound(
						builders.Outbound().WithAddress("127.0.0.1").WithPort(27777).WithTags(map[string]string{
							mesh_proto.ServiceTag:  "backend",
							mesh_proto.ProtocolTag: "http",
						}),
					).
					AddOutbound(
						builders.Outbound().WithAddress("127.0.0.1").WithPort(27778).WithTags(map[string]string{
							mesh_proto.ServiceTag:  "payment",
							mesh_proto.ProtocolTag: "http",
						}),
					).
					Build(),
				Policies: core_xds.MatchedPolicies{
					Dynamic: map[core_model.ResourceType]core_xds.TypedMatchingPolicies{
						v1alpha1.MeshLoadBalancingStrategyType: {
							Type:    v1alpha1.MeshLoadBalancingStrategyType,
							ToRules: given.toRules,
						},
					},
				},
				Routing: core_xds.Routing{
					OutboundTargets: map[core_xds.ServiceName][]core_xds.Endpoint{
						"backend": {
							{
								Tags: map[string]string{mesh_proto.ProtocolTag: core_mesh.ProtocolHTTP},
							},
						},
						"payment": {
							{
								Tags: map[string]string{mesh_proto.ProtocolTag: core_mesh.ProtocolHTTP},
							},
						},
					},
				},
			}
			plugin := plugin.NewPlugin().(core_plugins.PolicyPlugin)
			Expect(plugin.Apply(resources, context, &proxy)).To(Succeed())

			nameSplit := strings.Split(GinkgoT().Name(), " ")
			name := nameSplit[len(nameSplit)-1]

			Expect(getResource(resources, envoy_resource.ListenerType)).To(matchers.MatchGoldenYAML(filepath.Join("testdata", name+".listeners.golden.yaml")))
			Expect(getResource(resources, envoy_resource.ClusterType)).To(matchers.MatchGoldenYAML(filepath.Join("testdata", name+".clusters.golden.yaml")))
		},
		Entry("basic", testCase{
			resources: []core_xds.Resource{
				{
					Name:   "cluster-backend",
					Origin: generator.OriginOutbound,
					Resource: clusters.NewClusterBuilder(envoy_common.APIV3).
						Configure(policies_xds.WithName("backend")).
						MustBuild(),
				},
				{
					Name:   "cluster-payment",
					Origin: generator.OriginOutbound,
					Resource: clusters.NewClusterBuilder(envoy_common.APIV3).
						Configure(policies_xds.WithName("payment")).
						MustBuild(),
				},
				{
					Name:   "listener-backend",
					Origin: generator.OriginOutbound,
					Resource: NewListenerBuilder(envoy_common.APIV3).
						Configure(OutboundListener("outbound:127.0.0.1:27777", "127.0.0.1", 27777, core_xds.SocketAddressProtocolTCP)).
						Configure(FilterChain(NewFilterChainBuilder(envoy_common.APIV3).
							Configure(HttpConnectionManager("127.0.0.1:27777", false)).
							Configure(
								HttpOutboundRoute(
									"backend",
									envoy_common.Routes{{
										Clusters: []envoy_common.Cluster{envoy_common.NewCluster(
											envoy_common.WithService("backend"),
											envoy_common.WithWeight(100),
										)},
									}},
									map[string]map[string]bool{
										"kuma.io/service": {
											"backend": true,
										},
									},
								),
							),
						)).MustBuild(),
				},
				{
					Name:   "listener-payments",
					Origin: generator.OriginOutbound,
					Resource: NewListenerBuilder(envoy_common.APIV3).
						Configure(OutboundListener("outbound:127.0.0.1:27778", "127.0.0.1", 27778, core_xds.SocketAddressProtocolTCP)).
						Configure(FilterChain(NewFilterChainBuilder(envoy_common.APIV3).
							Configure(HttpConnectionManager("127.0.0.1:27778", false)).
							Configure(
								HttpOutboundRoute(
									"backend",
									envoy_common.Routes{{
										Clusters: []envoy_common.Cluster{envoy_common.NewCluster(
											envoy_common.WithService("payment"),
											envoy_common.WithWeight(100),
										)},
									}},
									map[string]map[string]bool{
										"kuma.io/service": {
											"payment": true,
										},
									},
								),
							),
						)).MustBuild(),
				},
			},
			toRules: core_xds.ToRules{
				Rules: []*core_xds.Rule{
					{
						Subset: core_xds.MeshService("backend"),
						Conf: v1alpha1.Conf{
							LoadBalancer: &v1alpha1.LoadBalancer{
								Type: v1alpha1.RandomType,
							},
						},
					},
					{
						Subset: core_xds.MeshService("payment"),
						Conf: v1alpha1.Conf{
							LoadBalancer: &v1alpha1.LoadBalancer{
								Type: v1alpha1.RingHashType,
								RingHash: &v1alpha1.RingHash{
									MinRingSize:  pointer.To[uint32](100),
									MaxRingSize:  pointer.To[uint32](1000),
									HashFunction: pointer.To(v1alpha1.MurmurHash2Type),
									HashPolicies: &[]v1alpha1.HashPolicy{
										{
											Type: v1alpha1.QueryParameterType,
											QueryParameter: &v1alpha1.QueryParameter{
												Name: "queryparam",
											},
											Terminal: pointer.To(true),
										},
										{
											Type: v1alpha1.ConnectionType,
											Connection: &v1alpha1.Connection{
												SourceIP: pointer.To(true),
											},
											Terminal: pointer.To(false),
										},
									},
								},
							},
						},
					},
				},
			},
		}),
	)
})
