package tests

import (
	"gitlab.com/devskiller-tasks/golang-anonymizer/anonymizer"
	"testing"
)

var offerTestCases = []struct {
	orig     string
	expected string
}{
	{`Lorem ipsum a@anonymizer.com. <a href="skype:loremipsum?call">call</a> +48 666 777 888`,
		`Lorem ipsum REPLACED@anonymizer.com. <a href="skype:REPLACED?call">call</a> +48 666 777 8XX`},
}

var offerTestAnonymizers = []anonymizer.Anonymizer{
	anonymizer.NewEmailAnonymizer("REPLACED"),
	anonymizer.NewSkypeAnonymizer("REPLACED"),
	buildPhoneAnonymizer(),
}

func buildPhoneAnonymizer() anonymizer.Anonymizer {
	phoneAnonymizer := anonymizer.NewPhoneAnonymizer().SetReplacement("XXX")
	_ = phoneAnonymizer.SetLastDigits(2)
	return phoneAnonymizer
}

func TestOfferAnonymize(t *testing.T) {
	offerAnonymizer := anonymizer.NewOfferAnonymizer()

	for _, value := range offerTestAnonymizers {
		offerAnonymizer.AddAnonymizer(value)
	}

	for _, testCase := range offerTestCases {
		anonymized := offerAnonymizer.Anonymize(testCase.orig)
		if anonymized != testCase.expected {
			t.Errorf("Anonymize(%s): expected %s, actual %s", testCase.orig, testCase.expected, anonymized)
		}
	}
}
