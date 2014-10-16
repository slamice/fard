package controllers

import (
    "net/http"
    "github.com/go-martini/martini"
    "encoding/json"
    "log"
)

type UserResponse struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    Id         int    `json:"id"`
}

func AddUserHandler(rw http.ResponseWriter, req *http.Request, uid int, p martini.Params) {
    
    decoder := json.NewDecoder(req.Body)
    log.Println(decoder)

    // err := decoder.Decode(&sf)

    // if err != nil {
    //     http.Error(rw, err.Error(), http.StatusBadRequest)
    // }

    return
}