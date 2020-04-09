package controllers

import (
  "api/domain"
  "api/usecase"
  "api/interfaces/database"
  // 何？
  "strconv"
)

// usecase層のUserInteractorを使用
type UserController struct {
  // usecase層でUserInteractorはinterfaces層にあるUserRepositoryを参照しているのでinterfaces/databaseをインポート
  Interactor usecase.UserInteractor
}

// 構造体を初期化する. 入れ子になっている.
func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

// ユーザーの Createメソッド
func (controller *UserController) Create(c Context) {
    // 入力されたユーザーを受け取るためのUser型を初期化
    u := domain.User{}
    // Bindは、Content-TypeをチェックしてバインドするContext（JsonとXML以外だとエラーを吐く）
    // Postリクエストで受け取ったユーザーをバインド
    c.Bind(&u)
    // ユーザーを追加・保存
    user, err := controller.Interactor.Add(u)
    if err != nil {
        c.JSON(500, err)
        return
    }
    c.JSON(201, user)
}

// Indexメソッド
func (controller *UserController) Index(c Context) {
    users, err := controller.Interactor.Users()
    if err != nil {
        c.JSON(500, err)
        return
    }
    c.JSON(200, users)
}

// Showメソッド
func (controller *UserController) Show(c Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := controller.Interactor.UserById(id)
    if err != nil {
        c.JSON(500, err)
        return
    }
    c.JSON(200, user)
}
