// gin.Context を使用するため

package controllers

type Context interface {
    Param(string) string
    Bind(interface{}) error
    Status(int)
    JSON(int, interface{})
}
