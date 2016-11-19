package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/minions", MinionIndex)
    router.HandleFunc("/minions/{minionId}", MinionShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}


func MinionIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "The minion Index page!")
}

func MinionShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    minionId := vars["minionId"]
    fmt.Fprintln(w, "Minion show:", minionId)
}
