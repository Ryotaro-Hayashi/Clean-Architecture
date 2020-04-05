// データのやりとり

package database

// domain層をインポート
impost "app/domain"

// infrastructures層で定義したデータベース接続を実行
type UserRepository struct {
  SqlHandler
}

// Userの作成・保存
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
  result, err := repo.Execute(
    "INSERT INTO users (first_name, last_name) VALUES(?,?)", u.FisrtName, u.LastName
  )
  if err != nil {
    return
  }

  // id64, LastInsertId って何？
  id64, err := result.LastInsertId()
  if err != nil {
    return
  }

  id = int(id64)
  return
}

// idによるUserの検索
func (repo *UserRepository) FindById(identifier id) (user domain.User, err error) {
  row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
  // わからん
  defer row.Close()
  if err != nil {
    return
  }

  var id int
  var FisrtName string
  var LastName string

  // わからん
  row.Next()
  // Sran, &, ; ？
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
    // insert, push などとの違いが分からない
    users = append(users, user)
  }
  return
}
