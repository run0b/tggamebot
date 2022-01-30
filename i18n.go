package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func CreateLocale() *i18n.Localizer {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("lang/ru.yml")
	localizer := i18n.NewLocalizer(bundle,
		language.Russian.String(),
		language.English.String())
	return localizer
}
