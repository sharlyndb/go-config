/**
 * @Time: 2022/2/22 02:03
 * @Author: yt.yin
 */

package configs

import (
	"github.com/goworkeryyt/configs/captcha"
	"github.com/goworkeryyt/configs/consul"
	"github.com/goworkeryyt/configs/database"
	"github.com/goworkeryyt/configs/email"
	"github.com/goworkeryyt/configs/ftp"
	"github.com/goworkeryyt/configs/jwt"
	"github.com/goworkeryyt/configs/mqtt"
	"github.com/goworkeryyt/configs/pay"
	"github.com/goworkeryyt/configs/redis"
	"github.com/goworkeryyt/configs/server"
	"github.com/goworkeryyt/configs/zap"
	"github.com/spf13/viper"
)

// Configs 项目单节点统一配置
type Configs struct {

	/** 服务实例配置 */
	Server            server.Server          `mapstructure:"server"             json:"server"            yaml:"server"`

	/** 注册中心配置 */
	Consul            consul.Consul          `mapstructure:"consul"             json:"consul"            yaml:"consul"`

	/** 验证码配置 */
	Captcha           captcha.Captcha        `mapstructure:"captcha"            json:"captcha"           yaml:"captcha"`

	/** MySQL 数据库配置 */
	MySQL             database.MySQL         `mapstructure:"mysql"              json:"mysql"             yaml:"mysql"`

	/** Postgre 数据库配置 */
	PostgreSQL        database.PostgreSQL    `mapstructure:"postgre"            json:"postgre"           yaml:"postgre"`

	/** sqlite 数据库配置 */
	SQLite            database.SQLite        `mapstructure:"sqlite"             json:"sqlite"            yaml:"sqlite"`

	/** redis 缓存数据库配置 */
	Redis             redis.Redis            `mapstructure:"redis"              json:"redis"             yaml:"redis"`

	/** 邮件发送相关配置 */
	Email             email.Email            `mapstructure:"email"              json:"email"             yaml:"email"`

	/** 文件服务器配置 */
	Ftp               ftp.Ftp                `mapstructure:"ftp"                json:"local"             yaml:"ftp"`

	/** jwt token 相关配置 */
	JWT               jwt.JWT                `mapstructure:"jwt"                json:"jwt"               yaml:"jwt"`

	/** mqtt 代理服务器相关配置 */
	Mqtt              mqtt.Mqtt              `mapstructure:"mqtt"               json:"mqtt"              yaml:"mqtt"`

	/** 日志相关配置 */
	Zap               zap.Zap                `mapstructure:"zap"                json:"zap"               yaml:"zap"`

	/** 支付相关配置-支付宝 */
	AliPay            pay.Alipay             `mapstructure:"alipay"             json:"alipay"            yaml:"alipay"`

	/** 支付相关配置-微信 */
	Weichat           pay.Weichat            `mapstructure:"weichat"            json:"weichat"           yaml:"weichat"`

	/** 可自己取一些扩展配置 */
	Viper             *viper.Viper
}
