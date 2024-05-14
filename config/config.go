package config

type HarborIP struct {
	HarborIP string `yaml:"harborIP"`

	Auth struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Project []string
}
