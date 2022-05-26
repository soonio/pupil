package http

import (
	json "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

// JsonSerializer implements JSON encoding using json-iterator/go.
type JsonSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d JsonSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d JsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(i)
	return err
}
