package domain

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/mocks"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
	"strings"
	"testing"
)

func TestContactsService_Validate(t *testing.T) {
	ctrl := gomock.NewController(t)

	const expectedName = "TEMP_MOCK_NAME"
	const expectedPhone = "+55 (75) 98800-4050"
	repoMock := mocks.NewMockContactsRepository(ctrl)
	dtoMock := mocks.NewMockContactDTO(ctrl)

	dtoMock.EXPECT().Name().Return(expectedName).AnyTimes()
	dtoMock.EXPECT().Phone().Return(expectedPhone).AnyTimes()

	// Asserts the that when
	repoMock.
		EXPECT().
		GetByPhone(gomock.Any()).
		Return(dtoMock, nil)

	service := NewContactsService(repoMock)

	tempDto := mocks.NewMockContactDTO(ctrl)
	tempDto.EXPECT().Phone().Return("5575988004050").AnyTimes()
	tempDto.EXPECT().Name().Return(strings.ToLower("not THE s@m3 nUmBeR")).AnyTimes()

	err := service.Add([]dtos.ContactsDTO{tempDto})
	if err != nil && !errors.Is(err, corerrs.ErrContactAlreadyExists) {
		t.Error(err)
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

	service := NewContactsService(repoMock)
	err := service.Add([]dtos.ContactsDTO{tempDto, dtoMock})
	if err != nil && !errors.Is(err, corerrs.ErrContactAlreadyExists) {
		t.Error(err)
	}

	if len(mockList) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(mockList))
	}
}
