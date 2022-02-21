/**
 * @Time: 2022/2/21 17:33
 * @Author: yt.yin
 */

package pay

// Alipay 支付宝 支付相关参数配置
type Alipay struct {

	/** 商户pid 即商户的账号id，在某些业务场景下要用到*/
	Pid            string          `mapstructure:"pid"               json:"pid"             yaml:"pid"`

	/** 应用id */
	AppId          string          `mapstructure:"app-id"            json:"appId"           yaml:"app-id"`

	/** 私钥 */
	PriKey         string          `mapstructure:"pri-key"           json:"priKey"          yaml:"pri-key"`

	/** 公钥 主要是回调 验签用 */
	PubKey         string          `mapstructure:"pub-key"           json:"pubKey"          yaml:"pub-key"`

	/** 签名方式  商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，私钥 1024位RSA */
	SignType       string          `mapstructure:"sign-type"         json:"signType"        yaml:"sign-type"`

	/** 支付宝回调的url */
	NotifyUrl      string          `mapstructure:"notify-url"        json:"notifyUrl"       yaml:"notify-url"`

	/** 默认订单标题 */
	Subject        string          `mapstructure:"subject"           json:"subject"         yaml:"subject"`

}
