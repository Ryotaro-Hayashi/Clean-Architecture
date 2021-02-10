package infrastructure

import (
    "api/interfaces/controllers"
    "github.com/gin-contrib/pprof"
    // gin をインポート
    "github.com/gin-gonic/gin"
)

// server.go で Run するためのエンジンを初期化
var Router *gin.Engine

func init() {
    // デフォルトのミドルウェア付きのEngineインスタンスを作成
    router := gin.Default()

    pprof.Register(router)

    // DB接続
    userController := controllers.NewUserController(NewSqlHandler())

    // interfaces/cotrollers のメソッドを使ってルーティングを設定
    // パラメーターは gin.Context が受け取っている
    router.POST("/users", func(c *gin.Context) { userController.Create(c) })
    router.GET("/users", func(c *gin.Context) { userController.Index(c) })
    router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

    // server.go で Run するためのエンジンに格納
    Router = router
}
