package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	abzEntities "github.com/jictyvoo/multi_client_rest_api/modules/apps/abzcore/domain/entities"
	xycEntities "github.com/jictyvoo/multi_client_rest_api/modules/apps/xyccore/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/mocks"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
	"strings"
	"testing"
)

type contactsTestCase[T any] struct {
	name         string
	service      func(retriever remy.DependencyRetriever) services.ContactsServiceFacade
	expectations T
}

func TestContactsService_Validate(t *testing.T) {
	var testCases = [...]contactsTestCase[[2][2]string]{
		{
			name: "ABZ_1 - Valid name and phone",
			service: func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
				return domain.NewContactsService[xycEntities.ContactEntity](
					remy.Get[interfaces.ContactsRepository](retriever),
				)
			},
			expectations: [2][2]string{
				{"TEMP_MOCK_NAME", "+55 (75) 98800-4050"},
				{"not THE s@m3 nUmBeR", "5575988004050"},
			},
		},
		{
			name: "XYC_2 - Valid name and phone",
			service: func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
				return domain.NewContactsService[abzEntities.ContactEntity](
					remy.Get[interfaces.ContactsRepository](retriever),
				)
			},
			expectations: [2][2]string{
				{"TEMP_MOCK_NAME", "+55 (75) 98800-4050"},
				{"NOT THE SAME NUMBER", "5575988004050"},
			},
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)

		injector := remy.NewInjector()
		remy.Register(injector, remy.Factory(func(retriever remy.DependencyRetriever) interfaces.ContactsRepository {
			repoMock := mocks.NewMockContactsRepository(ctrl)
			dtoMock := mocks.NewMockContactDTO(ctrl)

			dtoMock.EXPECT().Name().Return(testCase.expectations[0][0]).AnyTimes()
			dtoMock.EXPECT().Phone().Return(testCase.expectations[0][1]).AnyTimes()

			// Asserts the that when
			repoMock.
				EXPECT().
				GetByPhone(gomock.Any()).
				Return(dtoMock, nil).
				AnyTimes()
			return repoMock
		}))
		remy.Register(injector, remy.Factory(testCase.service))
		service := remy.Get[services.ContactsServiceFacade](injector)

		tempDto := mocks.NewMockContactDTO(ctrl)
		tempDto.EXPECT().Name().Return(strings.ToLower(testCase.expectations[1][0])).AnyTimes()
		tempDto.EXPECT().Phone().Return(testCase.expectations[1][1]).AnyTimes()

		err := service.Add([]dtos.ContactsDTO{tempDto})
		if err != nil && !errors.Is(err, corerrs.ErrContactAlreadyExists) {
			t.Error(err)
		}
	}
}

func TestContactsService_Add(t *testing.T) {
	ctrl := gomock.NewController(t)

	const expectedName = "TEMP_MOCK_NAME"
	const expectedPhone = "+55 (75) 98800-4050"
	mockList := make([]interfaces.ContactDTO, 0, 2)

	repoMock := mocks.NewMockContactsRepository(ctrl)

	// Asserts the that when
	repoMock.
		EXPECT().
		GetByPhone(gomock.Any()).
		DoAndReturn(func(phone string) (interfaces.ContactDTO, error) {
			for _, dto := range mockList {
				if dto.Phone() == phone {
					return dto, corerrs.ErrContactAlreadyExists
				}
			}
			return nil, nil
		}).Times(2)

	repoMock.
		EXPECT().
		Add(gomock.Any()).
		Do(func(contact entities.ContactEntity) {
			mockList = append(mockList, contact)
		}).
		AnyTimes()

	dtoMock := mocks.NewMockContactDTO(ctrl)
	dtoMock.EXPECT().Name().Return(expectedName).AnyTimes()
	dtoMock.EXPECT().Phone().Return(expectedPhone).AnyTimes()

	tempDto := mocks.NewMockContactDTO(ctrl)
	tempDto.EXPECT().Phone().Return(expectedPhone).AnyTimes()
	tempDto.EXPECT().Name().Return(strings.ToLower("not THE s@m3 nUmBeR")).AnyTimes()

	service := domain.NewContactsService[xycEntities.ContactEntity](repoMock)
	err := service.Add([]dtos.ContactsDTO{tempDto, dtoMock})
	if err != nil && !errors.Is(err, corerrs.ErrContactAlreadyExists) {
		t.Error(err)
	}

	if len(mockList) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(mockList))
	}
}

func TestContactsService_AddAll(t *testing.T) {
	// TODO: Implement this test
}
