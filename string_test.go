package envar

import "testing"

func TestString(t *testing.T) {
	e := basicEnv{"HOST": "1000", "NUM_FAILURES": "wat"}
	// copy over for the reference func version tests
	for k, v := range e {
		e["REF_"+k] = v
	}

	tests := []*struct {
		key           string
		defaultString string
		expected      string
		actual        string
		actualRef     *string
	}{
		{key: "HOST", expected: "1000"},
		{key: "NUM_FAILURES", expected: "wat"},
		{key: "NOT_IN_ENV", defaultString: "woot", expected: "woot"},
	}

	for _, test := range tests {
		StringVar(&test.actual, test.key, test.defaultString, "Phony usage for string "+test.key)
		test.actualRef = String("REF_"+test.key, test.defaultString, "Phony usage for string "+test.key+" REF")
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
