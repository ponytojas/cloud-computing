package messaging

import (
	"auth/internal/logger"
	"auth/internal/token"
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func SetupSubscribers(nc *nats.Conn) {
	nc.Subscribe("health", func(m *nats.Msg) {
		log.Info("Health check")
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("register", func(m *nats.Msg) {
		response, err := nc.Request("database.users.create", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error al crear el usuario:", err)
			nc.Publish(m.Reply, []byte("Error al crear el usuario"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error al crear el usuario"))
		}
	})

	nc.Subscribe("login", func(m *nats.Msg) {
		response, err := nc.Request("database.users.login", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error al iniciar sesión:", err)
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		}

		if string(response.Data) == "ERROR" {
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		} else {
			log.Debug("Login correcto")
		}

		var usercheck token.AuthCheck
		err = json.Unmarshal(response.Data, &usercheck)
		if err != nil {
			log.Error("Error al iniciar sesión:", err)
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		}

		token, err := token.CreateToken(usercheck)
		if err != nil {
			log.Error("Error al crear el token:", err)
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		}

		nc.Publish(m.Reply, []byte(token))

	})

}
