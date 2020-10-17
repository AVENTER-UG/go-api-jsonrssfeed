package types

type Config struct {
	APIPort           string
	APIBind           string
	LogLevel          string
	MinVersion        string
	AppName           string
	EnableSyslog      bool
}

type GetFeed struct {
	Url    string `json:"url"`
}

