package mqtt

// Mqtt 配置文件
type Mqtt struct {

	/** Mqtt代理服务器的Ip和端口 */
	Url              string       `mapstructure:"url"            json:"url"            yaml:"url"`

	/** 连接到代理服务器的用户名 */
	Username         string       `mapstructure:"username"       json:"username"       yaml:"username"`

	/** 密码 */
	Password         string       `mapstructure:"password"       json:"password"       yaml:"password"`


}

