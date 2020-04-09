// interfaces/databaseからの input port の役割

package usecase

import "api/domain"

// ロジックとしてUserRepository型とメソッドを定義
type UserRepository interface {
    Store(domain.User) (int, error)
    FindById(int) (domain.User, error)
    FindAll() (domain.Users, error)
}
