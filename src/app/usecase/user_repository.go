// interfaces/databaseからの input port の役割

package usecase

import "app/domain"

// ロジックとしてUserRepository型とメソッドを定義
type UserRepository interface {
    Store(domain.User) (int, error)
    FindById(int) (domain.User, error)
    FindAll() (domain.Users, error)
}

// interfacesでこのファイル（usecase）をインポートする必要があるのでは？
