/**
 * @Time: 2022/2/21 21:38
 * @Author: yt.yin
 */

package server

type Server struct {

	/** 端口 */
	Addr                      int         `mapstructure:"addr"                        json:"addr"                     yaml:"addr"`

	/** 服务名称 */
	ServerName                string      `mapstructure:"server-name"                 json:"serverName"               yaml:"server-name"`

	/** 请求根路径 */
	ContextPath               string      `mapstructure:"context-path"                json:"contextPath"              yaml:"context-path"`

	/** 数据库类型 */
	DataDriver                string      `mapstructure:"data-driver"                 json:"dataDriver"               yaml:"data-driver"`

	/** 是否开启请求方式检测 */
	HandleMethodNotAllowed    bool        `mapstructure:"handle-method-not-allowed"   json:"handleMethodNotAllowed"   yaml:"handle-method-not-allowed"`

	/** 语言    */
	Language                  string      `mapstructure:"language"                    json:"language"                 yaml:"language"`
}
