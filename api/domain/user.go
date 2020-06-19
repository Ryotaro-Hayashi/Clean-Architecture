package domain

// User型を宣言
type User struct {
  ID int
  FirstName string
  LastName string
}

// User型のスライスを宣言
// type で型に別名を付けることができる（ここでは Users型）
type Users []User
