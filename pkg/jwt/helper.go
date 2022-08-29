package jwt

import (
	"encoding/json"
	"github.com/IgorTkachuk/cartridge_accounting/internal/config"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/cache"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type RT struct {
	RefreshToken string `json:"refresh_token"`
}

type Helper interface {
	GenerateAccessToken(u user.User) ([]byte, error)
	UpdateRefreshToken(rt RT) ([]byte, error)
}

type helper struct {
	Logger  logging.Logger
	RTCache cache.Repository
}

func NewHelper(RTCache cache.Repository, logger logging.Logger) Helper {
	return &helper{
		Logger:  logger,
		RTCache: RTCache,
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	//Email string `json:"email"`
}

func (h helper) GenerateAccessToken(u user.User) ([]byte, error) {
	key := []byte(config.GetConfig().JWT.Secret)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(u.ID),
			Audience:  []string{"users"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	token, err := builder.Build(claims)
	if err != nil {
		return nil, err
	}

	h.Logger.Info("create refresh token")
	refreshTokenUUID := uuid.New()
	userBytes, _ := json.Marshal(u)
	err = h.RTCache.Set([]byte(refreshTokenUUID.String()), userBytes, 0)
	if err != nil {
		h.Logger.Error(err)
		return nil, err
	}

	jsonBytes, err := json.Marshal(map[string]string{
		"token":         token.String(),
		"refresh_token": refreshTokenUUID.String(),
	})

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil

}

func (h helper) UpdateRefreshToken(rt RT) ([]byte, error) {
	defer h.RTCache.Del([]byte(rt.RefreshToken))

	userBytes, err := h.RTCache.Get([]byte(rt.RefreshToken))
	if err != nil {
		return nil, err
	}

	var u user.User
	err = json.Unmarshal(userBytes, &u)
	if err != nil {
		return nil, err
	}

	return h.GenerateAccessToken(u)
}
