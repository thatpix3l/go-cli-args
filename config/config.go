package config

type Config struct {
	UseColor  *bool `json:"use_color,omit_empty"`
	ShowFunny *bool `json:"show_funny,omit_empty"`
}
