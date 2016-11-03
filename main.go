package main

import (
        "io"
        "log"
        "net/http"
        "fmt"
        "html"

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

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        })
        http.HandleFunc("/health", ping)
        log.Fatal(http.ListenAndServe(":8000", nil))
}
