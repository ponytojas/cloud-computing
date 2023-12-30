package messaging

import (
	"store/internal/logger"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func SetupSubscribers(nc *nats.Conn) {
	nc.Subscribe("store.health", func(m *nats.Msg) {
		log.Info("Health check")
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("createProduct", func(m *nats.Msg) {
		response, err := nc.Request("database.product.create", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating product:", err)
			nc.Publish(m.Reply, []byte("Error creating product"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating product"))
		}
	})

	nc.Subscribe("getProducts", func(m *nats.Msg) {
		response, err := nc.Request("database.product.get.all", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating product:", err)
			nc.Publish(m.Reply, []byte("Error creating product"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating product"))
		}
	})

	nc.Subscribe("setStock", func(m *nats.Msg) {
		response, err := nc.Request("database.stock.set", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating product:", err)
			nc.Publish(m.Reply, []byte("Error creating product"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating product"))
		}
	})

	nc.Subscribe("removeProduct", func(m *nats.Msg) {
		response, err := nc.Request("database.product.remove", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating product:", err)
			nc.Publish(m.Reply, []byte("Error creating product"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating product"))
		}
	})

	nc.Subscribe("getStock", func(m *nats.Msg) {
		response, err := nc.Request("database.stock.get", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating product:", err)
			nc.Publish(m.Reply, []byte("Error creating product"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating product"))
		}
	})

}
