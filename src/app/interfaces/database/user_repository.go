// データのやりとり

package database

// domain層をインポート
// 内側に依存しているので依存関係は守れている
import "app/domain"

// infrastructures層で定義したデータベース接続を実行
// 外側のレイヤーのルールを内側のレイヤーに持ち込んでいる！ように見えるが、
// 実態としてはuser_repository/goから呼び出している
type UserRepository struct {
  SqlHandler
}

// Userの作成・保存
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
  // Execute？
  result, err := repo.Execute(
    "INSERT INTO users (first_name, last_name) VALUES(?,?)", u.FisrtName, u.LastName
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
func (repo *UserRepository) FindById(identifier id) (user domain.User, err error) {
  // Queryで SELECT文を渡す
  row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)

  // 最後に実行する
  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var FisrtName string
  var LastName string

  // 行処理
  row.Next()

  // Scan()に変数ポインタを渡し、DBの結果をセット
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
