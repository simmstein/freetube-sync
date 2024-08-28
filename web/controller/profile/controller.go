package profile

import (
	"github.com/labstack/echo/v4"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

func Ko(c echo.Context, err error) error {
	return c.JSON(400, map[string]any{
		"code":    400,
		"message": err,
	})
}

func Ok(c echo.Context) error {
	return c.JSON(201, map[string]any{
		"code":    201,
		"message": "ok",
	})
}

func OkKo(c echo.Context, err error) error {
	if err != nil {
		return Ko(c, err)
	}

	return Ok(c)
}

func Init(c echo.Context) error {
	// payload := []model.Video{}
	// err := c.Bind(&payload)
	var err error

	return OkKo(c, err)
}

func Push(c echo.Context) error {
	// payload := []model.Video{}
	// err := c.Bind(&payload)
	var err error

	return OkKo(c, err)
}

func Pull(c echo.Context) error {
	// payload := []model.Video{}
	// err := c.Bind(&payload)
	var err error

	return OkKo(c, err)
}

func Register(e *echo.Echo) {
	e.POST(route.ProfileInit, Init)
	e.POST(route.ProfilePush, Push)
	e.GET(route.ProfilePull, Pull)
}
