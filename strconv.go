package idgen

import "math/bits"
import "errors"

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func formatInt(i int64, base int) string {
	_, s := formatBits(nil, uint64(i), base, false, false)
	return s
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

const host32bit = ^uint(0)>>32 == 0

const digits = "abcdefghijklmnopqrstuvwxyz"

// formatBits computes the string representation of u in the given base.
// If neg is set, u is treated as negative int64 value. If append_ is
// set, the string is appended to dst and the resulting byte slice is
// returned as the first result value; otherwise the string is returned
// as the second result value.
//
func formatBits(dst []byte, u uint64, base int, neg, append_ bool) (d []byte, s string) {

	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [64 + 1]byte // +1 for sign of 64bit value in base 2
	i := len(a)

	if neg {
		u = -u
	}

	// convert bits
	// We use uint values where we can because those will
	// fit into a single register even on a 32bit machine.
	if base == 10 {
		// common case: use constants for / because
		// the compiler can optimize it into a multiply+shift

		if host32bit {
			// convert the lower digits using 32bit operations
			for u >= 1e9 {
				// Avoid using r = a%b in addition to q = a/b
				// since 64bit division and modulo operations
				// are calculated by runtime functions on 32bit machines.
				q := u / 1e9
				us := uint(u - q*1e9) // u % 1e9 fits into a uint
				for j := 4; j > 0; j-- {
					is := us % 100 * 2
					us /= 100
					i -= 2
					a[i+1] = smallsString[is+1]
					a[i+0] = smallsString[is+0]
				}

				// us < 10, since it contains the last digit
				// from the initial 9-digit us.
				i--
				a[i] = smallsString[us*2+1]

				u = q
			}
			// u < 1e9
		}

		// u guaranteed to fit into a uint
		us := uint(u)
		for us >= 100 {
			is := us % 100 * 2
			us /= 100
			i -= 2
			a[i+1] = smallsString[is+1]
			a[i+0] = smallsString[is+0]
		}

		// us < 100
		is := us * 2
		i--
		a[i] = smallsString[is+1]
		if us >= 10 {
			i--
			a[i] = smallsString[is]
		}

	} else if isPowerOfTwo(base) {
		// It is known that base is a power of two and
		// 2 <= base <= len(digits).
		// Use shifts and masks instead of / and %.
		shift := uint(bits.TrailingZeros(uint(base))) & 31
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u >= b {
			i--
			a[i] = digits[uint(u)&m]
			u >>= shift
		}
		// u < base
		i--
		a[i] = digits[uint(u)]
	} else {
		// general case
		b := uint64(base)
		for u >= b {
			i--
			// Avoid using r = a%b in addition to q = a/b
			// since 64bit division and modulo operations
			// are calculated by runtime functions on 32bit machines.
			q := u / b
			a[i] = digits[uint(u-q*b)]
			u = q
		}
		// u < base
		i--
		a[i] = digits[uint(u)]
	}

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}

	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

func parseInt(s string, base int, bitSize int) (i int64, err error) {
	const fnParseInt = "ParseInt"

	// Empty string bad.
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}

	// Pick off leading sign.
	s0 := s
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	// Convert unsigned and check range.
	var un uint64
	un, err = ParseUint(s, base, bitSize)
	if err != nil && err.(*NumError).Err != ErrRange {
		err.(*NumError).Func = fnParseInt
		err.(*NumError).Num = s0
		return 0, err
	}

	if bitSize == 0 {
		bitSize = int(IntSize)
	}

	cutoff := uint64(1 << uint(bitSize-1))
	if !neg && un >= cutoff {
		return int64(cutoff - 1), rangeError(fnParseInt, s0)
	}
	if neg && un > cutoff {
		return -int64(cutoff), rangeError(fnParseInt, s0)
	}
	n := int64(un)
	if neg {
		n = -n
	}
	return n, nil
}

func ParseUint(s string, base int, bitSize int) (uint64, error) {
	const fnParseUint = "ParseUint"

	if len(s) == 0 {
		return 0, syntaxError(fnParseUint, s)
	}

	s0 := s
	switch {
	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal, hex prefix.
		switch {
		case s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X'):
			if len(s) < 3 {
				return 0, syntaxError(fnParseUint, s0)
			}
			base = 16
			s = s[2:]
		case s[0] == '0':
			base = 8
			s = s[1:]
		default:
			base = 10
		}

	default:
		return 0, baseError(fnParseUint, s0, base)
	}

	if bitSize == 0 {
		bitSize = int(IntSize)
	} else if bitSize < 0 || bitSize > 64 {
		return 0, bitSizeError(fnParseUint, s0, bitSize)
	}

	// Cutoff is the smallest number such that cutoff*base > maxUint64.
	// Use compile-time constants for common cases.
	var cutoff uint64
	switch base {
	case 10:
		cutoff = maxUint64/10 + 1
	case 16:
		cutoff = maxUint64/16 + 1
	default:
		cutoff = maxUint64/uint64(base) + 1
	}

	maxVal := uint64(1)<<uint(bitSize) - 1

	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case 'a' <= c && c <= 'z':
			d = c - 'a'
		default:
			return 0, syntaxError(fnParseUint, s0)
		}

		if d >= byte(base) {
			return 0, syntaxError(fnParseUint, s0)
		}

		if n >= cutoff {
			// n*base overflows
			return maxVal, rangeError(fnParseUint, s0)
		}
		n *= uint64(base)

		n1 := n + uint64(d)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			return maxVal, rangeError(fnParseUint, s0)
		}
		n = n1
	}

	return n, nil
}

// A NumError records a failed conversion.
type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing: " + e.Err.Error()
}

func syntaxError(fn, str string) *NumError {
	return &NumError{fn, str, ErrSyntax}
}

func rangeError(fn, str string) *NumError {
	return &NumError{fn, str, ErrRange}
}

func baseError(fn, str string, base int) *NumError {
	return &NumError{fn, str, errors.New("invalid base")}
}

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

func bitSizeError(fn, str string, bitSize int) *NumError {
	return &NumError{fn, str, errors.New("invalid bit size")}
}

const intSize = 32 << (^uint(0) >> 63)

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

const maxUint64 = (1<<64 - 1)
