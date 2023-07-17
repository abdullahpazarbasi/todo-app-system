package infrastructure_adapters_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

func extractCookiesFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (*[]drivenAppPortsRestful.Cookie, *[]drivenAppPortsRestful.ClientOption) {
	cookies := make([]drivenAppPortsRestful.Cookie, 0)
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.CookieOption:
			cookies = append(cookies, o.Cookie())
		case drivenAppPortsRestful.CookiesOption:
			for _, c := range *o.Cookies() {
				cookies = append(cookies, c)
			}
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return &cookies, &remainingOptions
}
