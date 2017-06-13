package humanize_test

import (
	"math"
	"testing"

	"github.com/git-lfs/git-lfs/tools/humanize"
	"github.com/stretchr/testify/assert"
)

type ParseBytesTestCase struct {
	Given    string
	Expected uint64
	Err      error
}

func (c *ParseBytesTestCase) Assert(t *testing.T) {
	got, err := humanize.ParseBytes(c.Given)
	if c.Err == nil {
		assert.NoError(t, err, "unexpected error: %s", err)
		assert.EqualValues(t, c.Expected, got)
	} else {
		assert.Equal(t, c.Err, err)
	}
}

func TestParseBytes(t *testing.T) {
	for desc, c := range map[string]*ParseBytesTestCase{
		"parse byte":     {"10B", uint64(10 * math.Pow(2, 0)), nil},
		"parse kibibyte": {"20KIB", uint64(20 * math.Pow(2, 10)), nil},
		"parse mebibyte": {"30MIB", uint64(30 * math.Pow(2, 20)), nil},
		"parse gibibyte": {"40GIB", uint64(40 * math.Pow(2, 30)), nil},
		"parse tebibyte": {"50TIB", uint64(50 * math.Pow(2, 40)), nil},
		"parse pebibyte": {"60PIB", uint64(60 * math.Pow(2, 50)), nil},

		"parse byte (lowercase)":     {"10b", uint64(10 * math.Pow(2, 0)), nil},
		"parse kibibyte (lowercase)": {"20kib", uint64(20 * math.Pow(2, 10)), nil},
		"parse mebibyte (lowercase)": {"30mib", uint64(30 * math.Pow(2, 20)), nil},
		"parse gibibyte (lowercase)": {"40gib", uint64(40 * math.Pow(2, 30)), nil},
		"parse tebibyte (lowercase)": {"50tib", uint64(50 * math.Pow(2, 40)), nil},
		"parse pebibyte (lowercase)": {"60pib", uint64(60 * math.Pow(2, 50)), nil},

		"parse byte (with space)":     {"10 B", uint64(10 * math.Pow(2, 0)), nil},
		"parse kibibyte (with space)": {"20 KIB", uint64(20 * math.Pow(2, 10)), nil},
		"parse mebibyte (with space)": {"30 MIB", uint64(30 * math.Pow(2, 20)), nil},
		"parse gibibyte (with space)": {"40 GIB", uint64(40 * math.Pow(2, 30)), nil},
		"parse tebibyte (with space)": {"50 TIB", uint64(50 * math.Pow(2, 40)), nil},
		"parse pebibyte (with space)": {"60 PIB", uint64(60 * math.Pow(2, 50)), nil},

		"parse byte (with space, lowercase)":     {"10 b", uint64(10 * math.Pow(2, 0)), nil},
		"parse kibibyte (with space, lowercase)": {"20 kib", uint64(20 * math.Pow(2, 10)), nil},
		"parse mebibyte (with space, lowercase)": {"30 mib", uint64(30 * math.Pow(2, 20)), nil},
		"parse gibibyte (with space, lowercase)": {"40 gib", uint64(40 * math.Pow(2, 30)), nil},
		"parse tebibyte (with space, lowercase)": {"50 tib", uint64(50 * math.Pow(2, 40)), nil},
		"parse pebibyte (with space, lowercase)": {"60 pib", uint64(60 * math.Pow(2, 50)), nil},

		"parse kilobyte": {"20KB", uint64(20 * math.Pow(10, 3)), nil},
		"parse megabyte": {"30MB", uint64(30 * math.Pow(10, 6)), nil},
		"parse gigabyte": {"40GB", uint64(40 * math.Pow(10, 9)), nil},
		"parse terabyte": {"50TB", uint64(50 * math.Pow(10, 12)), nil},
		"parse petabyte": {"60PB", uint64(60 * math.Pow(10, 15)), nil},

		"parse kilobyte (lowercase)": {"20kb", uint64(20 * math.Pow(10, 3)), nil},
		"parse megabyte (lowercase)": {"30mb", uint64(30 * math.Pow(10, 6)), nil},
		"parse gigabyte (lowercase)": {"40gb", uint64(40 * math.Pow(10, 9)), nil},
		"parse terabyte (lowercase)": {"50tb", uint64(50 * math.Pow(10, 12)), nil},
		"parse petabyte (lowercase)": {"60pb", uint64(60 * math.Pow(10, 15)), nil},

		"parse kilobyte (with space)": {"20 KB", uint64(20 * math.Pow(10, 3)), nil},
		"parse megabyte (with space)": {"30 MB", uint64(30 * math.Pow(10, 6)), nil},
		"parse gigabyte (with space)": {"40 GB", uint64(40 * math.Pow(10, 9)), nil},
		"parse terabyte (with space)": {"50 TB", uint64(50 * math.Pow(10, 12)), nil},
		"parse petabyte (with space)": {"60 PB", uint64(60 * math.Pow(10, 15)), nil},

		"parse kilobyte (with space, lowercase)": {"20 kb", uint64(20 * math.Pow(10, 3)), nil},
		"parse megabyte (with space, lowercase)": {"30 mb", uint64(30 * math.Pow(10, 6)), nil},
		"parse gigabyte (with space, lowercase)": {"40 gb", uint64(40 * math.Pow(10, 9)), nil},
		"parse terabyte (with space, lowercase)": {"50 tb", uint64(50 * math.Pow(10, 12)), nil},
		"parse petabyte (with space, lowercase)": {"60 pb", uint64(60 * math.Pow(10, 15)), nil},
	} {
		t.Run(desc, c.Assert)
	}
}
