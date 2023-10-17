package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

const (
	externalAppConfig = "/etc/config/app/"
	embeddedAppConfig = "config/"
	name              = "app_config"
)

type (
	AppConf struct {
		DataSource DataSource `mapstructure:"DATA_SOURCE"`
		App        App        `mapstructure:"APP"`
	}

	App struct {
		Name string `mapstructure:" NAME"`
	}

	DataSource struct {
		HostPort string `mapstructure:"PORT"`
		Url      string `mapstructure:"URL"`
	}
)

var Application AppConf

func LoadAppConfig() error {
	if _, err := os.Stat(fmt.Sprintf("%v%v.yaml", externalAppConfig, name)); errors.Is(err, os.ErrNotExist) {
		log.Warn().Msg("no environment config found, using default config")
		viper.AddConfigPath(embeddedAppConfig)
	} else {
		log.Info().Msg("using environment config")
		viper.AddConfigPath(externalAppConfig)
	}
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading in config: %v", err.Error())
	}

	err = viper.Unmarshal(&Application)
	if err != nil {
		return fmt.Errorf("error unmarshal config: %v", err.Error())
	}

	return nil
}
