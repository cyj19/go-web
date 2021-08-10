package initialize

import (
	"fmt"
	"go-web/internal/pkg/global"
	"log"
	"os"
)

// 使用标准log库
func InitLogger() {

	logFile, err := os.OpenFile("admin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件失败：%v", err))
	}
	global.LoggerIns = log.New(logFile, "<Custom>", log.Lshortfile|log.Ldate|log.Ltime)
	fmt.Println("初始化日志完成...")
}
