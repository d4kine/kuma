package postgres

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	postgres_config "github.com/kumahq/kuma/pkg/config/plugins/resources/postgres"
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/plugins"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/core/runtime/component"
	kuma_events "github.com/kumahq/kuma/pkg/events"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	postgres_events "github.com/kumahq/kuma/pkg/plugins/resources/postgres/events"
	test_postgres "github.com/kumahq/kuma/pkg/test/store/postgres"
	"github.com/kumahq/kuma/pkg/util/channels"
)

var _ = Describe("Events", func() {
	var cfg postgres_config.PostgresStoreConfig

	BeforeEach(func() {
		c, err := c.Config(test_postgres.WithRandomDb)
		Expect(err).ToNot(HaveOccurred())
		cfg = *c
		ver, err := MigrateDb(cfg)
		Expect(err).ToNot(HaveOccurred())
		Expect(ver).To(Equal(plugins.DbVersion(1677096751)))
	})

	DescribeTable("should receive a notification from pq listener",
		func(driverName string) {
			// given
			listenerStopCh, listenerErrCh, eventBusStopCh, storeErrCh := setupChannels()
			defer close(eventBusStopCh)
			defer close(listenerErrCh)
			listener := setupListeners(cfg, driverName, listenerErrCh, listenerStopCh)
			go triggerNotifications(cfg, storeErrCh)

			// when
			event, err := listener.Recv(eventBusStopCh)

			// then
			Expect(err).To(Not(HaveOccurred()))
			resourceChanged := event.(kuma_events.ResourceChangedEvent)
			Expect(resourceChanged.Operation).To(Equal(kuma_events.Create))
			Expect(resourceChanged.Type).To(Equal(model.ResourceType("Mesh")))

			// and shutdown gracefully
			close(listenerStopCh)
			close(storeErrCh)
			Eventually(channelClosesWithoutErrors(listenerErrCh), "5s", "10ms").Should(BeTrue())
			Eventually(channelClosesWithoutErrors(storeErrCh), "5s", "10ms").Should(BeTrue())
		},
		Entry("When using pq", postgres_config.DriverNamePq),
		Entry("When using pgx", postgres_config.DriverNamePgx),
	)

	DescribeTable("should continue handling notification after postgres recovery",
		func(driverName string) {
			// given
			listenerStopCh, listenerErrCh, eventBusStopCh, storeErrCh := setupChannels()
			defer close(eventBusStopCh)
			listener := setupListeners(cfg, driverName, listenerErrCh, listenerStopCh)
			go triggerNotifications(cfg, storeErrCh)

			// when
			event, err := listener.Recv(eventBusStopCh)

			// then
			Expect(err).To(Not(HaveOccurred()))
			resourceChanged := event.(kuma_events.ResourceChangedEvent)
			Expect(resourceChanged.Operation).To(Equal(kuma_events.Create))
			Expect(resourceChanged.Type).To(Equal(model.ResourceType("Mesh")))

			// when postgres is stopped
			err = c.Stop()

			// then
			Expect(err).To(Not(HaveOccurred()))
			Consistently(storeErrCh, "1s", "100ms").Should(Receive())

			// when postgres is restarted
			err = c.Start()

			// then
			Expect(err).To(Not(HaveOccurred()))
			event, err = listener.Recv(eventBusStopCh)
			Expect(err).To(Not(HaveOccurred()))
			resourceChanged = event.(kuma_events.ResourceChangedEvent)
			Expect(resourceChanged.Operation).To(Equal(kuma_events.Create))
			Expect(resourceChanged.Type).To(Equal(model.ResourceType("Mesh")))
		},
		Entry("When using pq", postgres_config.DriverNamePq),
		Entry("When using pgx", postgres_config.DriverNamePgx),
	)
})

func setupChannels() (chan struct{}, chan error, chan struct{}, chan error) {
	listenerStopCh := make(chan struct{})
	listenerErrCh := make(chan error)
	eventBusStopCh := make(chan struct{})
	storeErrCh := make(chan error)

	return listenerStopCh, listenerErrCh, eventBusStopCh, storeErrCh
}

func setupStore(cfg postgres_config.PostgresStoreConfig) store.ResourceStore {
	metrics, err := core_metrics.NewMetrics("Standalone")
	Expect(err).ToNot(HaveOccurred())
	pStore, err := NewStore(metrics, cfg)
	Expect(err).ToNot(HaveOccurred())
	return pStore
}

func setupListeners(cfg postgres_config.PostgresStoreConfig, driverName string, listenerErrCh chan error, listenerStopCh chan struct{}) kuma_events.Listener {
	cfg.DriverName = driverName
	eventsBus := kuma_events.NewEventBus()
	listener := eventsBus.New()
	l := postgres_events.NewListener(cfg, eventsBus)
	resilientListener := component.NewResilientComponent(core.Log.WithName("postgres-event-listener-component"), l)
	go func() {
		listenerErrCh <- resilientListener.Start(listenerStopCh)
	}()

	return listener
}

func triggerNotifications(cfg postgres_config.PostgresStoreConfig, storeErrCh chan error) {
	pStore := setupStore(cfg)
	defer GinkgoRecover()
	for i := 0; !channels.IsClosed(storeErrCh); i++ {
		err := pStore.Create(context.Background(), mesh.NewMeshResource(), store.CreateByKey(fmt.Sprintf("mesh-%d", i), ""))
		if err != nil {
			storeErrCh <- err
		}
	}
}

func channelClosesWithoutErrors(listenerErrCh chan error) func() bool {
	return func() bool {
		select {
		case err := <-listenerErrCh:
			Expect(err).ToNot(HaveOccurred())
			return true
		default:
			return false
		}
	}
}
