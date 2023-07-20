package di

import (
	"github.com/google/wire"

	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	repoAdapter "github.com/thnkrn/comet/api/pkg/repository/adapter"
	usecaseAdapter "github.com/thnkrn/comet/api/pkg/usecase/adapter"
)

var AdminSet = wire.NewSet(
	repoAdapter.NewAdminRepository, usecaseAdapter.NewAdminUseCase, handler.NewAdminHandler,
)
