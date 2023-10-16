package driven_adapter_restful

import (
	"encoding/base64"
	"fmt"
)

func extractHeaderFromClientOptions(options *[]ClientOption) (*map[string][]string, *[]ClientOption) {
	var header map[string][]string
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case BasicAuthenticationCredentialsOption:
			token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", o.Username(), o.Password())))
			header["Authorization"] = []string{fmt.Sprintf("Basic %s", token)}
		case AuthorizationSchemeAndTokenOption:
			header["Authorization"] = []string{fmt.Sprintf("%s %s", o.Scheme(), o.Token())}
		case ExtraHeaderLineOption:
			name := o.Name()
			_, existent := header[name]
			if !existent {
				header[name] = []string{}
			}
			header[name] = append(header[name], o.Value())
		case CustomHeaderOption:
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
