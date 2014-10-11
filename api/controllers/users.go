package controllers

import (
    "strconv"
    "github.com/go-martini/martini"
)

type UserResponse struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    Id         int    `json:"id"`
}

func createUserHandler(rw http.ResponseWriter, req *http.Request, uid int, p martini.Params) {
    id, err := strconv.Atoi(p["id"])
}