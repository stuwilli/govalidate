package rules

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/stuwilli/govalidate/helper"
)

//URL Validates a URL using url.Parse() in the net/url library.
// For a valid URL, the following need to be present in a parsed URL:
// * Scheme (either http or https)
// * Host (without a backslash)
func URL(data ValidationData) error {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	parsed, err := url.Parse(v)
	if err != nil {
		return ErrInvalid{
			ValidationData: data,
			Failure:        "is not a valid URL",
			Message:        data.Message,
		}
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf("has an invalid scheme '%s'", parsed.Scheme),
			Message:        data.Message,
		}
	}

	if parsed.Host == "" || strings.IndexRune(parsed.Host, '\\') > 0 {
		return ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf("has an invalid host ('%s')", parsed.Host),
			Message:        data.Message,
		}
	}

	return nil
}
