package driven_adapter_todo

import (
	"fmt"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
)

func translateHTTPErrorToFault(f domainFaultPort.Factory, httpError error) domainFaultPort.Fault {
	return f.WrapError(
		httpError,
		f.Message(fmt.Sprintf("error %T occurred", httpError)),
	)
}
