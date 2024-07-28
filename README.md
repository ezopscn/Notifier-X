<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">🥳 Notifier X</h1>
<h3 align="center">一个适合小规模团队的全能型消息管理平台</h3>

<p align="center">
  <a>
    <img src="https://img.shields.io/badge/-Golang 1.20-blue?style=flat-square&logo=go&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Gin 1.9.1-blue?style=flat-square&logo=gin&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-MySQL-blue?style=flat-square&logo=mysql&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Redis-c14438?style=flat-square&logo=redis&logoColor=white&link=mailto:ezops.cn@gmail.com" alt="">
  </a>
</p>

<hr>

### 📩 支持信息

| 平台    | 是否支持 | 补充说明 |
|:------:|:----:|:----:|
| 钉钉    |  ✅   | 开发中  |
| 飞书    |  ❌   | 开发中  |
| 企业微信  |  ❌   | 开发中  |
| 邮件    |  ❌   | 开发中  |
| 阿里云短信 |  ❌   | 开发中  |
| 腾讯云短信 |  ❌   | 开发中  |

### 🤔 技术栈

- [x] Go：Google 开发的开源编程语言，诞生于 2006 年 1 月 2 日 15 点 4 分 5 秒 [🌏](https://github.com/golang/go)
- [x] Cobra：CLI 开发参数处理工具 [🌏](https://github.com/spf13/cobra)
- [x] Embed：go 1.16 新增的文件嵌入属性, 轻松将静态文件打包到编译后的二进制应用中
- [x] Gin：用 Go 编写的高性能 HTTP Web 框架 [🌏](https://github.com/gin-gonic/gin)
- [x] Viper：配置管理工具, 支持多种配置文件类型 [🌏](https://github.com/spf13/viper)
- [x] Zap：提供快速、结构化、分级的日志记录 [🌏](https://pkg.go.dev/go.uber.org/zap)
- [x] Lumberjack：日志滚动切割工具 [🌏](https://github.com/natefinch/lumberjack)
- [x] Gorm：数据库 ORM 管理框架, 可自行扩展多种数据库类型 [🌏](https://gorm.io/gorm)
- [x] Carbon：简单、语义化且对开发人员友好的 datetime 包 [🌏](https://github.com/golang-module/carbon)
- [x] Redis：Redis 客户端 [🌏](https://github.com/redis/go-redis)
- [x] Sonic：字节开源的高性能 JSON 库 [🌏](https://github.com/bytedance/sonic)
- [x] Multitemplate：Gin 多模板继承库 [🌏](https://github.com/gin-contrib/multitemplate)

<br>

### ⚡ 开发依赖

```bash
# 命令行工具
go get -u github.com/spf13/cobra

# Golang web 开发框架
go get -u github.com/gin-gonic/gin

# YAML 配置文件解析成结构体
go get -u github.com/spf13/viper

# 日志
go get -u go.uber.org/zap

# 日志切割
go get -u github.com/natefinch/lumberjack

# 数据库 GORM
go get -u gorm.io/gorm

# MySQL 连接驱动
go get -u gorm.io/driver/mysql

# 日期时间
go get -u github.com/golang-module/carbon/v2

# Redis 客户端
go get -u github.com/redis/go-redis/v9

# Sonic Json 库
go get -u github.com/bytedance/sonic

# gin 多模板继承
go get -u github.com/gin-contrib/multitemplate

# JSON 快速解析
go get -u github.com/tidwall/gjson

# 异步并发任务
go get -u golang.org/x/sync/semaphore
```