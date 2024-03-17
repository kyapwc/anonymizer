package tests

import (
	"gitlab.com/devskiller-tasks/golang-anonymizer/anonymizer"
	"testing"
)

var skypeTestCases = []struct {
	original    string
	expected    string
	description string
}{
	{`Lorem ipsum`, `Lorem ipsum`, "no skype username found"},
	{`Better call skype:MySkype`, `Better call skype:#`, "simple skype username"},
	{`Better call skype:MySkype or at least skype:YourAdmin`, `Better call skype:# or at least skype:#`, "Two simple skype usernames"},
	{`Better call skype:32My1000Skype32`, `Better call skype:#`, "Skype username with numbers"},
	{`Lorem ipsum <a href="skype:loremipsum?call">call</a> dolor sit amet`, `Lorem ipsum <a href="skype:#?call">call</a> dolor sit amet`, "username inside HTML tag"},
	{`Lorem ipsum  <a href="skype:loremipsum?call">call</a>, dolor sit <a href="skype:IPSUMLOREM?chat">chat</a> amet`, `Lorem ipsum  <a href="skype:#?call">call</a>, dolor sit <a href="skype:#?chat">chat</a> amet`, "multiple skype usernames"},
}

func TestSkypeAnonymize(t *testing.T) {
	skypeAnonymizer := anonymizer.NewSkypeAnonymizer("#")
	for _, test := range skypeTestCases {
		t.Run(test.description, func(t *testing.T) {
			anonymized := skypeAnonymizer.Anonymize(test.original)
			if anonymized != test.expected {
				t.Errorf("Anonymize(%s): expected '%s', actual '%s'", test.original, test.expected, anonymized)
			}
		})
	}
}
