package processor

import (
	"context"
	"testing"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

func Test_Login(t *testing.T) {
	ctx := context.Background()
	testCredentials := entity.UserAuth{
		Login:    "test_login",
		Password: "test_password",
	}

	mockRepo := newMockRepository(t)
	mockRepo.EXPECT().GetUserPasswordHash(ctx, testCredentials.Login).Return("0x0", nil).Once()

	secretsProcessor := New(mockRepo)

	err := secretsProcessor.LoginUser(ctx, testCredentials)
	if err != nil {
		return
	}
}
