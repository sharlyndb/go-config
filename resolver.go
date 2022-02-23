/**
 * @Time: 2022/2/23 10:02
 * @Author: yt.yin
 */

package goconfig

import (
	"github.com/fsnotify/fsnotify"
	"github.com/goworkeryyt/go-config/env"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	// 全局配置
	globalConfig *Config
	// 为该全局变量创建一个锁
	glcMutex sync.Mutex
)

const (

	// ConfigSuffix 配置文件默认后缀
	ConfigSuffix = "_config"

	// ConfigType 配置文件类型
	ConfigType = "yaml"

	// ConfigPath 配置文件所在路径
	ConfigPath = "./resources"
)

// SubItem 从配置中获取指定的配置子项
func (c *Config) SubItem(subKey string, v interface{}) {
	glcMutex.Lock()
	defer glcMutex.Unlock()
	if globalConfig == nil || globalConfig.Viper == nil {
		active := env.Active()
		filename := active.Value() + ConfigSuffix
		v := viper.New()
		v.SetConfigName(filename)
		v.SetConfigType(ConfigType)
		v.AddConfigPath(ConfigPath)
		log.Println("读取配置文件:", filename)
		err := v.ReadInConfig()
		if err != nil {
			log.Fatalf("读取配置文件异常 : %s \n", err)
			return
		}
		globalConfig = &Config{}
		if err := v.Unmarshal(globalConfig); err != nil {
			log.Fatalf("读取配置文件异常 : %s \n", err)
			return
		}
		if globalConfig != nil && globalConfig.Viper != nil {
			globalConfig.Viper = v
		}
		// 获取子项的值
		value := v.Sub(subKey)
		if value == nil {
			return
		}
		err = value.Unmarshal(v)
		if err != nil {
			log.Fatalf("读取子配置项异常 : %s \n", err)
		}
		return
	}
	value := globalConfig.Viper.Sub(subKey)
	if value == nil {
		return
	}
	err := value.Unmarshal(v)
	if err != nil {
		log.Println("读取子配置项异常:", err)
	}
	return
}

// GlobalConfig 获取全局配置
func GlobalConfig(envArr ...string) *Config {
	glcMutex.Lock()
	defer glcMutex.Unlock()
	if globalConfig != nil {
		return globalConfig
	}
	v := viper.New()
	filename := ""
	if len(envArr) == 0 {
		active := env.Active()
		filename = active.Value() + ConfigSuffix
	} else {
		filename = envArr[0] + ConfigSuffix
	}
	v.SetConfigName(filename)
	v.SetConfigType(ConfigType)
	v.AddConfigPath(ConfigPath)
	log.Println("读取配置文件:", filename)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件异常 : %s \n", err)
		return nil
	}
	globalConfig = &Config{}
	if err := v.Unmarshal(globalConfig); err != nil {
		log.Fatalf("读取配置文件异常 : %s \n", err)
		return nil
	}
	globalConfig.Viper = v
	// 监控配置改变
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件内容发生改变:", e.Name)
		if err := v.Unmarshal(globalConfig); err != nil {
			log.Fatalf("读取配置文件异常 : %s \n", err)
			return
		}
		globalConfig.Viper = v
	})
	return globalConfig
}
