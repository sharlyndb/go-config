/**
 * @Time: 2022/2/21 18:08
 * @Author: yt.yin
 */

package env


// Profiles 多配置文件指定配置
type Profiles struct{

	/** 指定的配置文件
	dev Development environment，开发环境，用于开发者调试使用
	fat Feature Acceptance Test environment，功能验收测试环境，用于软件测试者测试使用
	uat User Acceptance Test environment，用户验收测试环境，用于生产环境下的软件测试者测试使用
	pro Production environment，生产环境
	*/
	Active         string       `mapstructure:"active"              json:"active"                  yaml:"active"`
}
