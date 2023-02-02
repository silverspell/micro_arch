package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReadingList struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Books []string `json:"books"`
}

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Type   string `json:"type"`
}

type ReadingListResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

var lists [3]ReadingList

func init() {
	l1 := ReadingList{
		Id:    "LIST_1",
		Name:  "List 1",
		Books: []string{"BOOK_1"},
	}

	l2 := ReadingList{
		Id:    "LIST_2",
		Name:  "List 2",
		Books: []string{"BOOK_1", "BOOK_2"},
	}

	l3 := ReadingList{
		Id:    "LIST_3",
		Name:  "List 3",
		Books: []string{"BOOK_2"},
	}

	lists = [3]ReadingList{l1, l2, l3}
}

func main() {
	e := echo.New()
	e.GET("/:id", func(c echo.Context) error {
		var list ReadingList
		p := c.Param("id")
		for _, v := range lists {
			if v.Id == p {
				list = v
				break
			}
		}

		books := make([]Book, 0)
		for _, v := range list.Books {
			var b Book

			bookResp, err := http.Get("http://books-api:8080/" + v)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return err
			}
			defer bookResp.Body.Close()
			bookStr, _ := io.ReadAll(bookResp.Body)
			json.Unmarshal(bookStr, &b)
			books = append(books, b)
		}
		readingListResp := ReadingListResponse{
			Id:    list.Id,
			Name:  list.Name,
			Books: books,
		}
		return c.JSON(200, readingListResp)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
