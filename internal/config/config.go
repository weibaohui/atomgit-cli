package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfg *viper.Viper

func init() {
	cfg = viper.New()

	cfg.SetConfigName("config")
	cfg.SetConfigType("toml")
	cfg.AddConfigPath(getConfigDir())
	cfg.Set("base_url", "https://api.atomgit.com")
	cfg.Set("token", "")

	_ = cfg.ReadInConfig()
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return fmt.Sprintf("%s/.config/atomgit-cli", home)
}

func GetBaseURL() string {
	if env := os.Getenv("ATOMGIT_BASE_URL"); env != "" {
		return strings.TrimSpace(env)
	}
	if baseURL := cfg.GetString("base_url"); baseURL != "" {
		return strings.TrimSpace(baseURL)
	}
	return "https://api.atomgit.com"
}

func GetToken() string {
	if env := os.Getenv("ATOMGIT_TOKEN"); env != "" {
		return strings.TrimSpace(env)
	}
	if token := cfg.GetString("token"); token != "" {
		return strings.TrimSpace(token)
	}
	return ""
}

func TokenSource() string {
	if env := os.Getenv("ATOMGIT_TOKEN"); env != "" && strings.TrimSpace(env) != "" {
		return "env"
	}
	if token := cfg.GetString("token"); token != "" && strings.TrimSpace(token) != "" {
		return "config"
	}
	return "none"
}

func SetToken(token string) error {
	cfg.Set("token", token)
	configPath := getConfigDir()
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}
	configFile := fmt.Sprintf("%s/config.toml", configPath)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return cfg.SafeWriteConfig()
	}
	return cfg.WriteConfig()
}

func DeleteToken() error {
	cfg.Set("token", "")
	configPath := getConfigDir()
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}
	configFile := fmt.Sprintf("%s/config.toml", configPath)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return cfg.SafeWriteConfig()
	}
	return cfg.WriteConfig()
}

func UserConfigHint() string {
	return fmt.Sprintf("Set a token with `amc auth login` or export ATOMGIT_TOKEN. Config home: %s/.config/atomgit-cli", os.Getenv("HOME"))
}

func MaskToken(token string) string {
	if token == "" {
		return "(none)"
	}
	if len(token) <= 8 {
		return strings.Repeat("*", len(token))
	}
	return token[:4] + "…" + token[len(token)-4:]
}
