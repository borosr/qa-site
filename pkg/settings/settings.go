package settings

import (
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	DBConnectionString string `json:"db_connection_string"`
	Port               string `json:"port"`
	JwtHMAC            string `json:"jwt_hmac"`
	JwtSecret          string `json:"jwt_secret"`
	RevokeHMAC         string `json:"revoke_hmac"`
	RevokeSecret       string `json:"revoke_secret"`
	Hostname           string `json:"hostname"`
	BadgerPath         string `json:"badger_path"`
	GithubClientID     string `json:"github_client_id"`
	GithubClientSecret string `json:"github_client_secret"`
	set                bool   `json:"-"`
	Visibility         string `json:"visibility"`
}

func Get() Config {
	if !config.set {
		config = setup()
	}
	return config
}

func setup() Config {
	c := Config{}
	setDefaults()
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	c.set = true

	return c
}

func setDefaults() {
	viper.SetDefault("DBConnectionString", "postgres://qa_site:qa_site@localhost:26257/qm_site?sslmode=disable")
	viper.SetDefault("Port", "8000")
	viper.SetDefault("GithubClientID", "")
	viper.SetDefault("GithubClientSecret", "")
	viper.SetDefault("BadgerPath", "/tmp/badger")
}
