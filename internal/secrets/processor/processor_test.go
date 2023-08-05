package processor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

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
			name:     "TextEdit: Succes",
			index:    "1234",
			intIndex: 1234,
			text:     "string",
			textObj: entity.UserText{},
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
