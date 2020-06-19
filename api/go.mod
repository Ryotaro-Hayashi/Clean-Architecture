// go mod init プロジェクト名（ここではapi）でモジュールを初期化
// go build でパッケージをインストール （Docker環境ではDockerfileに定義）
module api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.14
)
