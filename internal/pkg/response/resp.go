package response

import (
	"fmt"
	"strings"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/labstack/echo/v4"
)

func Respond(c echo.Context, resp models.Response) error {
	if resp.Status == 0 {
		resp.Status = 200
	}

	return c.JSON(resp.Status, resp.ToMap())
}

func RespondError(c echo.Context, status echo.HTTPError, msg ...string) error {
	message := ""

	if len(msg) == 0 {
		message = fmt.Sprint(status.Message)
	} else {
		message = strings.Join(msg, " ")
	}

	return Respond(c, models.Response{
		Status:           status.Code,
		Message:          message,
		DocumentationUrl: "https://example.com"})
}
