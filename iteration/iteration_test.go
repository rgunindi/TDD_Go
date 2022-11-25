package iteration

import "testing"

func TestIteration(t *testing.T){
got:=Iterate("a",5)
expected:="aaaaa"

if got!=expected {
	t.Errorf("Expected %q but got %q",expected,got)
}

}