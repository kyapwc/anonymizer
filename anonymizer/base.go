package anonymizer

type Anonymizer interface {
	Anonymize(text string) string
}
