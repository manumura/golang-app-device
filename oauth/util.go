package oauth

import (
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
)

// OutputJSON encodes the Response to JSON and writes to the http.ResponseWriter
func OutputJSON(rs *osin.Response, c echo.Context) error {
	// Add headers
	for i, k := range rs.Headers {
		for _, v := range k {
			c.Response().Header().Add(i, v)
		}
	}

	if rs.Type == osin.REDIRECT {
		// Output redirect with parameters
		u, err := rs.GetRedirectUrl()
		if err != nil {
			return err
		}
		return c.Redirect(http.StatusFound, u)
	}

	return c.JSON(rs.StatusCode, rs.Output)
}
