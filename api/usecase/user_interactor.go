// interfaces/controllersへのGatewayの役割

package usecase

import "api/domain"

// interfaces層で定義したルールを持ち込んでいるように見える
type UserInteractor struct {
    UserRepository UserRepository
}

// Addメソッド
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
    // userを保存
    identifier, err := interactor.UserRepository.Store(u)
    if err != nil {
		  return
	  }
    // 作成したユーザーを返す
	  user, err = interactor.UserRepository.FindById(identifier)
	  return
}

// Usersメソッド
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
    // userを一覧で返す
    user, err = interactor.UserRepository.FindAll()
    return
}

// UserByIdメソッド
func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
    // userをid検索して返す
    user, err = interactor.UserRepository.FindById(identifier)
    return
}
