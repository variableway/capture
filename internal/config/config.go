package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/variableway/innate/capture/internal/model"
)

func DefaultDataDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".capture")
}

func Load(dataDir string) (*model.Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(dataDir)

	cfg := model.DefaultConfig()

	if err := v.ReadInConfig(); err != nil {
		// Config file doesn't exist yet, return defaults
		return cfg, nil
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func Save(dataDir string, cfg *model.Config) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(dataDir)

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(dataDir, "config.yaml")
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("# Capture Configuration\n")
	if err != nil {
		return err
	}

	// Simple YAML writing via viper
	v.Set("app", cfg.App)
	v.Set("defaults", cfg.Defaults)
	v.Set("feishu", cfg.Feishu)
	v.Set("bitable", cfg.Bitable)
	v.Set("bot", cfg.Bot)

	return v.WriteConfigAs(configPath)
}
