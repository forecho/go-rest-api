package server

import (
	"fmt"
	"github.com/forecho/go-rest-api/internal/responses"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type customErrHandler struct {
	e *echo.Echo
}

type resp struct {
	Message interface{} `json:"message"`
}

var validationErrors = map[string]string{
	"required": " is required, but was not received",
	"min":      "'s value or length is less than allowed",
	"max":      "'s value or length is bigger than allowed",
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}

func (ce *customErrHandler) handler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if ce.e.Debug {
		msg = err.Error()
		switch err.(type) {
		case *echo.HTTPError:
			code = err.(*echo.HTTPError).Code
		case validator.ValidationErrors:
			code = http.StatusBadRequest
		}
	} else {
		switch err.(type) {
		case *echo.HTTPError:
			e := err.(*echo.HTTPError)
			code = e.Code
			msg = e.Message
			if e.Internal != nil {
				msg = fmt.Sprintf("%v, %v", err, e.Internal)
			}
		case validator.ValidationErrors:
			var errMsg []string
			e := err.(validator.ValidationErrors)
			for _, v := range e {
				errMsg = append(errMsg, fmt.Sprintf("%s%s", v.Field(), getVldErrorMsg(v.ActualTag())))
			}
			msg = resp{Message: errMsg}
			code = http.StatusBadRequest
		default:
			msg = http.StatusText(code)
		}
		if _, ok := msg.(string); ok {
			msg = resp{Message: msg}
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = responses.Response(c, code, msg)
		}
		if err != nil {
			ce.e.Logger.Error(err)
		}
	}
}
