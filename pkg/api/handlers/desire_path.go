package handlers

import (
	"github.com/feldtsen/farrago/pkg/models"
	"github.com/feldtsen/farrago/view/page/desirePath"
	"github.com/labstack/echo/v4"
)

type DesirePath struct{}

func (h *DesirePath) Handler(c echo.Context) error {
	paginationModel := c.Get("paginationModel").(models.Pagination)

	return renderByHXRequest(c, desirePath.DesirePathPartial(paginationModel), desirePath.DesirePathPage(paginationModel))
}
