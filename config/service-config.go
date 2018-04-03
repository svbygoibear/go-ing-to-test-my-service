package config

type Service struct {
	Port	 int64 	`json:"port"`
	Hostname string `json:"hostname"`
}

func Load(file string) (*Service, error) {
	conf := Service{ 8080, "127.0.0.1" }
	return &conf, nil
}
