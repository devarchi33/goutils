package ctxdb

import (
	"context"

	"github.com/devarchi33/goutils/kafka"
	"xorm.io/xorm"
)

type ContextDB struct {
	*xorm.Engine
}

func New(db *xorm.Engine, service string, config kafka.Config) *ContextDB {
	// db.ShowExecTime()
	if len(config.Brokers) != 0 {
		if producer, err := kafka.NewProducer(config.Brokers, config.Topic,
			kafka.WithDefault(),
			kafka.WithTLS(config.SSL)); err == nil {
			db.SetLogger(&dbLogger{serviceName: service, Producer: producer})
			db.ShowSQL()
		}
	}

	return &ContextDB{Engine: db}
}

func (db *ContextDB) NewSession(ctx context.Context) *xorm.Session {
	session := db.Engine.NewSession()

	func(session interface{}, ctx context.Context) {
		if s, ok := session.(interface{ SetContext(context.Context) }); ok {
			s.SetContext(ctx)
		}
	}(session, ctx)

	return session
}
