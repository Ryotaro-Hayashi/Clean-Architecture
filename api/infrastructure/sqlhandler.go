// ①DB接続は、外部パッケージを利用するので内側にルールを持ち込まないようにinfrastructure層に定義

package infrastructure

import (
  // database/sqlパッケージ
  "database/sql"
  // mysql用のドライバー
  _ "github.com/go-sql-driver/mysql"
  // ②interfacesでDB接続ができるように、interfacesで定義したロジックをインポートして依存関係を逆転させる
  "api/interfaces/database"
  // gorm
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
  "api/domain"  
)

type SqlHandler struct {
  // DB型
  Conn *sql.DB
}

type GormHandler struct {
  Conn *gorm.DB
}

// New + 構造体名 という構造体を初期化する関数名の命名慣例
func NewSqlHandler() database.SqlHandler {
  // データベースへ接続するためのhandlerを取得。ドライバ名（mysql）と、user:password@tcp(host:port)/dbnameを指定。
  // ローカル環境で、tcp を入れると nil pointer エラーが出る(mysql.server start も忘れずに)
  // Docker環境では、tcp を入れないと nil pointer エラーが出る
  conn, err := sql.Open("mysql", "root@tcp(db:3306)/CleanArchitecture")

  //接続でエラーが発生した場合の処理
  if err != nil {
      panic(err.Error)
  }

  // newは、指定した型のポインタ型を生成する関数（構造体の初期化）
  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}

func NewGormHandler() database.GormHandler {
    db, err := gorm.Open("mysql", "root@tcp(db:3306)/CleanArchitecture")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(domain.User{}) // マイグレーション

	db.Create(domain.User{ // データの挿入
        FirstName: "test",
        LastName: "man",
	})

	gormHandler := new(GormHandler)
	gormHandler.Conn = db

	return gormHandler
}

// Executeメソッド￥
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
    res := SqlResult{}
    // Exec は Query と違い、行を返さず、要約して返す
    // 引数は(クエリ, クエリ内のパラメーター)
    result, err := handler.Conn.Exec(statement, args...)
    if err != nil {
        return res, err
    }
    res.Result = result
    return res, nil
}

// Queryメソッド
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
    // 行を返す
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
    // Result型
    Result sql.Result
}

// LastInsertIdメソッド
func (r SqlResult) LastInsertId() (int64, error) {
    // 最後に挿入された要素のidを返す
    return r.Result.LastInsertId()
}

// RowsAffectedメソッド
func (r SqlResult) RowsAffected() (int64, error) {
    // 更新・挿入・削除によって影響を受けた行を返す
    return r.Result.RowsAffected()
}

// Scanメソッド, Nextメソッド, Closeメソッドの呼び出しもと
type SqlRow struct {
    // Rows型
    Rows *sql.Rows
}

// Scanメソッド
func (r SqlRow) Scan(dest ...interface{}) error {
    // dest に行の値をコピーする
    return r.Rows.Scan(dest...)
}

// Nextメソッド
// Scanメソッドが読み取れるように結果行をセット
// Scanメソッドより先行しなくてはいけない
func (r SqlRow) Next() bool {
    // セットに成功するとtrueを返す
    return r.Rows.Next()
}

// Closeメソッド
func (r SqlRow) Close() error {
    return r.Rows.Close()
}
