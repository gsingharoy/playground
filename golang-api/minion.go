package main

import "time"

type Minion struct {
  Id            int       `json:"id"`
  Name          string    `json:"name"`
  LinkesBanana  bool      `json:"likes_banana"`
  DOB           time.Time `json:"dob"`
}

type Minions []Minion
