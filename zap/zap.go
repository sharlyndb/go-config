/**
 * @Time: 2022/2/21 23:44
 * @Author: yt.yin
 */

package zap

// Zap 日志配置
type Zap struct {

	/** 日志级别 */
	Level                string       `mapstructure:"level"               json:"level"             yaml:"level"`

	/** 日志格式 */
	Format               string       `mapstructure:"format"              json:"format"            yaml:"format"`

	/** 日志前缀 */
	Prefix               string       `mapstructure:"prefix"              json:"prefix"            yaml:"prefix"`

	/** 日志目录 */
	Director             string       `mapstructure:"director"            json:"director"          yaml:"director"`

	/** 日志最大保留时间 单位：天 */
	MaxAge               int          `mapstructure:"max-age"             json:"maxAge"            yaml:"max-age"`

	/** 日志软连接文件 */
	LinkName             string       `mapstructure:"link-name"           json:"linkName"          yaml:"link-name"`

	/** 是否在日志中输出源码所在的行 */
	ShowLine             bool         `mapstructure:"show-line"           json:"showLine"          yaml:"showLine"`

	/** 日志编码等级，指定不通过等级可以有不同颜色 */
	EncodeLevel          string       `mapstructure:"encode-level"        json:"encodeLevel"       yaml:"encode-level"`

	/** 堆栈捕捉标识 */
	StacktraceKey        string       `mapstructure:"stacktrace-key"      json:"stacktraceKey"     yaml:"stacktrace-key"`

	/** 是否在控制台打印日志 */
	LogInConsole         bool         `mapstructure:"log-in-console"      json:"logInConsole"      yaml:"log-in-console"`
}
