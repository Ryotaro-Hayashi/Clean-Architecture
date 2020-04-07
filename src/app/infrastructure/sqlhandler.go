// DB接続は、外部パッケージを利用するので内側にルールを持ち込まないようにinfrastructure層に定義

package infrastructure

// mysqlをgithubからインポート
import (
  // database/sqlパッケージ
  "database/sql"
  // mysql用のドライバー
  _ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
  // database/sqlパッケージによるデータベース接続に必要なtype
  Conn *sql.DB
}

// New + 構造体名 という構造体を初期化する関数名の命名慣例
func NewSqlHandler() *SqlHandler {
  // データベースへ接続するためのhandlerを取得。ドライバ名（mysql）と、user:password@tcp(host:port)/dbnameを指定。
  // tcp, := は何？
  conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")

  //接続でエラーが発生した場合の処理
  if err != nil {
      panic(err.Error)
  }

  // newは、指定した型のポインタ型を生成する関数（構造体の初期化）
  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}
