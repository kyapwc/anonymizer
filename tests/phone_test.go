package tests

import (
	"gitlab.com/devskiller-tasks/golang-anonymizer/anonymizer"
	"testing"
)

var phoneNumbersTestCases = []struct {
	replacement           string
	digitsAnonimizedCount int
	original              string
	expected              string
	description           string
}{
	{"X", 3, "Lorem ipsum", "Lorem ipsum", "no phone number"},
	{"X", 0, "Lorem +48 666 666 666 d00r", "Lorem +48 666 666 666 d00r", "Anonymizer turned off"},
	{"X", 3, "Lorem +48 666 666 666 d11r", "Lorem +48 666 666 XXX d11r", "Phone number in the middle of the sentence"},
	{"*", 3, "Lorem +48 666 666 666, +48 777 777 777 sit +88 888 888 888 amet", "Lorem +48 666 666 ***, +48 777 777 *** sit +88 888 888 *** amet", "multiple phone numbers"},
}

func TestPhoneAnonymize(t *testing.T) {
	for _, testCase := range phoneNumbersTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			pa := anonymizer.NewPhoneAnonymizer()
			pa.SetReplacement(testCase.replacement)
			err := pa.SetLastDigits(testCase.digitsAnonimizedCount)
			if err != nil {
				t.Errorf("Cannot set last digits count. Reason: '%s'", err.Error())
			}
			anonymized := pa.Anonymize(testCase.original)
			if anonymized != testCase.expected {
				t.Errorf("Anonymize(%s): expected %s, actual %s", testCase.original, testCase.expected, anonymized)
			}
		})
	}
}
