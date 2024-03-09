package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION"`

	MySQLConfig    MySQLConfig        `mapstructure:"DATABASE"`
	RedisConfig    RedisConfig        `mapstructure:"REDIS"`
	RabbitMQConfig RabbitMQConfig     `mapstructure:"RABBITMQ"`
	BankAcquirer   []ThirdPartyConfig `mapstructure:"BANK_ACQUIRER"`
	SnapCoreURL    string             `mapstructure:"SNAP_CORE_URL"`
}

type ThirdPartyConfig struct {
	Name    string `mapstructure:"NAME"`
	BaseUrl string `mapstructure:"BASE_URL"`
}

type Secret struct {
	MySQLSecret    MySQLSecret    `mapstructure:"DATABASE"`
	RedisSecret    RedisSecret    `mapstructure:"REDIS"`
	RabbitMQSecret RabbitMQSecret `mapstructure:"RABBITMQ"`
	SecuritySecret SecuritySecret `mapstructure:"SECURITY"`

	NewRelicLicenseKey string             `mapstructure:"NEW_RELIC_LICENSE_KEY"`
	InternalServiceKey string             `mapstructure:"INTERNAL_SERVICE_KEY"`
	BankAcquirer       []ThirdPartySecret `mapstructure:"BANK_ACQUIRER"`
	SnapCoreKey        string             `mapstructure:"SNAP_CORE_KEY"`
	SnapCoreSecret     string             `mapstructure:"SNAP_CORE_SECRET"`
}

type ThirdPartySecret struct {
	Name       string `mapstructure:"NAME"`
	ClientID   string `mapstructure:"CLIENT_ID"`
	ClientKey  string `mapstructure:"CLIENT_KEY"`
	PublicKey  string `mapstructure:"SECRET_KEY"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`
}

type MySQLConfig struct {
	Dialect      string `mapstructure:"DIALECT"`
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	MaxIdleConns int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleTime  int    `mapstructure:"MAX_IDLE_TIME"`
	MaxLifeTime  int    `mapstructure:"MAX_LIFE_TIME"`
}

type MySQLSecret struct {
	Database string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type RedisConfig struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	CacheDB int    `mapstructure:"CACHE_DB"`
}

type RedisSecret struct {
	Password string `mapstructure:"PASSWORD"`
}

type RabbitMQConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type RabbitMQSecret struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type SecuritySecret struct {
	JwtSecretKey  string         `mapstructure:"JWT_SECRET_KEY"`
	ChannelApiKey []ClientApiKey `mapstructure:"CHANNEL_API_KEY"`
}

type ClientApiKey struct {
	ClientName   string `mapstructure:"CLIENT_NAME"`
	ClientID     string `mapstructure:"CLIENT_ID"`
	ClientSecret string `mapstructure:"CLIENT_SECRET"`
}

func LoadConfig(configPath, secretPath string) (*Config, *Secret, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	viper.Reset()

	// Load Secret
	viper.SetConfigFile(secretPath)
	var secret Secret
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading secret file: %w", err)
	}
	if err := viper.Unmarshal(&secret); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, &secret, nil
}

func (c *SecuritySecret) GetClientApi(key string) *ClientApiKey {
	for _, v := range c.ChannelApiKey {
		if v.ClientID == key {
			return &v
		}
	}
	return nil
}

func (c *SecuritySecret) GetClientApiByName(clientName string) *ClientApiKey {
	for _, v := range c.ChannelApiKey {
		if v.ClientName == clientName {
			return &v
		}
	}
	return nil
}

func (c *Config) GetBankAcquirer(key string) *ThirdPartyConfig {
	for _, v := range c.BankAcquirer {
		if v.Name == key {
			return &v
		}
	}

	return nil
}

func (c *Secret) GetBankAcquirerSecret(key string) *ThirdPartySecret {
	for _, v := range c.BankAcquirer {
		if v.Name == key {
			return &v
		}
	}
	return nil
}
