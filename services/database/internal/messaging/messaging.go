package messaging

import (
	"database/internal/database"
	"database/internal/logger"
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func SetupSubscribers(nc *nats.Conn, db *sql.DB) {
	nc.Subscribe("database.health", func(m *nats.Msg) {
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

	nc.Subscribe("database.product.create", func(m *nats.Msg) {
		var product database.Product
		err := json.Unmarshal(m.Data, &product)
		if err != nil {
			log.Error("Error on json product data:", err)
			nc.Publish(m.Reply, []byte("ERROR 2001"))
			return
		}

		productId, err := database.CreateProduct(db, product)
		if err != nil {
			log.Error("Error on product create:", err)
			nc.Publish(m.Reply, []byte("ERROR 2002"))
			return
		}

		productIdJSON, err := json.Marshal(productId)
		if err != nil {
			log.Error("Error on json encoding:", err)
			nc.Publish(m.Reply, []byte("ERROR 2003"))
			return
		}

		log.Debug("Success product create:", productId)
		nc.Publish(m.Reply, productIdJSON)
	})

	nc.Subscribe("database.product.get.all", func(m *nats.Msg) {
		productData, err := database.GetAllProductsWithStock(db)
		if err != nil {
			log.Error("Error on product get:", err)
			nc.Publish(m.Reply, []byte("ERROR 2004"))
			return
		}

		productDataJSON, err := json.Marshal(productData)
		if err != nil {
			log.Error("Error on json encoding:", err)
			nc.Publish(m.Reply, []byte("ERROR 2005"))
			return
		}

		log.Debug("Success product get:", productData)
		nc.Publish(m.Reply, productDataJSON)
	})

	nc.Subscribe("database.product.get.*", func(m *nats.Msg) {
		parts := strings.Split(m.Subject, ".")
		numberPart := parts[len(parts)-1]
		productId, err := strconv.Atoi(numberPart)
		if err != nil {
			log.Error("Error al convertir '%s' a número: %v", numberPart, err)
			nc.Publish(m.Reply, []byte("ERROR 2006"))
			return
		}

		productData, err := database.GetProduct(db, productId)
		if err != nil {
			log.Error("Error on productId get:", err)
			nc.Publish(m.Reply, []byte("ERROR 2007"))
			return
		}

		productDataJSON, err := json.Marshal(productData)
		if err != nil {
			log.Error("Error on json encoding:", err)
			nc.Publish(m.Reply, []byte("ERROR 2008"))
			return
		}

		log.Debug("Success productId get:", productData)
		nc.Publish(m.Reply, productDataJSON)
	})

	nc.Subscribe("database.product.set.stock", func(m *nats.Msg) {
		var productStock database.ProductStock
		err := json.Unmarshal(m.Data, &productStock)
		if err != nil {
			log.Error("Error on json productStock data:", err)
			nc.Publish(m.Reply, []byte("ERROR 2009"))
			return
		}

		err = database.UpsertProductStock(db, productStock.ProductID, productStock.Quantity)
		if err != nil {
			log.Error("Error on productStock set:", err)
			nc.Publish(m.Reply, []byte("ERROR 2010"))
			return
		}

		log.Debug("Success productStock set:", productStock)
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("database.product.remove", func(m *nats.Msg) {
		var product database.Product
		err := json.Unmarshal(m.Data, &product)
		if err != nil {
			log.Error("Error on json product data:", err)
			nc.Publish(m.Reply, []byte("ERROR 2011"))
			return
		}

		err = database.RemoveProduct(db, product.ProductID)
		if err != nil {
			log.Error("Error on product delete:", err)
			nc.Publish(m.Reply, []byte("ERROR 2012"))
			return
		}

		log.Debug("Success product delete:", product)
		nc.Publish(m.Reply, []byte("OK"))
	})
}
