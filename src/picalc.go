package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func say(word string) string {
    title := "[ 20191123 ] Hello world, this is"
    return title + " " + word
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Say something!")
    fmt.Fprintf(w, "%s\n", say("BLUE11231 Hello taoxb it's knative testing!!!"))
}

func main() {
    log.Print("Say something started.")

    http.HandleFunc("/", handler)   

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
