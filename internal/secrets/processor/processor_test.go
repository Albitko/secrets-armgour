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
