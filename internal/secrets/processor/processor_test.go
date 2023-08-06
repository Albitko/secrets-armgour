package processor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

func Test_Login(t *testing.T) {
	authTests := []struct {
		name        string
		login       string
		password    string
		hashedPass  string
		expectedErr error
	}{
		{
			name:        "Auth: success",
			login:       "user",
			password:    "secretpass",
			hashedPass:  "$2a$10$D7P7FfZnPw6UKJfHofAWEuuFmEacoHpl/T9LH3zFCOlMJhaVTpxGy",
			expectedErr: nil,
		},
		{
			name:        "Auth: fail",
			login:       "user",
			password:    "secretpass",
			hashedPass:  "$2a$10$D7P7FfZnPw6UKJfHofAWGuuFmEacoHpl/T9LH3zFCOlMJhaVTpxGy",
			expectedErr: entity.ErrInvalidCredentials,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			testCredentials := entity.UserAuth{
				Login:    tt.login,
				Password: tt.password,
			}
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().GetUserPasswordHash(ctx, tt.login).Return(tt.hashedPass, nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.LoginUser(ctx, testCredentials)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_CardEdit(t *testing.T) {
	authTests := []struct {
		name        string
		index       string
		intIndex    int
		card        entity.UserCard
		expectedErr interface{}
	}{
		{
			name:        "CardEdit: Succes",
			index:       "1234",
			intIndex:    1234,
			card:        entity.UserCard{},
			expectedErr: nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().UpdateCard(ctx, tt.intIndex, tt.card).Return(nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.CardEdit(ctx, tt.index, tt.card)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_BinaryEdit(t *testing.T) {
	authTests := []struct {
		name        string
		index       string
		intIndex    int
		b64data     []byte
		binary      entity.UserBinary
		expectedErr interface{}
	}{
		{
			name:     "BinaryEdit: Succes",
			index:    "1234",
			intIndex: 1234,
			b64data:  []byte("test string"),
			binary: entity.UserBinary{
				B64Content: "dGVzdCBzdHJpbmc=",
			},
			expectedErr: nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().UpdateBinary(ctx, tt.intIndex, tt.binary, tt.b64data).Return(nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.BinaryEdit(ctx, tt.index, tt.binary)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_TextEdit(t *testing.T) {
	authTests := []struct {
		name        string
		index       string
		intIndex    int
		text        string
		textObj     entity.UserText
		expectedErr error
	}{
		{
			name:        "TextEdit: Succes",
			index:       "1234",
			intIndex:    1234,
			text:        "string",
			textObj:     entity.UserText{},
			expectedErr: nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().UpdateText(ctx, tt.intIndex, tt.textObj).Return(nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.TextEdit(ctx, tt.index, tt.textObj)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_CredentialsEdit(t *testing.T) {
	authTests := []struct {
		name        string
		index       string
		intIndex    int
		credentials entity.UserCredentials
		expectedErr error
	}{
		{
			name:        "CredentialsEdit: Succes",
			index:       "1234",
			intIndex:    1234,
			credentials: entity.UserCredentials{},
			expectedErr: nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().UpdateCredentials(ctx, tt.intIndex, tt.credentials).Return(nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.CredentialsEdit(ctx, tt.index, tt.credentials)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_GetUserData(t *testing.T) {
	authTests := []struct {
		name        string
		data        string
		id          string
		user        string
		res         entity.UserBinary
		expectedErr error
	}{
		{
			name:        "CredentialsEdit: Succes",
			data:        "binary",
			id:          "4",
			user:        "user",
			res:         entity.UserBinary{},
			expectedErr: nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().GetUserData(ctx, tt.data, tt.id, tt.user).Return(entity.UserBinary{}, nil).Once()
			secretsProcessor := New(mockRepo)
			res, err := secretsProcessor.GetUserData(ctx, tt.data, tt.id, tt.user)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.res, res)
		})
	}
}

func Test_BinaryCreation(t *testing.T) {
	authTests := []struct {
		name           string
		binary         entity.UserBinary
		user           string
		decodedContent []byte
		expectedErr    error
	}{
		{
			name: "CredentialsEdit: Succes",
			binary: entity.UserBinary{
				B64Content: "dGVzdCBzdHJpbmc=",
			},
			user:           "user",
			decodedContent: []byte("test string"),
			expectedErr:    nil,
		},
	}
	for _, tt := range authTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := newMockRepository(t)
			mockRepo.EXPECT().InsertBinary(ctx, tt.binary, tt.decodedContent, tt.user).Return(nil).Once()
			secretsProcessor := New(mockRepo)
			err := secretsProcessor.BinaryCreation(ctx, tt.binary, tt.user)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestCardCreation(t *testing.T) {
	mockRepo := newMockRepository(t)
	processor := &processor{
		repo: mockRepo,
	}
	mockRepo.On("InsertCard", mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	err := processor.CardCreation(context.TODO(), entity.UserCard{}, "user")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_ListUserData(t *testing.T) {
	mockRepo := newMockRepository(t)
	processor := &processor{
		repo: mockRepo,
	}
	mockRepo.On("SelectUserData", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil).Once()
	result, err := processor.ListUserData(context.TODO(), "data", "user")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestCredentialsCreation(t *testing.T) {
	mockRepo := newMockRepository(t)
	processor := &processor{
		repo: mockRepo,
	}
	mockRepo.On("InsertCredentials", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil).Once()
	err := processor.CredentialsCreation(context.TODO(), entity.UserCredentials{}, "user")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestTextCreation(t *testing.T) {
	mockRepo := newMockRepository(t)
	processor := &processor{
		repo: mockRepo,
	}
	mockRepo.On("InsertText", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil).Once()
	err := processor.TextCreation(context.TODO(), entity.UserText{}, "user")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestDeleteUserData(t *testing.T) {
	mockRepo := newMockRepository(t)
	processor := &processor{
		repo: mockRepo,
	}
	mockRepo.On("DeleteUserData", mock.Anything, "binary", "user").
		Return(nil, nil).Once()
	err := processor.DeleteUserData(context.TODO(), "binary", "user")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}
