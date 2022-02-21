/**
 * @Time: 2022/2/21 15:47
 * @Author: yt.yin
 */

package redis

type Redis struct {

	/** redis 数据服务器ip和端口 */
	Addr             string       `mapstructure:"addr"              json:"addr"             yaml:"addr"`

	/** 指定连接的数据库 默认连数据库0 */
	DB               int          `mapstructure:"db"                json:"db"               yaml:"db"`

	/** 连接密码 */
	Password         string       `mapstructure:"password"          json:"password"         yaml:"password"`

	/** 最大重试次数 */
	MaxRetries       int          `mapstructure:"max-retries"       json:"maxRetries"       yaml:"max-retries"`

	/** 最大空闲连接数 */
	MinIdleConns     int          `mapstructure:"min-idle-conns"    json:"minIdleConns"     yaml:"min-idle-conns"`

	/** 连接池大小 */
	PoolSize         int          `mapstructure:"pool-size"         json:"poolSize"         yaml:"pool-size"`
}
