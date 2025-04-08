package initializers

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.SetConfigName("server")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
