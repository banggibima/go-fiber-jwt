package memory

import (
	"github.com/banggibima/go-fiber-jwt/internal/domain/entity"
)

type TokenMemory interface {
	ReadByRefreshToken(refreshToken string) (*entity.Token, error)
	Create(token *entity.Token) error
	DeleteByRefreshToken(refreshToken string) error
}
