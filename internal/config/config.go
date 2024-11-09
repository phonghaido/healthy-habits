package config

import "github.com/spf13/viper"

type CommonConfig struct {
	Port string
}
type USDAConfig struct {
	URL    string `json:"usdaURL"`
	APIKey string `json:"usdaAPIKey"`
}

type MongoDBConfig struct {
	MongoDBConnStr string `json:"mongodbConnStr"`
}

func SetupViper() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetCommonConfig() CommonConfig {
	return CommonConfig{
		Port: viper.GetString("APP_PORT"),
	}
}
func GetUSDAConfig() (USDAConfig, error) {
	return USDAConfig{
		URL:    viper.GetString("USDA_URL"),
		APIKey: viper.GetString("USDA_API_KEY"),
	}, nil
}

func GetMongoDBConfig() (MongoDBConfig, error) {
	return MongoDBConfig{
		MongoDBConnStr: viper.GetString("MONGODB_CONN_STR"),
	}, nil
}
