package components

type Security struct {
	APIKey *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=noresponse_api_key"`
}

func (s *Security) GetAPIKey() *string {
	if s == nil {
		return nil
	}
	return s.APIKey
}
