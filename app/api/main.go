package main

import (
	"net/http"

	"api/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	// reids
	// repository.SetupRedis()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.GET("users/:uuid", getUserList)
	// e.GET("ranking", ranking)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUserList(c echo.Context) error {
	// リクエストからIDを取得
	uuid := c.Param("uuid")

	// redisからデータを取得
	userList, err := repository.GetUserList(uuid)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, userList)
}

func ranking(c echo.Context) error {
	result, err := repository.GetRankings()

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}
