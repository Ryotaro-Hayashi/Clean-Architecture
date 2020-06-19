// データのやりとり

package database

// domain層をインポート
// 内側に依存しているので依存関係は守れている
import (
  "api/domain"
  "log"
)
// usecaseをインポートする必要があるのでは？？
// interfaces/controllersでインポートしているので大丈夫（？）

// infrastructures層で定義したデータベース接続を実行
// 外側のレイヤーのルールを内側のレイヤーに持ち込んでいる！ように見えるが、
// 実態としてはuser_repository/goから呼び出している
type UserRepository struct {
  SqlHandler
  GormHandler
}

// Userの保存
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
  // クエリを実行して結果を要約して返す
  result, err := repo.Execute(
    "INSERT INTO users (first_name, last_name) VALUES(?,?)", u.FirstName, u.LastName,
  )
  if err != nil {
    return
  }

  // LastInsertIdメソッドで最後に挿入（保存）されたidを取得
  id64, err := result.LastInsertId()
  if err != nil {
    return
  }

  // 保存されたid
  id = int(id64)
  return
}

// idによるUserの検索
func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
  // クエリを実行して行を返す
  row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)

  // 最後に実行する
  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var firstName string
  var lastName string

  // Scanメソッドで読み取れるように結果行をセット
  row.Next()

  // Scan()に変数ポインタを渡し、各変数にrowの値をコピー
  // err を定義して ; で条件文と仕切る
  if err = row.Scan(&id, &firstName, &lastName); err != nil {
        return
  }

  user.ID = id
  user.FirstName = firstName
  user.LastName = lastName
  return
}

// User一覧
func (repo *UserRepository) FindAll() (users domain.Users, err error) {
  rows, err := repo.Query("SELECT id, first_name, last_name FROM users")

  defer rows.Close()

  if err != nil {
    return
  }

  // 一行一行処理を行う
  for rows.Next() {
    var id int
    var firstName string
    var lastName string
    if err := rows.Scan(&id, &firstName, &lastName); err != nil {
        continue
    }
    user := domain.User{
        ID:        id,
        FirstName: firstName,
        LastName:  lastName,
    }
    // usersにuserを追加
    users = append(users, user)
  }
  return
}

// ユーザー情報を全取得して返す
func (repo *UserRepository) GormFindAll() (users domain.Users, err error) {
  // ユーザー情報を全取得
  users, err = repo.Find()
  if err != nil {
    return
  }

  log.Print("The users are ", users)

  return
}
