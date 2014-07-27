package models

import (
	"fmt"
	"github.com/robfig/revel"
	"regexp"
)

type User struct {
	Name     string
	UserName string
	Email    string
	Password []byte
}