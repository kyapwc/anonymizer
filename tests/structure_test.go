package tests

import (
	"fmt"
	"gitlab.com/devskiller-tasks/golang-anonymizer/anonymizer"
	"reflect"
	"runtime"
	"testing"
)

var anonymizers = []struct {
	name       string
	anonymizer anonymizer.Anonymizer
}{
	{"email", anonymizer.NewEmailAnonymizer("...")},
	{"skype", anonymizer.NewSkypeAnonymizer("...")},
	{"phone", anonymizer.NewPhoneAnonymizer()},
	{"offer", anonymizer.NewOfferAnonymizer()},
}

func TestAnonymizersStructure(t *testing.T) {
	structureTests := []func(a anonymizer.Anonymizer, t *testing.T){testAnonymizerImplementsInterface, testAnonymizerReturnType, testAnonymizerParameterType}
	for _, anonymizerTested := range anonymizers {
		for _, testcase := range structureTests {
			testMethodName := runtime.FuncForPC(reflect.ValueOf(testcase).Pointer()).Name()
			t.Run(fmt.Sprintf("Anonimizer: '%s', method: '%s'", anonymizerTested.name, testMethodName), func(t *testing.T) {
				testcase(anonymizerTested.anonymizer, t)
			})
		}
	}
}
func testAnonymizerImplementsInterface(a anonymizer.Anonymizer, t *testing.T) {
	interfaceType := reflect.TypeOf((*anonymizer.Anonymizer)(nil)).Elem()
	if !reflect.TypeOf(a).Implements(interfaceType) {
		t.Error("EmailAnonymizer doesn't implement Anonymizer interface")
	}
}

func testAnonymizerReturnType(a anonymizer.Anonymizer, t *testing.T) {
	anonymized := a.Anonymize("string")
	kind := reflect.TypeOf(anonymized).Kind()
	if kind != reflect.String {
		t.Error("EmailAnonymizer return type is not string")
	}
}

func testAnonymizerParameterType(a anonymizer.Anonymizer, t *testing.T) {
	anonymizeMethod := reflect.TypeOf(a.Anonymize)
	if anonymizeMethod.NumIn() != 1 {
		t.Error("EmailAnonymizer.Anonymize expects only a single parameter")
	}
	if anonymizeMethod.In(0).Kind() != reflect.String {
		t.Error("EmailAnonymizer.Anonymize should expect only a string parameter")
	}
}
