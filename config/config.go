package config

import (
	"encoding/json"
	"os"
)

type SDKConfig struct {
	ProjectName string     `json:"projectName"`
	MetricsPort string     `json:"metricsPort"`
	AlertEmail  EmailAlert `json:"alertEmail"`
}

type EmailAlert struct {
	Enable   bool     `json:"enable"`
	SMTPHost string   `json:"smtpHost"`
	SMTPPort string   `json:"smtpPort"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	From     string   `json:"from"`
	To       []string `json:"to"`
}

var Cfg SDKConfig

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&Cfg)
}
