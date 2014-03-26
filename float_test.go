package envar

import "testing"

func TestFloat(t *testing.T) {
	e := basicEnv{"PORT": "1000", "NUM_FAILURES": "wat", "FLOAT_SUCCESS": "1.10", "SCIENCE": "1e-5"}
	// copy over for the reference func version tests
	for k, v := range e {
		e["REF_"+k] = v
	}

	tests := []*struct {
		key          string
		defaultFloat float64
		expected     float64
		actual       float64
		actualRef    *float64
	}{
		{key: "PORT", defaultFloat: -1.0, expected: 1000},
		{key: "NUM_FAILURES", defaultFloat: -1.9, expected: -1.9},
		{key: "GOING_TO_BE_DEFAULT", defaultFloat: -100.021, expected: -100.021},
		{key: "FLOAT_SUCCESS", defaultFloat: -1000.123, expected: 1.10},
		{key: "SCIENCE", defaultFloat: -90.123, expected: 0.00001},
	}

	for _, test := range tests {
		FloatVar(&test.actual, test.key, test.defaultFloat, "Phony usage for float "+test.key)
		test.actualRef = Float("REF_"+test.key, test.defaultFloat, "Phony usage for float "+test.key+" REF")
	}

	err := ParseFromEnvironment(e)
	if err != nil {
		t.Errorf("Error processing environment err=%s", err)
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("Error with FloatVar. %q: got %f, expected %f", test.key, test.actual, test.expected)
		}

		if *test.actualRef != test.expected {
			t.Errorf("Error with Float. %q: got %f, expected %f", test.key, *test.actualRef, test.expected)
		}
	}
}
