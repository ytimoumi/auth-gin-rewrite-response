package limengo

import (
	"auth/redis"
	"context"
	"github.com/google/uuid"
	"log"

	authsvc "auth/gen/authenticate"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authSvc struct {
	logger *log.Logger
	r      redis.RedisServicer
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger) authsvc.Service {
	return &authSvc{
		logger: logger,
		r:      redis.NewDefaultRedisServicer(),
	}
}

func (a *authSvc) Authenticate(ctx context.Context, auth *authsvc.InputAuth) (*authsvc.OutputAuth, error) {
	var err error
	key, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	err = a.r.RedisSetAccessToken(key.String(), "yassine")
	if err != nil {
		return nil, err
	}
	t := key.String()
	locale := "fr_FR"
	providerToken := ""
	people := "/people/peoples/1"
	res := &authsvc.OutputAuth{
		Login:         auth.Login,
		RefreshToken:  t,
		Subscriber:    *auth.XProvider,
		ProviderToken: &providerToken,
		People:        &people,
		Locale:        &locale,
	}
	return res, nil
}
