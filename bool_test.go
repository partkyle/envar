package envar

import "testing"

func TestBool(t *testing.T) {
	e := basicEnv{
		"INVALID_INTEGER": "1000",
		"FAILS":           "wat",
		"ZERO":            "0",
		"ONE":             "1",
		"SHOULD_FALSE":    "false",
		"SHOULD_TRUE":     "true",
	}

	// copy over for the reference func version tests
	for k, v := range e {
		e["REF_"+k] = v
	}

	tests := []*struct {
		key         string
		defaultBool bool
		expected    bool
		actual      bool
		actualRef   *bool
	}{
		{key: "INVALID_INTEGER", defaultBool: true, expected: true},
		{key: "FAILS", defaultBool: true, expected: true},
		{key: "ZERO", defaultBool: true, expected: false},
		{key: "ONE", defaultBool: true, expected: true},
		{key: "SHOULD_FALSE", defaultBool: true, expected: false},
		{key: "SHOULD_TRUE", defaultBool: true, expected: true},
		{key: "MISSING_WILL_DEFAULT", defaultBool: true, expected: true},

		// check inversion of default since there are only 2 cases
		{key: "INVALID_INTEGER", defaultBool: false, expected: false},
		{key: "FAILS", defaultBool: false, expected: false},
		{key: "ZERO", defaultBool: false, expected: false},
		{key: "ONE", defaultBool: false, expected: true},
		{key: "SHOULD_FALSE", defaultBool: false, expected: false},
		{key: "SHOULD_TRUE", defaultBool: false, expected: true},
		{key: "MISSING_WILL_DEFAULT", defaultBool: false, expected: false},
	}

	for _, test := range tests {
		BoolVar(&test.actual, test.key, test.defaultBool)
		test.actualRef = Bool("REF_"+test.key, test.defaultBool)
	}

	err := ParseFromEnvironment(e)
	if err != nil {
		t.Errorf("Error processing environment err=%s", err)
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("Error with BoolVar. %q: got %v, expected %v", test.key, test.actual, test.expected)
		}

		if *test.actualRef != test.expected {
			t.Errorf("Error with Bool. %q: got %v, expected %v", test.key, *test.actualRef, test.expected)
		}
	}
}
