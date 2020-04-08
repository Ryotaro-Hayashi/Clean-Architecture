package domain

// import "fmt"

// User型を宣言
type User struct {
  ID int
  FirstName string
  LastName string
  // FullName  string
}

// User型のスライスを宣言
type Users []User

// func (u *User) Build() *User {
// 	u.FullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
// 	return u
// }
