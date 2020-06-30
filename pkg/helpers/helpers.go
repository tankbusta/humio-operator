package helpers

import (
	"crypto/sha256"
	"fmt"
	"os"
	"reflect"

	humioapi "github.com/humio/cli/api"
)

func GetTypeName(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func ContainsElement(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func RemoveElement(list []string, s string) []string {
	for i, v := range list {
		if v == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// TODO: refactor, this is copied from the humio/cli/api/parsers.go
func MapTests(vs []string, f func(string) humioapi.ParserTestCase) []humioapi.ParserTestCase {
	vsm := make([]humioapi.ParserTestCase, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// TODO: refactor, this is copied from the humio/cli/api/parsers.go
func ToTestCase(line string) humioapi.ParserTestCase {
	return humioapi.ParserTestCase{
		Input:  line,
		Output: map[string]string{},
	}
}

// IsOpenShift returns whether the operator is running in OpenShift-mode
func IsOpenShift() bool {
	sccName, found := os.LookupEnv("OPENSHIFT_SCC_NAME")
	return found && sccName != ""
}

// AsSHA256 does a sha 256 hash on an object and returns the result
func AsSHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
