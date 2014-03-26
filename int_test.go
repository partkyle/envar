package envar

import "testing"

type env map[string]string

func (e env) Get(key string) string {
	return e[key]
}

func TestInt(t *testing.T) {
	e := env{"PORT": "1000", "NUM_FAILURES": "wat", "FLOAT_TO_FAIL": "1.10",
		"REF_PORT": "1000", "REF_NUM_FAILURES": "wat", "REF_FLOAT_TO_FAIL": "1.10"}

	tests := []*struct {
		key        string
		defaultInt int
		expected   int
		actual     int
		actualRef  *int
	}{
		{key: "PORT", defaultInt: -1, expected: 1000},
		{key: "NUM_FAILURES", defaultInt: -1, expected: -1},
		{key: "GOING_TO_BE_DEFAULT", defaultInt: -100, expected: -100},
		{key: "FLOAT_TO_FAIL", defaultInt: -1000, expected: -1000},
	}

	for _, test := range tests {
		IntVar(&test.actual, test.key, test.defaultInt)
		test.actualRef = Int("REF_"+test.key, test.defaultInt)
	}

	err := ParseFromEnvironment(e)
	if err != nil {
		t.Errorf("Error processing environment err=%s", err)
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("Error with IntVar. %q: got %d, expected %d", test.key, test.actual, test.expected)
		}

		if *test.actualRef != test.expected {
			t.Errorf("Error with Int. %q: got %d, expected %d", test.key, *test.actualRef, test.expected)
		}
	}
}
