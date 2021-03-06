package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/languagetranslatorv3"
)

var translator *languagetranslatorv3.LanguageTranslatorV3

func init() {
	apiKey, exist := os.LookupEnv("LANGUAGE_TRANSLATOR_APIKEY")
	if !exist {
		fmt.Println("No apikey defined")
	}
	urlService, exist := os.LookupEnv("LANGUAGE_TRANSLATOR_URL")
	if !exist {
		fmt.Println("No url defined")
	}

	// initial everything
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	langVersion := "2018-05-01"

	options := &languagetranslatorv3.LanguageTranslatorV3Options{
		Version:       &langVersion,
		Authenticator: authenticator,
	}

	var translateErr error

	translator, translateErr = languagetranslatorv3.NewLanguageTranslatorV3(options)

	if translateErr != nil {
		panic(translateErr)
	}

	translator.SetServiceURL(urlService)

}

func TranslateWords(words string, lang string) string {
	fmt.Println("in translate for " + words)

	// real translation
	if lang != "en" {
		result, _, translateErr := translator.Translate(
			&languagetranslatorv3.TranslateOptions{
				Text:    []string{words},
				ModelID: core.StringPtr(lang + "-en"),
			},
		)

		if translateErr != nil {
			panic(translateErr)
		}

		b, _ := json.Marshal(result.Translations[0].Translation)
		fmt.Println(string(b))

		return string(b)
	}

	return words

}
