package ftp

type Ftp struct {

	/** ftp 服务器Ip和端口 */
	Addr            string       `mapstructure:"addr"                json:"addr"              yaml:"addr"`

	/** 用户 */
	Username        string       `mapstructure:"username"            json:"username"          yaml:"username"`

	/** 密码 */
	Password        string       `mapstructure:"password"            json:"password"          yaml:"password"`

	/** 指定目录 */
	Cwd             string       `mapstructure:"cwd"                 json:"cwd"               yaml:"cwd"`

}
