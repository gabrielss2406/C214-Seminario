package main

import (
    "log"
    "net/http"
    "todolist/routes"
)

func main() {
    r := routes.SetupRoutes()
    log.Println("Servidor rodando na porta 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
