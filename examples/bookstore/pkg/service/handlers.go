package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// booksHandler returns all books from db
func booksHandler(ctx echo.Context) error {
	var books []Book
	result := Db.Find(&books)
	if result.Error != nil {
		return internalServerErr(ctx, result.Error)
	}
	return ctx.JSON(http.StatusOK, books)
}

// getBookByIDHandler returns a particular book
func getBookByIDHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var book Book
	result := Db.First(&book, id)
	if result.Error != nil {
		if "record not found" == result.Error.Error() {
			resp := make(map[string]string)
			resp["message"] = "Not Found"
			return ctx.JSON(http.StatusNotFound, resp)
		}
		return internalServerErr(ctx, result.Error)
	}

	return ctx.JSON(http.StatusOK, book)
}

// addBookHandler adds a book to db
func addBookHandler(ctx echo.Context) error {

	bdy := ctx.Request().Body
	defer func(bdy io.ReadCloser) {
		_ = bdy.Close()
	}(bdy)

	var book Book

	err := json.NewDecoder(bdy).Decode(&book)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	result := Db.Create(&book)
	if result.Error != nil {
		return internalServerErr(ctx, result.Error)
	}

	return ctx.NoContent(http.StatusCreated)
}

// updateBookHandler updates data for a particular book
func updateBookHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	bdy := ctx.Request().Body
	defer func(bdy io.ReadCloser) {
		_ = bdy.Close()
	}(bdy)

	var book Book

	err = json.NewDecoder(bdy).Decode(&book)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	book.ID = uint(id)

	res := Db.Save(&book)
	if res.Error != nil {
		return internalServerErr(ctx, res.Error)
	}

	return ctx.NoContent(http.StatusAccepted)
}

// deleteBookHandler removes a book from db
func deleteBookHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res := Db.Delete(&Book{}, id)
	if res.Error != nil {
		return internalServerErr(ctx, res.Error)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// internalServerErr returns a 500 HTTP response
func internalServerErr(ctx echo.Context, err error) error {
	zap.L().Error(err.Error())
	return ctx.NoContent(http.StatusInternalServerError)
}
