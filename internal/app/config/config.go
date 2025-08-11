package config

import (
	"github.com/nuanxinqing123/flying-bird/internal/app/config/autoload"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Configuration struct {
	App autoload.App `mapstructure:"app" json:"app" yaml:"app"`
}

var (
	Config Configuration
	Log    *zap.Logger
	VP     *viper.Viper
)
