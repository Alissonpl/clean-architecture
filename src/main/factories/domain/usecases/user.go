package usecases

import (
"clean-go/src/domain/usecase/user"
"clean-go/src/main/factories/infra/repos"

)

func MakeUserService() user.UseCase {
	userRepository := repos.MakeUserRepository()

    userService := user.NewService(userRepository)
	return userService
}