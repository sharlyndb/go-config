/**
 * @Time: 2022/2/21 03:14
 * @Author: yt.yin
 */

package mqtt

// Mqtt 配置文件
type Mqtt struct {

	/** Mqtt代理服务器的Ip和端口 */
	Url                   string       `mapstructure:"url"                      json:"url"                       yaml:"url"`

	/** 连接到代理服务器的用户名 */
	Username              string       `mapstructure:"username"                 json:"username"                  yaml:"username"`

	/** 密码 */
	Password              string       `mapstructure:"password"                 json:"password"                  yaml:"password"`

	/** Mqtt 协议版本号 4是3.1.1，3是3.1 */
	ProtocolVersion       uint         `mapstructure:"protocol-ver"             json:"protocolVer"               yaml:"protocol-ver"`

	/** 设置客户端掉线服务端是否清除session */
	CleanSession          bool         `mapstructure:"clean-session"            json:"cleanSession"              yaml:"clean-session"`

	/** 断开后是否重新连接 */
	AutoReconnect         bool         `mapstructure:"auto-reconnect"           json:"autoReconnect"             yaml:"auto-reconnect"`

	/** 保活时间间隔 */
	KeepAlive             int          `mapstructure:"keep-alive"               json:"keepAlive"                 yaml:"keep-alive"`

	/** 最大连接间隔时间 单位：秒 */
	MaxReconnectInterval  int          `mapstructure:"max-reconnect-interval"   json:"maxReconnectInterval"      yaml:"max-reconnect-interval"`

	/** ping 超时时间 单位：秒 */
	PingTimeout           int          `mapstructure:"pingTimeout"              json:"ping-timeout"              yaml:"ping-timeout"`

	/** 写超时时间 单位：秒 */
	WriteTimeout          int          `mapstructure:"write-timeout"            json:"writeTimeout"              yaml:"write-timeout"`

	/** 连接超时时间 单位：秒 */
	ConnectTimeout        int          `mapstructure:"connect-timeout"          json:"connectTimeout"            yaml:"connect-timeout"`

	/** 遗言发送的topic */
	WillTopic             string       `mapstructure:"will-topic"               json:"willTopic"                 yaml:"will-topic"`
}

