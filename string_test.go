package envar

import "testing"

func TestString(t *testing.T) {
	e := basicEnv{"HOST": "1000", "NUM_FAILURES": "wat",
		"REF_HOST": "1000", "REF_NUM_FAILURES": "wat"}

	tests := []*struct {
		key       string
		expected  string
		actual    string
		actualRef *string
	}{
		{key: "HOST", expected: "1000"},
		{key: "NUM_FAILURES", expected: "wat"},
	}

	for _, test := range tests {
		StringVar(&test.actual, test.key)
		test.actualRef = String("REF_" + test.key)
	}

	err := ParseFromEnvironment(e)
	if err != nil {
		t.Errorf("Error processing environment err=%s", err)
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("Error with StringVar. %q: got %v, expected %v", test.key, test.actual, test.expected)
		}

		if *test.actualRef != test.expected {
			t.Errorf("Error with String. %q: got %v, expected %v", test.key, *test.actualRef, test.expected)
		}
	}
}
