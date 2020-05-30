package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func say(word string) string {
    title := "[ 202005301 ] Hello world, this is"
    return title + " " + word
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Say something, it's in container...05301!")
    fmt.Fprintf(w, "%s\n", say("GREEN202005301 Hello leo it's knative testing GREEN on github!!!"))
}

func main() {
    log.Print("Say something started.202005301 test hello leo")

    http.HandleFunc("/", handler)   

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
