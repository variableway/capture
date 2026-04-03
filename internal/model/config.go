package model

type Config struct {
	App      AppConfig      `yaml:"app"`
	Defaults DefaultsConfig `yaml:"defaults"`
	Feishu   FeishuConfig   `yaml:"feishu"`
	Bitable  BitableConfig  `yaml:"bitable"`
	Bot      BotConfig      `yaml:"bot"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	DataDir string `yaml:"data_dir"`
}

type DefaultsConfig struct {
	Priority string   `yaml:"priority"`
	Tags     []string `yaml:"tags"`
	Editor   string   `yaml:"editor"`
}

type FeishuConfig struct {
	AppID             string `yaml:"app_id"`
	AppSecret         string `yaml:"app_secret"`
	EncryptKey        string `yaml:"encrypt_key"`
	VerificationToken string `yaml:"verification_token"`
}

type BitableConfig struct {
	Enabled bool   `yaml:"enabled"`
	AppToken string `yaml:"app_token"`
	TableID  string `yaml:"table_id"`
}

type BotConfig struct {
	Mode string `yaml:"mode"` // webhook, websocket
	Port int    `yaml:"port"`
}

func DefaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:    "Capture",
			Version: "0.1.0",
			DataDir: "~/.capture",
		},
		Defaults: DefaultsConfig{
			Priority: "medium",
			Editor:   "vim",
		},
		Feishu: FeishuConfig{},
		Bitable: BitableConfig{
			Enabled: false,
		},
		Bot: BotConfig{
			Mode: "websocket",
			Port: 8080,
		},
	}
}
