package config

type ServiceConfig struct {
	BaseURL     string `yaml:"base_url"`
	HTTPMethod  string `yaml:"http_method"`
	ContentType string `yaml:"content_type"`
	XSI         string `yaml:"xsi"`
	XSD         string `yaml:"xsd"`
	Soap12      string `yaml:"soap12"`
	XMLName     string `yaml:"xml_name"`
	XMLNS       string `yaml:"xmlns"`
}
