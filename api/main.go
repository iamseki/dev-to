package main

import (
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/hc", func(ctx iris.Context) {
		time.Sleep(1 * time.Minute)
		ctx.JSON(iris.Map{"status": "ok"})
	})

	app.Get("/ue", func(ctx iris.Context) {
		ctx.StatusCode(422)
	})

	app.Get("/fb", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusForbidden)
	})

	app.Get("/fibonacci/{nth}", func(ctx iris.Context) {
		nth, _ := strconv.Atoi(ctx.Params().Get("nth"))

		result := fibonnaci(nth)

		ctx.Header("Cache-Control", "max-age=3600")
		ctx.JSON(iris.Map{"fibonacci": result})
	})

	app.Listen(":8080")
}

func fibonnaci(n int) int {
	x := 0
	y := 1

	for i := 0; i <= n; i++ {
		tmp := y
		y += x
		x = tmp
	}

	return y
}
