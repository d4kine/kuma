package postgres

import (
	"context"
	"sync"

	"github.com/go-logr/logr"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/config/plugins/resources/postgres"
)

// PgxListener will listen for NOTIFY commands on a channel.
type PgxListener struct {
	notificationsCh chan *Notification
	err             error
	mu              sync.Mutex

	logger logr.Logger

	db *pgxpool.Pool

	stopFn func()
}

func (l *PgxListener) Error() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.err
}

var _ Listener = (*PgxListener)(nil)

// NewPgxListener will create and initialize a PgxListener which will automatically connect and listen to the provided channel.
func NewPgxListener(config postgres.PostgresStoreConfig, logger logr.Logger) (Listener, error) {
	ctx := context.Background()
	connectionString, err := config.ConnectionString()
	if err != nil {
		return nil, err
	}
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	l := &PgxListener{
		notificationsCh: make(chan *Notification, 32),
		logger:          logger,
		db:              db,
	}
	l.start(ctx)

	return l, nil
}

func (l *PgxListener) start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	l.stopFn = cancel

	go l.run(ctx)
	l.logger.Info("started")
}

func (l *PgxListener) Close() error {
	l.stopFn() // this closes the channel by canceling the context
	l.logger.Info("stopped")
	return nil
}

func (l *PgxListener) run(ctx context.Context) {
	err := l.handleNotifications(ctx)
	if err != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.err = err
		close(l.notificationsCh)
	}
}

func (l *PgxListener) handleNotifications(ctx context.Context) error {
	conn, err := l.db.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting connection")
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "listen "+channelName)
	if err != nil {
		return err
	}

	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			return err
		}

		l.logger.V(1).Info("event happened", "event", notification)
		select {
		case l.notificationsCh <- toBareNotification(notification):
		default:
			l.logger.V(1).Info("buffer full, event dropped", "event", notification)
		}
	}
}

func toBareNotification(notification *pgconn.Notification) *Notification {
	return &Notification{
		Payload: notification.Payload,
	}
}

func (l *PgxListener) Notify() chan *Notification { return l.notificationsCh }
