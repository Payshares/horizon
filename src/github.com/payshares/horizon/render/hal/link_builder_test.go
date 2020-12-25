package hal

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestLinkBuilder(t *testing.T) {

	Convey("Link Expansion", t, func() {

		check := func(href string, base string, expectedResult string) {
			lb := LinkBuilder{mustParseURL(base)}
			result := lb.expandLink(href)
			So(result, ShouldEqual, expectedResult)
		}

		check("/root", "", "/root")
		check("/root", "//payshares.org", "//payshares.org/root")
		check("/root", "https://payshares.org", "https://payshares.org/root")
		check("//else.org/root", "", "//else.org/root")
		check("//else.org/root", "//payshares.org", "//else.org/root")
		check("//else.org/root", "https://payshares.org", "//else.org/root")
		check("https://else.org/root", "", "https://else.org/root")
		check("https://else.org/root", "//payshares.org", "https://else.org/root")
		check("https://else.org/root", "https://payshares.org", "https://else.org/root")

		// Regression: ensure that parameters are not escaped
		check("/accounts/{id}", "https://payshares.org", "https://payshares.org/accounts/{id}")
	})

}

func mustParseURL(base string) *url.URL {
	if base == "" {
		return nil
	}

	u, err := url.Parse(base)
	if err != nil {
		panic(err)
	}
	return u
}
