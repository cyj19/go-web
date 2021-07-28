package auth

import (
	"fmt"
	"go-web/internal/pkg/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	Post int
}

func NewApp(path, name, fileType string) *App {
	//加载配置文件
	err := config.Init(path, name, fileType)
	if err != nil {
		panic(err)
	}

	return &App{Post: config.GetServerPort()}
}

func (app *App) Run() {

	g := gin.Default()
	//初始化路由
	initRouter(g)
	g.Run(fmt.Sprintf(":%d", app.Post))

}
