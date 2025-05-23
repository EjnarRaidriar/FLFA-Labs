package regex

import (
	"testing"
)

func Test1(t *testing.T) {
	test := "a|bc"
	compare := "(| a (_ b c))"
	tree, err := NewParseTree(test)
	if err != nil && tree.String() != compare {
		t.Errorf(`NewParseTree("%s") != %s`, test, compare)
	}
}

func Test2(t *testing.T) {
	test := "ab|c"
	compare := "(_ a (| b c))"
	tree, err := NewParseTree(test)
	if err != nil && tree.String() != compare {
		t.Errorf(`NewParseTree("%s") != %s`, test, compare)
	}
}

func Test3(t *testing.T) {
	test := "a+"
	compare := "(+ a)"
	tree, err := NewParseTree(test)
	if err != nil && tree.String() != compare {
		t.Errorf(`NewParseTree("%s") != %s`, test, compare)
	}
}

func Test4(t *testing.T) {
	test := "a+b"
	compare := "(_ (+ a) b)"
	tree, err := NewParseTree(test)
	if err != nil && tree.String() != compare {
		t.Errorf(`NewParseTree("%s") != %s`, test, compare)
	}
}

func Test5(t *testing.T) {
	test := "a^2"
	compare := "(^ a 2)"
	tree, err := NewParseTree(test)
	if err != nil && tree.String() != compare {
		t.Errorf(`NewParseTree("%s") != %s`, test, compare)
	}
}

func TestInvalid1(t *testing.T) {
	test := "a^b"
	_, err := NewParseTree(test)
	if err == nil {
		t.Errorf("%s should be invalid", test)
	}
}
