package builder

import (
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type TwigBuilder struct {
}

func NewTwigBuilder() *TwigBuilder {
	return &TwigBuilder{}
}

func (b *TwigBuilder) Build(template entity.Template, params map[string]string) (body string, title string, err error) {
	return "", "", errors.New("not implemented")
}
