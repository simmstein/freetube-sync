package helper

import "github.com/labstack/echo/v4"

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
