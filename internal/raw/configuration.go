package raw

type Configuration struct {
	Images *ConfigurationImages `json:"images"`
}

func (c *Configuration) SetDefaults() {}

type ConfigurationImages struct {
	SecureBaseUrl string `json:"secure_base_url"`
}
