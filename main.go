package main

import (
        "io"
        "log"
        "net/http"
        "os"

        "gopkg.in/redis.v5"
)

func ping(w http.ResponseWriter, r *http.Request) {
        client := redis.NewClient(&redis.Options{
                Addr:     "redis:6379",
                Password: "", // no password set
                DB:       0,  // use default DB
        })

        pong, err := client.Ping().Result()
        if err != nil {
                io.WriteString(w, err.Error())
        } else {
                io.WriteString(w, pong)
        }
}

func returnHostname(w http.ResponseWriter, r *http.Request) {
        name, err := os.Hostname()
        if err != nil {
                io.WriteString(w, err.Error())
        } else {
                io.WriteString(w, name)
        }
}

func main() {
        http.HandleFunc("/health", ping)
        http.HandleFunc("/", returnHostname)
        log.Fatal(http.ListenAndServe(":8000", nil))
}
