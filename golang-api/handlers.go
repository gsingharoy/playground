package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func MinionIndex(w http.ResponseWriter, r *http.Request) {
    minions := Minions{
        Minion{Name: "Dave"},
        Minion{Name: "Bob"},
    }

    if err := json.NewEncoder(w).Encode(minions); err != nil {
        panic(err)
    }
}

func MinionShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    minionId := vars["minionId"]
    fmt.Fprintln(w, "Minion show:", minionId)
}
