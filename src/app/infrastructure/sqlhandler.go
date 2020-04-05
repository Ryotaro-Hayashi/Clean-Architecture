// DB接続は、外部パッケージを利用するので内側にルールを持ち込まないようにinfrastructure層に定義

package infrastructure

// mysqlをgithubからインポート
import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// わからん
type SqlHandler struct {
  // * の使い方が分からない
  Conn *sql.DB
}

// わからん
func NewSqlHandler() *SqlHandler {
  // データベースへ接続。ドライバ名（mysql）と、user:password@tcp(host:port)/dbnameを指定。
  // tcp, := は何？
  conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")

  //接続でエラーが発生した場合の処理
  if err != nil {
      panic(err.Error)
  }

  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}
