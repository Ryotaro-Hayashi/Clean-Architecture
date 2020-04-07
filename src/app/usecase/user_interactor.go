// interfaces/controllersへのGatewayの役割

package usecase

import "app/domain"

// interfaces層で定義したルールを持ち込んでいるように見える
type UserInteractor struct {
    UserRepository UserRepository
}

// Addメソッド
func (interactor *UserInteractor) Add(u domain.User) (err error) {
    _, err := interactor.UserRepository.Store(u)
    return
}

// Usersメソッド
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
    user, err = interactor.UserRepository.FindAll()
    return
}

// UserByIdメソッド
func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
    user, err = interactor.UserRepository.FindById(identifier)
    return
}
