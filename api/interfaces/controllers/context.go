// gin.Context を使用するため

package controllers

// Context で使う組み込みメソッドを定義
type Context interface {
    Param(string) string
    Bind(interface{}) error
    Status(int)
    JSON(int, interface{})
}
