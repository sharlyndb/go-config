/**
 * @Time: 2022/2/21 21:38
 * @Author: yt.yin
 */

package server

type Server struct {

	/** 端口 */
	Addr                      int         `mapstructure:"addr"                        json:"addr"                     yaml:"addr"`

	/** 服务名称 */
	ServiceName               string      `mapstructure:"service-name"                json:"serviceName"              yaml:"service-name"`

	/** 请求根路径 */
	ContextPath               string      `mapstructure:"context-path"                json:"contextPath"              yaml:"context-path"`

	/** 数据库类型 */
	DataDriver                string      `mapstructure:"data-driver"                 json:"dataDriver"               yaml:"data-driver"`

	/** 日志打印级别 */
	LogLevel                  string      `mapstructure:"log-level"                   json:"logLevel"                 yaml:"log-level"`

	/** 是否开启请求方式检测 */
	HandleMethodNotAllowed    bool        `mapstructure:"handle-method-not-allowed"   json:"handleMethodNotAllowed"   yaml:"handle-method-not-allowed"`

	/** 语言    */
	Language                  string      `mapstructure:"language"                    json:"language"                 yaml:"language"`
}
