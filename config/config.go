package config

import (
	"errors"
	"fmt"

	"github.com/odpf/salt/config"
	"github.com/odpf/siren/domain"
)

// LoadConfig returns application configuration
func LoadConfig(configFile string) (*domain.Config, error) {
	var cfg domain.Config
	loader := config.NewLoader(config.WithFile(configFile))

	if err := loader.Load(&cfg); err != nil {
		if errors.As(err, &config.ConfigFileNotFoundError{}) {
			fmt.Println(err)
			return &cfg, nil
		}
		return nil, err
	}
	return &cfg, nil
}

// func LoadConfig() *domain.Config {
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath("./")
// 	viper.AddConfigPath("../")
// 	viper.SetConfigType("yaml")
// 	viper.AutomaticEnv()
// 	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 			fmt.Println("config file was not found. Env vars and defaults will be used")
// 		} else {
// 			panic(err)
// 		}
// 	}

// 	err, configKeys := getFlattenedStructKeys(domain.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Bind each conf fields to environment vars
// 	for key := range configKeys {
// 		err := viper.BindEnv(configKeys[key])
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	var config domain.Config
// 	defaults.SetDefaults(&config)

// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		fmt.Errorf("unable to unmarshal config to struct: %v\n", err)
// 	}
// 	return &config
// }

// func getFlattenedStructKeys(config domain.Config) (error, []string) {
// 	var structMap map[string]interface{}
// 	err := mapstructure.Decode(config, &structMap)
// 	if err != nil {
// 		return err, nil
// 	}

// 	flat, err := flatten.Flatten(structMap, "", flatten.DotStyle)
// 	if err != nil {
// 		return err, nil
// 	}

// 	keys := make([]string, 0, len(flat))
// 	for k := range flat {
// 		keys = append(keys, k)
// 	}

// 	return nil, keys
// }
