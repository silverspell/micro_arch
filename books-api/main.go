package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Type   string `json:"type"`
}

var books [2]Book
var mayEncounterError bool

func init() {
	mayEncounterError = os.Getenv("MAY_ENCOUNTER_ERROR") == "1"

	b1 := Book{
		Id:     "BOOK_1",
		Name:   "Death of the Fake Lynx",
		Author: "Fake Lynx",
		Type:   "criminal",
	}

	b2 := Book{
		Id:     "BOOK_2",
		Name:   "Sign of the Ghostly Tuba",
		Author: "Ghostly Tuba",
		Type:   "sci-fi",
	}
	books = [2]Book{b1, b2}
}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, books)
	})

	e.GET("/:id", func(c echo.Context) error {
		fmt.Printf("mayEncounterError: %v\n", mayEncounterError)

		if mayEncounterError {
			r := rand.Intn(100)
			fmt.Printf("r: %v\n", r)
			if r < 30 {
				return c.JSON(500, "Internal server error")
			}
		}
		bookId := c.Param("id")
		for _, v := range books {
			if bookId == v.Id {
				return c.JSON(200, v)
			}
		}
		return c.JSON(404, "Not found")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
