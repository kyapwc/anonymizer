package anonymizer

// Simple usage:
// o := anonymizer.NewOfferAnonymizer()
// e := anonymizer.NewEmailAnonymizer("***")
// s := anonymizer.NewSkypeAnonymizer("#")
// p := anonymizer.NewPhoneAnonymizer()
// o.AddAnonymizer(e)
// o.AddAnonymizer(s)
// o.AddAnonymizer(p)
// o.Anonymize(`Lorem ipsum a@a.com. <a href="skype:loremipsum?call">call</a> +48 666 777 888`)) -> Lorem ipsum ***@a.com. <a href="skype:#?call">call</a> +48 666 777 XXX

type offerAnonymizer struct {
	anonymizers []Anonymizer
}

func NewOfferAnonymizer() *offerAnonymizer {
	return &offerAnonymizer{}
}

func (oa *offerAnonymizer) AddAnonymizer(a Anonymizer) {
	oa.anonymizers = append(oa.anonymizers, a)
}

func (oa *offerAnonymizer) Anonymize(s string) string {
	for _, anonymizer := range oa.anonymizers {
		s = anonymizer.Anonymize(s)
	}
	return s
}
