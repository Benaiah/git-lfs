package humanize

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/git-lfs/git-lfs/errors"
)

const (
	Byte = 1 << (iota * 10)
	Kibibyte
	Mebibyte
	Gibibyte
	Tebibyte
	Pebibyte

	Kilobyte = 1000 * Byte
	Megabyte = 1000 * Kilobyte
	Gigabyte = 1000 * Megabyte
	Terabyte = 1000 * Gigabyte
	Petabyte = 1000 * Terabyte
)

var bytesTable = map[string]uint64{
	"b": Byte,

	"kib": Kibibyte,
	"mib": Mebibyte,
	"gib": Gibibyte,
	"tib": Tebibyte,
	"pib": Pebibyte,

	"kb": Kilobyte,
	"mb": Megabyte,
	"gb": Gigabyte,
	"tb": Terabyte,
	"pb": Petabyte,
}

// ParseBytes parses a given human-readable bytes or ibytes string into a number
// of bytes, or an error if the string was unable to be parsed.
func ParseBytes(str string) (uint64, error) {
	var sep int
	for _, r := range str {
		if !(unicode.IsDigit(r) || r == '.' || r == ',') {
			break
		}

		sep = sep + 1
	}

	f, err := strconv.ParseFloat(strings.Replace(str[:sep], ",", "", -1), 64)
	if err != nil {
		return 0, err
	}

	unit := strings.ToLower(strings.TrimSpace(str[sep:]))

	if m, ok := bytesTable[unit]; ok {
		f = f * float64(m)
		if f >= math.MaxUint64 {
			return 0, errors.New("number of bytes too large")
		}
		return uint64(f), nil
	}
	return 0, errors.Errorf("unknown unit: %q", unit)
}
