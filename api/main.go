package api

import (
  "github.com/go-martini/martini"
  "github.com/slamice/fard/api/controllers"
)

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Group("/users", func(r martini.Router) {
//    r.Get("/:id", controllers.getUser)
    r.Post("/add", controllers.AddUserHandler)
//    r.Put("/update/:id", controllers.UpdateUser)
//    r.Delete("/delete/:id", controllers.deleteUser)
  })

  m.Run()
}
