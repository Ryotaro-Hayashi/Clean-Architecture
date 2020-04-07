// ①DB接続は、外部パッケージを利用するので内側にルールを持ち込まないようにinfrastructure層に定義

package infrastructure

// mysqlをgithubからインポート
import (
  // database/sqlパッケージ
  "database/sql"
  // mysql用のドライバー
  _ "github.com/go-sql-driver/mysql"
  // ②interfacesでDB接続ができるように、interfacesで定義したロジックをインポートして依存関係を逆転させる
  "app/interfaces/database"
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

// interfacesで使うメソッドを定義

// Executeメソッド
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
    res := SqlResult{}
    result, err := handler.Conn.Exec(statement, args...)
    if err != nil {
        return res, err
    }
    res.Result = result
    return res, nil
}

// Queryメソッド
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
    rows, err := handler.Conn.Query(statement, args...)
    if err != nil {
        return new(SqlRow), err
    }
    row := new(SqlRow)
    row.Rows = rows
    return row, nil
}

// LastInsertIdメソッドの呼び出しもと
type SqlResult struct {
    Result sql.Result
}

// LastInsertIdメソッド
func (r SqlResult) LastInsertId() (int64, error) {
    return r.Result.LastInsertId()
}

// RowsAffectedメソッド
func (r SqlResult) RowsAffected() (int64, error) {
    return r.Result.RowsAffected()
}

// Scanメソッド, Nextメソッド, Closeメソッドの呼び出しもと
type SqlRow struct {
    Rows *sql.Rows
}

// Scanメソッド
func (r SqlRow) Scan(dest ...interface{}) error {
    return r.Rows.Scan(dest...)
}

// Nextメソッド
func (r SqlRow) Next() bool {
    return r.Rows.Next()
}

// Closeメソッド
func (r SqlRow) Close() error {
    return r.Rows.Close()
}
