package proxmox

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CpuLimit_Validate(t *testing.T) {
	testData := []struct {
		name  string
		input CpuLimit
		err   error
	}{
		// Invalid
		{name: "Invalid errors.New(CpuLimit_Error_UpperBound)",
			input: 129,
			err:   errors.New(CpuLimit_Error_UpperBound),
		},
		// Valid
		{name: "Valid LowerBound",
			input: 0,
		},
		{name: "Valid UpperBound",
			input: 128,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(*testing.T) {
			require.Equal(t, test.input.Validate(), test.err, test.name)
		})
	}
}

func Test_CpuSockets_Validate(t *testing.T) {
	testData := []struct {
		name  string
		input CpuSockets
		err   error
	}{
		// Invalid
		{name: "Invalid errors.New(CpuSockets_Error_LowerBound)",
			input: 0,
			err:   errors.New(CpuSockets_Error_LowerBound),
		},
		{name: "Invalid errors.New(CpuSockets_Error_UpperBound)",
			input: 5,
			err:   errors.New(CpuSockets_Error_UpperBound),
		},
		// Valid
		{name: "Valid LowerBound",
			input: 1,
		},
		{name: "Valid UpperBound",
			input: 4,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(*testing.T) {
			require.Equal(t, test.input.Validate(), test.err, test.name)
		})
	}
}

func Test_CpuUnits_Validate(t *testing.T) {
	testData := []struct {
		name  string
		input CpuUnits
		err   error
	}{
		// Invalid
		{name: "Invalid errors.New(CpuUnits_Error_UpperBound)",
			input: 262145,
			err:   errors.New(CpuUnits_Error_UpperBound),
		},
		// Valid
		{name: "Valid LowerBound",
			input: 0,
		},
		{name: "Valid UpperBound",
			input: 262144,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(*testing.T) {
			require.Equal(t, test.input.Validate(), test.err, test.name)
		})
	}
}
