package prose

import (
	"fmt"
	"testing"
)
func TestOneElements(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if want != got {
		t.Errorf(errString(list,got,want))
	}
}
func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errString(list,got,want))
	}
}
func TestThreeElement(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple,orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errString(list,got,want))
	}
	
}

func errString(list []string,got string,want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\",want \"%s\"",list,got,want)
}
