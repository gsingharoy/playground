package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
  "io"
  "strconv"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func MinionIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(minions); err != nil {
    panic(err)
  }
}

func MinionShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var minionId int
  var err error
  if minionId, err = strconv.Atoi(vars["minionId"]); err != nil {
		panic(err)
	}
  minion := RepoFindMinion(minionId)
  if minion.Id > 0 {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(minion); err != nil {
      panic(err)
    }
    return
  }

  // If we didn't find it, 404
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusNotFound)
  if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
    panic(err)
  }
}

func MinionCreate(w http.ResponseWriter, r *http.Request) {
  var minion Minion
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &minion); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  m := RepoCreateMinion(minion)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(m); err != nil {
    panic(err)
  }
}
