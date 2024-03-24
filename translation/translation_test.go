package translation_test

import (
	"shipping_go/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	//word := "Hi"
	//language := "english"
	//res := translation.Translate("Hi", "english")
	//if res != "Hi" {
	//	t.Errorf(`expected "Hi" but received "%s"`, res)
	//}

	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "Hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
	}

	for _, test := range tt {
		r := translation.Translate(test.Word, test.Language)

		if r != test.Translation {
			t.Errorf(`expected %s but received "%s"`, test.Translation, r)
		}
	}

}
