/**
 * @Time: 2022/2/22 11:44
 * @Author: yt.yin
 */

package env

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// 设置和获取程序运行环境
var (

	/** 当前激活的环境 */
	active Environment

	/** dev Development environment，开发环境，用于开发者调试使用 */
	dev Environment = &environment{value: "dev"}

	/** fat Feature Acceptance Test environment，功能验收测试环境，用于软件测试者测试使用 */
	fat Environment = &environment{value: "fat"}

	/** uat User Acceptance Test environment，用户验收测试环境，用于生产环境下的软件测试者测试使用 */
	uat Environment = &environment{value: "uat"}

	/** pro Production environment，生产环境 */
	pro Environment = &environment{value: "pro"}
)

var _ Environment = (*environment)(nil)

const (
	ConfigFile = "./resources/active.yaml"
)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsUat() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsFat() bool {
	return e.value == "fat"
}

func (e *environment) IsUat() bool {
	return e.value == "uat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

// Active 当前配置的env
func Active() Environment {
	if active == nil || active.Value() == "" {
		activeFile := LoadActiveFile()
		if activeFile != nil {
			env := activeFile.Active
			switch strings.ToLower(strings.TrimSpace(env)) {
			case "dev":
				active = dev
			case "fat":
				active = fat
			case "uat":
				active = uat
			case "pro":
				active = pro
			default:
				active = fat
				log.Println("找不到对应的环境，默认使用 fat")
			}
			return active
		}
	}
	return active
}

// SetActive 配置当前环境
func SetActive(env string) string {
	switch strings.ToLower(strings.TrimSpace(env)) {
	case "dev":
		active = dev
	case "fat":
		active = fat
	case "uat":
		active = uat
	case "pro":
		active = pro
	default:
		active = fat
		log.Println("找不到对应的环境，默认使用 fat")
	}
	return env
}

// LoadActiveFile 读取默认配置文件 active.yaml
func LoadActiveFile(path ...string) *Profiles {
	v := viper.New()
	if len(path) == 0 {
		v.SetConfigFile(ConfigFile)
		log.Println("读取默认配置文件:", ConfigFile)
	} else {
		v.SetConfigFile(path[0])
		log.Println("读取指定配置文件:", path[0])
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("严重错误的配置文件 : %s \n", err))
	}
	var p Profiles
	if err := v.Unmarshal(&p); err != nil {
		log.Println("读取active配置文件异常:", err)
		return nil
	}
	return &p
}
