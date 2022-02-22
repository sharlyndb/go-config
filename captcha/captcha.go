/**
 * @Time: 2022/2/22 17:46
 * @Author: yt.yin
 */

package captcha

// Captcha 验证码格式配置
type Captcha struct {

	/** 数字或字符串长度 */
	KeyLen             int          `mapstructure:"key-len"           json:"keyLen"         yaml:"key-len"`

	/** 验证码宽度 */
	ImgWidth           int          `mapstructure:"img-width"         json:"imgWidth"       yaml:"img-width"`

	/** 验证码高度 */
	ImgHeight          int          `mapstructure:"img-height"        json:"imgHeight"      yaml:"img-height"`

	/** 最大歪曲度 0.5-1.0 */
	MaxSkew            float64      `mapstructure:"max-skew"          json:"maxSkew"        yaml:"max-skew"`

	/** 分布的点的数量  推荐设置 100左右 */
	DotCount           int          `mapstructure:"dot-count"         json:"dotCount"       yaml:"dot-count"`
}