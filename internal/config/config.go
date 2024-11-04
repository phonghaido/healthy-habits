package config

import "github.com/spf13/viper"

type USDAConfig struct {
	URL    string `json:"usdaURL"`
	APIKey string `json:"usdaAPIKey"`
}

type MongoDBConfig struct {
	MongoDBUsername string `json:"mongodbUsername"`
	MongoDBPassword string `json:"mongodbPassword"`
	MongoDBHost     string `json:"mongodbHost"`
	MongoDBPort     string `json:"mongodbPort"`
	MongoDBDatabase string `json:"mongodbDatabase"`
}

func SetupViper() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetUSDAConfig() (USDAConfig, error) {
	return USDAConfig{
		URL:    viper.GetString("USDA_URL"),
		APIKey: viper.GetString("USDA_API_KEY"),
	}, nil
}

func GetMongoDBConfig() (MongoDBConfig, error) {
	return MongoDBConfig{
		MongoDBUsername: viper.GetString("MONGODB_USERNAME"),
		MongoDBPassword: viper.GetString("MONGODB_PASSWORD"),
		MongoDBHost:     viper.GetString("MONGODB_HOST"),
		MongoDBPort:     viper.GetString("MONGODB_PORT"),
		MongoDBDatabase: viper.GetString("MONGODB_DATABASE"),
	}, nil
}
