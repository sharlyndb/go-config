/**
 * @Time: 2022/2/21 16:07
 * @Author: yt.yin
 */

package database

// PostgreSQL 配置
type PostgreSQL struct {

	/** 数据库IP地址 */
	Host               string       `mapstructure:"host"                json:"host"              yaml:"host"`

	/** 端口 */
	Port               string       `mapstructure:"port"                json:"port"              yaml:"port"`

	/** 后缀配置  默认配置 charset=utf8mb4&parseTime=True&loc=Local */
	Config             string       `mapstructure:"config"              json:"config"            yaml:"config"`

	/** 数据库名称 */
	Dbname             string       `mapstructure:"db-name"             json:"dbname"            yaml:"db-name"`

	/** 数据库用户名 */
	Username           string       `mapstructure:"username"            json:"username"          yaml:"username"`

	/** 数据库密码 */
	Password           string       `mapstructure:"password"            json:"password"          yaml:"password"`

	/** 最大空闲连接数 */
	MaxIdleConns       int          `mapstructure:"max-idle-conns"      json:"maxIdleConns"      yaml:"max-idle-conns"`

	/** 最大连接数 */
	MaxOpenConns       int          `mapstructure:"max-open-conns"      json:"maxOpenConns"      yaml:"max-open-conns"`

}
