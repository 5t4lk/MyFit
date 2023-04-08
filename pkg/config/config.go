package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken  string
	TelegramBotURL string `mapstructure:"bot_url"`
	APIKey         string
	APIHost        string
	DBuri          string

	Messages Messages
}

type Messages struct {
	Responses
}

type Responses struct {
	UnknownMessage    string `mapstructure:"unknown_message"`
	DisconnectConsult string `mapstructure:"disconnect_consult"`
	UnknownCommand    string `mapstructure:"unknown_command"`
	ConnectConsult    string `mapstructure:"connect_consult"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TOKEN", "6164310826:AAGweHp6I4UzrNmBfy6krwbjJCg1kkIA5i0")
	os.Setenv("API_KEY", "85e69f9465msh9b8963f3d02fc11p1fbbc9jsn64207893b164")
	os.Setenv("API_HOST", "chatgpt-ai-chat-bot.p.rapidapi.com")
	os.Setenv("DB_URI", "mongodb://localhost:27017")

	if err := viper.BindEnv("api_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("api_host"); err != nil {
		return err
	}
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("db_uri"); err != nil {
		return err
	}

	cfg.APIKey = viper.GetString("api_key")
	cfg.APIHost = viper.GetString("api_host")
	cfg.TelegramToken = viper.GetString("token")
	cfg.DBuri = viper.GetString("db_uri")

	return nil
}
