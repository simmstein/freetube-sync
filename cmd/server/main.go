package main

import (
	"github.com/labstack/echo/v4"
	"gitnet.fr/deblan/freetube-sync/model"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

func main() {
	e := echo.New()

	e.POST(route.HistoryInit, func(c echo.Context) error {
		payload := []model.Video{}
		err := c.Bind(&payload)

		if err != nil {
			return c.JSON(400, map[string]any{
				"code":    400,
				"message": err,
			})
		}

		return c.JSON(201, map[string]any{
			"code":    201,
			"message": "ok",
		})
	})

	e.POST(route.HistoryPush, func(c echo.Context) error {
		payload := []model.Video{}
		err := c.Bind(&payload)

		if err != nil {
			return c.JSON(400, map[string]any{
				"code":    400,
				"message": err,
			})
		}

		return c.JSON(201, map[string]any{
			"code":    201,
			"message": "ok",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
