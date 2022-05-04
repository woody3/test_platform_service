package utils

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *viper.Viper
var lock sync.Mutex

func Viper(filename string) {
	config = viper.New()
	config.AddConfigPath("../")
	config.AddConfigPath("../config/")
	config.SetConfigFile(filename)
	config.SetConfigType(strings.Split(filename, ".")[1])
	readConfig(config)
	config.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
}

func GetConfig() *viper.Viper {
	return config
}

func readConfig(v *viper.Viper) {
	lock.Lock()
	defer lock.Unlock()
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
}
