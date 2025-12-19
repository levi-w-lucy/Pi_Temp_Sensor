package model

type Secrets struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TemperaturePayload struct {
	Temperature    float64 `json:"temperature"`
	DatasourceName string  `json:"datasource_name"`
	ClientID       string  `json:"client_id"`
	ClientSecret   string  `json:"client_secret"`
}

type CmdArgs struct {
	ConfigFilePath string
	ReadInterval   int
}
