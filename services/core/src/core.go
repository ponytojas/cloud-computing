package main

import (
    "log"
    "net/http"
    "time"

    "github.com/nats-io/nats.go"
)

func main() {
    // Conectar a NATS
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Endpoint para la autenticación
    http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
        // Publicar la solicitud al tema de autenticación y esperar por la respuesta
        msg, err := nc.Request("auth-service", []byte("Payload de autenticación"), 2*time.Second)
        if err != nil {
            log.Println(err)
            http.Error(w, "Error de autenticación", http.StatusInternalServerError)
            return
        }

        // Escribir la respuesta de la autenticación al cliente
        w.WriteHeader(http.StatusOK)
        w.Write(msg.Data)
    })

    // Otros endpoints...

    log.Println("Core microservice running on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
