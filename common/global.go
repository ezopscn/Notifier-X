package common

import "go.uber.org/zap"

// Logo 图形生成网站：http://patorjk.com/software/taag/
const Logo = `
 ▄▄        ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄       ▄       ▄ 
▐░░▌      ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░▌     ▐░▌
▐░▌░▌     ▐░▌▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀  ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌      ▐░▌   ▐░▌ 
▐░▌▐░▌    ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌     ▐░▌               ▐░▌     ▐░▌          ▐░▌       ▐░▌       ▐░▌ ▐░▌  
▐░▌ ▐░▌   ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌     ▐░█▄▄▄▄▄▄▄▄▄      ▐░▌     ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌        ▐░▐░▌   
▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌     ▐░░░░░░░░░░░▌     ▐░▌     ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌         ▐░▌    
▐░▌   ▐░▌ ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌     ▐░█▀▀▀▀▀▀▀▀▀      ▐░▌     ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀█░█▀▀         ▐░▌░▌   
▐░▌    ▐░▌▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌     ▐░▌               ▐░▌     ▐░▌          ▐░▌     ▐░▌         ▐░▌ ▐░▌  
▐░▌     ▐░▐░▌▐░█▄▄▄▄▄▄▄█░▌     ▐░▌      ▄▄▄▄█░█▄▄▄▄ ▐░▌           ▄▄▄▄█░█▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌      ▐░▌       ▐░▌   ▐░▌ 
▐░▌      ▐░░▌▐░░░░░░░░░░░▌     ▐░▌     ▐░░░░░░░░░░░▌▐░▌          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌     ▐░▌     ▐░▌
 ▀        ▀▀  ▀▀▀▀▀▀▀▀▀▀▀       ▀       ▀▀▀▀▀▀▀▀▀▀▀  ▀            ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀       ▀       ▀
`

// 全局工具
var (
	SystemLog *zap.SugaredLogger // 系统日志工具
	AccessLog *zap.SugaredLogger // 访问日志工具
)

// 系统信息
var (
	APPName     = "Notifier X"          // 应用名称
	APIPrefix   = "/api/v1"             // API 前缀
	Version     = "1.0"                 // 版本信息
	ConfigFile  = "config/default.yaml" // 默认配置文件
	VersionFile = "config/version"      // 版本配置文件
)

// 时间格式
const (
	MsecTimeFormat = "2006-01-02 15:04:05.000"
	SecTimeFormat  = "2006-01-02 15:04:05"
	DateTimeFormat = "2006-01-02"
)
