/**
 * @Time: 2022/2/21 20:49
 * @Author: yt.yin
 */

package consul

// Consul 注册中心配置
type Consul struct {

	/** 注册中心地址 */
	Addr                  string     `mapstructure:"addr"                 json:"addr"                 yaml:"addr"`

	/** 间隔 单位秒 */
	RegisterInterval      int        `mapstructure:"register-interval"    json:"registerInterval"     yaml:"register-interval"`
}
