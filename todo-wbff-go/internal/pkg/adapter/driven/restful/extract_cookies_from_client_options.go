package driven_adapter_restful

func extractCookiesFromClientOptions(options *[]ClientOption) (*[]Cookie, *[]ClientOption) {
	cookies := make([]Cookie, 0)
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case CookieOption:
			cookies = append(cookies, o.Cookie())
		case CookiesOption:
			for _, c := range *o.Cookies() {
				cookies = append(cookies, c)
			}
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return &cookies, &remainingOptions
}
