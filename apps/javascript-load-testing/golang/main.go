package main

import (
	"strconv"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/hc", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": "ok"})
	})

	app.Get("/fibonacci/{nth}", func(ctx iris.Context) {
		nth, _ := strconv.Atoi(ctx.Params().Get("nth"))

		result := fibonnaci(nth)
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
