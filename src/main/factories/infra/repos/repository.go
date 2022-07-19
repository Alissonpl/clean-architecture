package repos

import (
"clean-go/src/infra/repos"

)

func MakeUserRepository() repos.UserRepository {
	userRepository := repos.NewUserRepository()

	return userRepository
}