package models

import (
    "errors"
    "time"
    "github.com/coopernurse/gorp"
)

type User struct {
    Id                int            `db:"id"`
    Password          string         `db:"password"`
    Email             string         `db:"email"`
    FirstName         string         `db:"firstName"`
    LastName          string         `db:"lastName"`
    ctime             time.Time      `db:"ctime" json:"ctime,omitempty"`
    utime             time.Time      `db:"utime" json:"utime,omitempty"`
}

func InsertUser(password, email, firstName, lastName) (id int, err error) {
    m := User{
        Email:      email,
        FirstName:  firstName,
        LastName:   lastName,
        ctime:      ctime,
        utime:      utime
    }

    u := User(m)
    err = dbmap.Insert(&u)
    if err != nil {
        return id, err
    }

    id = u.Id

    return id, nil
}