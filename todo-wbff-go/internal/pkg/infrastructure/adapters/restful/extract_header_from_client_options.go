package infrastructure_adapters_restful

import (
	"encoding/base64"
	"fmt"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

func extractHeaderFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (*map[string][]string, *[]drivenAppPortsRestful.ClientOption) {
	var header map[string][]string
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.BasicAuthenticationCredentialsOption:
			token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", o.Username(), o.Password())))
			header["Authorization"] = []string{fmt.Sprintf("Basic %s", token)}
		case drivenAppPortsRestful.AuthorizationSchemeAndTokenOption:
			header["Authorization"] = []string{fmt.Sprintf("%s %s", o.Scheme(), o.Token())}
		case drivenAppPortsRestful.ExtraHeaderLineOption:
			name := o.Name()
			_, existent := header[name]
			if !existent {
				header[name] = []string{}
			}
			header[name] = append(header[name], o.Value())
		case drivenAppPortsRestful.CustomHeaderOption:
			for n, l := range o.Map() {
				_, existent := header[n]
				if !existent {
					header[n] = []string{}
				}
				for _, v := range l {
					header[n] = append(header[n], v)
				}
			}
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return &header, &remainingOptions
}
