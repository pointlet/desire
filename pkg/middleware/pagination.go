package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/feldtsen/farrago/pkg/models"
	"github.com/labstack/echo/v4"
)

func Pagination(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid 'page' parameter, must be an integer"})
		}

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid 'pageSize' parameter, must be an integer"})
		}

		if pageSize < 1 || pageSize > 100 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid 'pageSize' parameter, must be between 1 and 100"})
		}

		offset := (page - 1) * pageSize

		paginationModel := models.PaginationModel{
			Page:     page,
			PageSize: pageSize,
			Offset:   offset,
		}

		c.Set("paginationModel", paginationModel)

		fmt.Println("paginationModel: ", paginationModel)

		return next(c)
	}
}

func GeneratePaginatedURL(endpoint string, page int, pageSize int) string {
	fmt.Println("endpoint: ", endpoint)
	return fmt.Sprintf("%s?page=%d&pageSize=%d", endpoint, nextPage(page), pageSize)
}

func nextPage(page int) int {
	return page + 1
}
