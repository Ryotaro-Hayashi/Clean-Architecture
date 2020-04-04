// なんのため？
package domain

// Userモデルのカラムとその型を宣言？
type User struct {
  ID int
  FirstName string
  LastName string
}

// Usersという変数がUserの配列ということを宣言？
type Users []User
