package domain

import (
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/utils"
	"github.com/wrapped-owls/goremy-di/remy"
)

type ContactsService struct {
	repository interfaces.ContactsRepository
}

func NewContactsService() *ContactsService {
	return &ContactsService{
		repository: remy.Get[interfaces.ContactsRepository](utils.Injector),
	}
}
