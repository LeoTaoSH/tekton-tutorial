package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func say(word string) string {
    title := "[ 2020062105 ] Hello world, this is"
    return title + " " + word
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Say something, it's in container..20200621!")
    fmt.Fprintf(w, "%s\n", say("GREEN20201006 Hello leo it's knative testing GREEN on github!!!"))
}

func main() {
    log.Print("Say something started.2020062101 test hello leo")

    http.HandleFunc("/", handler)   

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
