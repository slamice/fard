package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Group("/users", func(r martini.Router) {
    r.Get("/:id", controllers.getUser)
    r.Post("/new", controllers.addUser)
    r.Put("/update/:id", controllers.UpdateUser)
    r.Delete("/delete/:id", controllers.deleteUser)
  })

  m.Run()
}
