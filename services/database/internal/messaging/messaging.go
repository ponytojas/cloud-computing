package messaging

import (
	"database/internal/database"
	"database/internal/logger"
	"database/sql"
	"encoding/json"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func SetupSubscribers(nc *nats.Conn, db *sql.DB) {
	nc.Subscribe("health", func(m *nats.Msg) {
		log.Info("Health check")
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("database.users.create", func(m *nats.Msg) {
		var user database.User
		err := json.Unmarshal(m.Data, &user)
		if err != nil {
			log.Error("Error on json user data:", err)
			nc.Publish(m.Reply, []byte("ERROR 1000"))
			return
		}

		id, err := database.CreateUser(db, user)
		if err != nil {
			log.Error("Error creating user:", err)
			nc.Publish(m.Reply, []byte("ERROR 1001"))
			return
		}

		log.Debug("User created with ID:", id)
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("database.users.login", func(m *nats.Msg) {
		var user database.User
		err := json.Unmarshal(m.Data, &user)
		if err != nil {
			log.Error("Error on json user data:", err)
			nc.Publish(m.Reply, []byte("ERROR 1002"))
			return
		}

		usercheck, err := database.LoginUser(db, user)
		if err != nil {
			log.Error("Error on user login:", err)
			nc.Publish(m.Reply, []byte("ERROR 1003"))
			return
		}

		usercheckJSON, err := json.Marshal(usercheck)
		if err != nil {
			log.Error("Error on json encoding:", err)
			nc.Publish(m.Reply, []byte("ERROR 1004"))
			return
		}

		log.Debug("Success login:", usercheck)
		nc.Publish(m.Reply, usercheckJSON)
	})

}
