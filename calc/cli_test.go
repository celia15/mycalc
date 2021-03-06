package mycalc

import (
	"bytes"
	"testing"
)

type runTest struct {
	name  string
	input string
	value string
}

var runTests = []runTest{
	{"single", "1", "1\n"},
	{"add", "1+1", "2\n"},
	{"sub", "1-1", "0\n"},
	{"mul", "2*3", "6\n"},
	{"div", "4/2", "2\n"},
	{"addmul", "4+2*5", "14\n"},
	{"subdiv", "6-4/2", "4\n"},
	{"twoline", "1+2\n3+4", "3\n7\n"},
	{"op", "+", "unexpected item\n"},
	{"op and success", "+\n1+2", "unexpected item\n3\n"},
	{"not expected symbol", "&", "unexpected item\n"},
	{"paren", "(1)", "1\n"},
	{"paren calc", "(1+2)", "3\n"},
	{"paren calc mul", "(1+2)*3", "9\n"},
	{"unary minus", "-3", "-3\n"},
	{"unary minus", "1--3", "4\n"},
	{"sushi", "🍣", "980\n"},
	{"double sushi", "🍣+🍣", "1960\n"},
	{"assign", "answer=42\nanswer", "42\n42\n"},
	{"unassigned", "x", "unassigned variable x\n"},
	{"unassigned2", "2*pi", "unassigned variable pi\n"},
	{"unassigned3", "a=a", "unassigned variable a\n"},
	{"two number", "1 2", "1\n2\n"},
}

func TestRun(t *testing.T) {
	for _, test := range runTests {
		b := bytes.Buffer{}
		RunString(test.input, &b)
		sb := b.String()
		if sb != test.value {
			t.Errorf("%s: got\n\t%v\nexpeected\n\t%v", test.name, sb, test.value)
		}
	}
}
