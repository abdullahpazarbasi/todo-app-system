package client_interface_adapters_rest_api_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	drivingAppPortsError "todo-app-wbff/internal/pkg/app/ports/driving/error"
)

func ErrorHandler(cause error, ec echo.Context) {
	if ec.Response().Committed {
		return
	}

	fullDetailed := ec.Echo().Debug
	var httpStatusCode int
	var rs map[string]interface{}
	switch err := cause.(type) {
	case drivingAppPortsError.ServiceError:
		httpStatusCode = err.ProposedHTTPStatusCode()
		rs = err.Normalize(fullDetailed)
	case *echo.HTTPError:
		httpStatusCode = err.Code
		if fullDetailed {
			message, ok := err.Message.(string)
			if !ok {
				message = http.StatusText(httpStatusCode)
			}
			if err.Internal == nil {
				rs = map[string]interface{}{"message": message}
			} else {
				rs = map[string]interface{}{"message": message, "error": err.Internal.Error()}
			}
		} else {
			rs = map[string]interface{}{"message": http.StatusText(httpStatusCode)}
		}
	default:
		httpStatusCode = http.StatusInternalServerError
		if fullDetailed {
			rs = map[string]interface{}{"message": err.Error()}
		} else {
			rs = map[string]interface{}{"message": "error occurred"}
		}
	}

	var err error
	if ec.Request().Method == http.MethodHead {
		err = ec.NoContent(httpStatusCode)
	} else {
		err = ec.JSON(httpStatusCode, rs)
	}
	if err != nil {
		ec.Logger().Error(err)
	}
}
