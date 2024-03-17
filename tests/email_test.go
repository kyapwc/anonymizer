package tests

import (
	"gitlab.com/devskiller-tasks/golang-anonymizer/anonymizer"
	"testing"
)

var mailAnonymizerTestCases = []struct {
	original    string
	expected    string
	description string
}{
	{"Lorem ipsum", "Lorem ipsum", "no email"},
	{"Lorem ipsum a@a.com dolor sit amet", "Lorem ipsum ...@a.com dolor sit amet", "single small letter character mail name"},
	{"Lorem ipsum a@a.com dolor sit amet a@a.com ber", "Lorem ipsum ...@a.com dolor sit amet ...@a.com ber", "2 email adresses"},
	{"Lorem ipsum --@--.--", "Lorem ipsum --@--.--", "mail name without recognizable characters"},
	{"Lorem some@data ipsum", "Lorem some@data ipsum", "malformed mail - mail DNS is malformed, no anonimization expected"},
	{"Lorem B@bb12.com ipsum", "Lorem ...@bb12.com ipsum", "single big letter character mail name"},
	{"Lorem abc-abc@abc.edu.co.uk am", "Lorem ...@abc.edu.co.uk am", "multipart mail DNS"},
	{"Lorem cBa-abC@abc.edu.co.uk. dolor", "Lorem ...@abc.edu.co.uk. dolor", "mail DNS ending with a dot"},
	{"Lorem dsad BB12@BB-12.COM. dolor", "Lorem dsad ...@BB-12.COM. dolor", "big letters with numbers"},
	{"Lorem XXd -abc_ABC@abc.edu. dolor", "Lorem XXd -...@abc.edu. dolor", "Leading special character ('-') is not anonymized"},
	{"Lorem XXd abc_ABC-@abc.edu. dolor", "Lorem XXd abc_ABC-@abc.edu. dolor", "Wrong username (special character just before @)"},
	{"Lorem b12@bb12.com- dolor", "Lorem ...@bb12.com- dolor", "mail DNS ending special unrecognised characters omitted"},
}

func TestEmailAnonymize(t *testing.T) {
	emailAnonymizer := anonymizer.NewEmailAnonymizer("...")
	for _, test := range mailAnonymizerTestCases {
		t.Run(test.description, func(t *testing.T) {
			anonymized := emailAnonymizer.Anonymize(test.original)
			if anonymized != test.expected {
				t.Errorf("Anonymize(%s): expected '%s', actual '%s'", test.original, test.expected, anonymized)
			}
		})
	}
}
