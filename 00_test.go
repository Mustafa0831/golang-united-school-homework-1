package solution

import (
	"strings"
	"testing"
)

func Test0(t *testing.T) {
	want := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	check := PrintS()
	// fmt.Println(check,"here")
	if !strings.EqualFold(want, check) {
		t.Errorf("Msg got was incorrect, check: %s, want: %s", want, check)

	}
}
