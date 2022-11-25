package thaana_test

import (
	"testing"

	"github.com/yaameen/go-thaana"
)

func TestAscii(t *testing.T) {

	testCases := []struct {
		name        string
		test        string
		expectation string
		fn          func(string) string
	}{
		{
			name:        "converts ascii to unicode equivalent",
			test:        "Oleh",
			expectation: "ހެލޯ",
			fn:          thaana.AsciiToUnicode,
		},
		{
			name:        "converts ascii to unicode equivalent",
			test:        "udwmcawHum cnImWy",
			expectation: "ޔާމީން މުޙައްމަދު",
			fn:          thaana.AsciiToUnicode,
		},
		{
			name:        "thinki",
			test:        "udwmcawHum",
			expectation: "މުޙައްމަދު",
			fn:          thaana.AsciiToUnicode,
		},
		{
			name:        "arabic",
			test:        "Qcaudcbwa",
			expectation: "އަބްދުއްﷲ",
			fn:          thaana.AsciiToUnicode,
		},
		{
			name:        "with numbers",
			test:        "1 Q",
			expectation: "ﷲ 1",
			fn:          thaana.AsciiToUnicode,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			utf := v.fn(v.test)
			if utf != v.expectation {
				t.Errorf("expected %s but got %v", v.expectation, utf)
			}
		})
	}

}
