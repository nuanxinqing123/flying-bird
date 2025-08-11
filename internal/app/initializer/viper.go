package initializer

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nuanxinqing123/flying-bird/internal/app/config"
	"github.com/spf13/viper"
)

// Viper 初始化配置
func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("configs/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err = v.Unmarshal(&config.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&config.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
