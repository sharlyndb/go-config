# configs 
### 介绍
go 开发中常用的一些配置
### 安装
```
go get -u github.com/goworkeryyt/configs
```
### 例子
默认的程序根目录下必须包含 resources 文件夹，且文件夹内必须有 active.yaml和四种不同环境的开发文件至少一种
```go
package main

import (
	"fmt"
	"github.com/goworkeryyt/configs"
	"github.com/goworkeryyt/configs/redis"
	"os"
)

func main() {
	// 获取全局配置
	globalConfigs := configs.GlobalConfigs()
	if globalConfigs == nil{
		fmt.Println("未能读取配置")
		os.Exit(1)
	}
	// 读取 redis的配置,方法一
	redis1 := globalConfigs.Redis
	fmt.Println(redis1)

	// 获取redis的配置,方法二
	redis2 := redis.Redis{}
	globalConfigs.SubItem("redis",&redis2)
	fmt.Println(redis2)
}

```

### 目录结构
```shell
configs
├
├── consul(注册中心配置)
├
├── database(数据库配置)
├
├── email(邮件配置)
├
├── ftp(文件服务器配置)
├
├── jwt(jwt token生成和校验配置)
├
├── mqtt(mqtt物联网配置)
├
├── pay(支付相关配置 支付宝和微信)
├
├── profile(多配置文件相关配置)
├
├── redis(redis缓存数据库相关配置)
├
├── resources(项目整合配置文件示例)
│   ├── active.yaml      配置指定要激活启用的配置文件
│   └── dev_config.yaml  开发环境配置文件
│   └── fat_config.yaml  功能验收测试环境配置文件
│   └── pro_config.yaml  生产环境配置文件
│   └── uat_config.yaml  用户验收测试环境配置文件
├
├── server(服务端口等相关配置)
├
├── zap(日志相关配置)
```



