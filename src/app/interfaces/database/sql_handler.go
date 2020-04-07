// user_repository.goから呼び出す

package database

type SqlHandler interface {
  // interface型に何もメソッドを定義していない（よってどんな型でも代入できる）
  // Result型, error型, Row型
  Execute(string, ...interface{}) (Result, error)
  Query(string, ...interface{}) (Row, error)
}

type Result interface {
  LastInsertId() (int64, error)
  RowsAffected() (int64, error)
}

type Row interface {
  Scan(...interface{}) error
  Next() bool
  Close() error
}
