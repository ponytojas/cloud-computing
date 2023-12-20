package main

import (
    "log"

    "github.com/nats-io/nats.go"
)

func main() {
    // Conexión a NATS
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Suscripción a un tema para recibir datos y evaluar alertas
    nc.Subscribe("patient-data", func(m *nats.Msg) {
        // Aquí va la lógica para evaluar los datos y enviar alertas
        log.Printf("Recibido mensaje: %s\n", string(m.Data))
    })

    log.Println("Servicio de alertas iniciado y escuchando mensajes")
    select {} // Bloquea la goroutine principal indefinidamente
}
