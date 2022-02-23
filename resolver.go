/**
 * @Time: 2022/2/23 10:02
 * @Author: yt.yin
 */

package configs

import (
	"github.com/goworkeryyt/configs/env"
	"github.com/spf13/viper"
	"log"
)

var (
	// 全局配置
	globalConfigs  *Configs
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
func SubItem(subKey string,v interface{}){
	if globalConfigs == nil || globalConfigs.Viper == nil {
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
		if err := v.Unmarshal(globalConfigs); err != nil {
			log.Fatalf("读取配置文件异常 : %s \n", err)
			return
		}
		if globalConfigs != nil &&  globalConfigs.Viper != nil {
			globalConfigs.Viper = v
		}
		// 获取子项的值
		value := v.Sub(subKey)
		if value == nil {
			return
		}
		err = value.Unmarshal(v)
		if err != nil{
			log.Fatalf("读取子配置项异常 : %s \n", err)
		}
		return
	}
	value := globalConfigs.Viper.Sub(subKey)
	if value == nil {
		return
	}
	err := value.Unmarshal(v)
	if err != nil{
		log.Println("读取子配置项异常:", err)
	}
	return
}