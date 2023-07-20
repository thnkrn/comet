package di

import (
	"github.com/google/wire"

	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	repoAdapter "github.com/thnkrn/comet/api/pkg/repository/adapter"
	usecaseAdapter "github.com/thnkrn/comet/api/pkg/usecase/adapter"
)

var UserSet = wire.NewSet(
	repoAdapter.NewUserRepository, usecaseAdapter.NewUserUseCase, handler.NewUserHandler,
)
