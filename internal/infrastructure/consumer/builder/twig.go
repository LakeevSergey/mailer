package builder

import (
	"bytes"

	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type TwigBuilder struct {
}

func NewTwigBuilder() *TwigBuilder {
	return &TwigBuilder{}
}

func (b *TwigBuilder) Build(template entity.Template, params map[string]string) (body string, title string, err error) {
	env := twig.New(nil)

	values := make(map[string]stick.Value, len(params))
	for key, value := range params {
		values[key] = value
	}

	bufBody := new(bytes.Buffer)

	err = env.Execute(template.Body, bufBody, values)
	if err != nil {
		return "", "", err
	}

	bufTitle := new(bytes.Buffer)

	err = env.Execute(template.Title, bufTitle, values)
	if err != nil {
		return "", "", err
	}

	return bufBody.String(), bufTitle.String(), nil
}
