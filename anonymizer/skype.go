package anonymizer

// Simple usage:
// NewSkypeAnonymizer("#").Anonymize(`<a href="skype:loren?call"/>`) -> <a href="skype:#?call"/>

type skypeAnonymizer struct {
	Replacement string
}

func NewSkypeAnonymizer(replacement string) *skypeAnonymizer {
	return &skypeAnonymizer{
		Replacement: replacement,
	}
}

func (sa *skypeAnonymizer) Anonymize(s string) string {
	// TODO: Return text with anonymized skype
	return ""
}
