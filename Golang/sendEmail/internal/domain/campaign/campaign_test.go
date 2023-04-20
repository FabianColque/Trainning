package campaign

import ("testing"
		"time"
		"github.com/stretchr/testify/assert")

var(
	name = "Campaign X"
	content = "Body"
	contacts = []string{"fbcolque@asd.com", "ernesto@asd.com"}
)

func TestNewCampaign_CreateCampaign(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	//assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaign_IDIsNotNil(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func TestNewCampaign_CreatedOn_MustToBeNow(t *testing.T){
	assert := assert.New(t)
	now	:= time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func TestNewCampaign_MustValidateName(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name is required", err.Error())
}

func TestNewCampaign_MustValidateContent(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("Content is required", err.Error())
}

func TestNewCampaign_MustValidateList(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("List is empty", err.Error())
}