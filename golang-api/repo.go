package main

import "fmt"

var currentId int

var minions Minions

// Give us some seed data
func init() {
  RepoCreateMinion(Minion{Name: "Bob"})
  RepoCreateMinion(Minion{Name: "Stuart"})
}

func RepoFindMinion(id int) Minion {
  for _, m := range minions {
    if m.Id == id {
      return m
    }
  }
  // return empty Minion if not found
  return Minion{}
}

func RepoCreateMinion(m Minion) Minion {
  currentId += 1
  m.Id = currentId
  minions = append(minions, m)
  return m
}

func RepoDestroyMinion(id int) error {
  for i, m := range minions {
    if m.Id == id {
      minions = append(minions[:i], minions[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find Minion with id of %d to delete", id)
}
