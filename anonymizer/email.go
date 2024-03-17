package anonymizer

// Simple usage:
// ea := NewEmailAnonymizer("***")
// ea.Anonymize("a-a@a.b.c") -> ***@a.b.c

type emailAnonymizer struct {
	Replacement string
}

func NewEmailAnonymizer(replacement string) *emailAnonymizer {
	return &emailAnonymizer{
		Replacement: replacement,
	}
}

func (ea *emailAnonymizer) Anonymize(s string) string {
	// TODO: Return text with anonymized email
	return ""
}
