package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/languagetranslatorv3"
)

func TranslateWords(words string) string {

	apiKey := os.Getenv("WATSON_API_KEY")
	urlService := os.Getenv("WATSON_URL_SERVICE")

	// initial everything
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	langVersion := "2018-05-01"

	options := &languagetranslatorv3.LanguageTranslatorV3Options{
		Version:       &langVersion,
		Authenticator: authenticator,
	}

	translator, err := languagetranslatorv3.NewLanguageTranslatorV3(options)

	if err != nil {
		panic(err)
	}

	translator.SetServiceURL(urlService)

	// real translation
	result, _, translateErr := translator.Translate(
		&languagetranslatorv3.TranslateOptions{
			Text:    []string{words},
			ModelID: core.StringPtr("id-en"),
		},
	)

	if translateErr != nil {
		panic(translateErr)
	}

	b, _ := json.Marshal(result.Translations[0].Translation)
	fmt.Println(string(b))

	return string(b)

}
