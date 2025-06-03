package raw

type ApiStatus struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_message"`
}

func (s *ApiStatus) SetDefaults() {}
