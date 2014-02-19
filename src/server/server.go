package server

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"strconv"
	"todoRepos"
)

type HttpErr struct {
	Message string `json:"message"`
}

func Start() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func() string {
		return "Hello, World!"
	})

	m.Get("/todos/:id", func(params martini.Params, r render.Render) {
		id, convErr := strconv.Atoi(params["id"])
		if convErr != nil {
			r.JSON(400, HttpErr{"invalid id param"})
			return
		}
		todo, getErr := todoRepos.Get(id)
		if getErr != nil {
			r.JSON(404, HttpErr{fmt.Sprintf("Todo %v not found", id)})
			return
		}
		r.JSON(200, todo)
	})

	m.Run()
}
