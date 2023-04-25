package campaign

import ("testing"
		"sendEmail/internal/contract"
		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/mock"
		"errors")

type repositoryMock struct{
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error{
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name : "Test Y",
		Content : "Body",
		Emails : []string{"test1@test.com"},
	}
	service = Service{}	
)

/*func Test_Create_Campaign(t *testing.T){
	assert := assert.New(t)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}*/

func Test_Create_ValidateDomainError(t *testing.T){
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("Name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T){
	repositoryMock := new(repositoryMock)	
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}else if campaign.Content != newCampaign.Content {
			return false
		}else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock	

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T){
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)	
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock	

	_, err := service.Create(newCampaign)

	assert.True(erros.Is(internalerros.ErrInternal, err))
	assert.Equal("error to save on database", err.Error())
}