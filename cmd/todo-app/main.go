package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katsuharu/todo-app/application"
	"github.com/katsuharu/todo-app/handler"
	"github.com/katsuharu/todo-app/infra/dao"
	"github.com/katsuharu/todo-app/infra/db"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db, err := db.NewDB()
	if err != nil {
		return
	}

	tr := dao.NewTodo(db)
	ta := application.NewTodo(tr)
	th := handler.NewTodo(ta)

	e.POST("/todos", th.Create)
	e.GET("/todos", th.List)

	e.Logger.Fatal(e.Start(":1323"))
}
