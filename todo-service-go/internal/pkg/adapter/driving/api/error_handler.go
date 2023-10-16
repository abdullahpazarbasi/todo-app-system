package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
)

func errorHandler(cause error, ec echo.Context) {
	if ec.Response().Committed {
		return
	}

	fullDetailed := ec.Echo().Debug
	var httpStatusCode int
	var rs map[string]interface{}
	switch outerError := cause.(type) {
	case domainFaultPort.Fault:
		httpStatusCode = outerError.ProposedHTTPStatusCode()
		rs = outerError.Normalize(fullDetailed)
	case *echo.HTTPError:
		switch innerError := outerError.Message.(type) {
		case domainFaultPort.Fault:
			httpStatusCode = innerError.ProposedHTTPStatusCode()
			rs = innerError.Normalize(fullDetailed)
		default:
			httpStatusCode = outerError.Code
			if fullDetailed {
				message, ok := outerError.Message.(string)
				if !ok {
					message = http.StatusText(httpStatusCode)
				}
				if outerError.Internal == nil {
					rs = map[string]interface{}{"message": message}
				} else {
					rs = map[string]interface{}{"message": message, "error": outerError.Internal.Error()}
				}
			} else {
				rs = map[string]interface{}{"message": http.StatusText(httpStatusCode)}
			}
		}
	default:
		httpStatusCode = http.StatusInternalServerError
		if fullDetailed {
			rs = map[string]interface{}{"message": outerError.Error()}
		} else {
			rs = map[string]interface{}{"message": http.StatusText(httpStatusCode)}
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
