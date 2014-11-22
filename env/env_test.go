package env

import (
	"os"
	"strconv"
	"testing"
)

func TestBool(t *testing.T) {
	in, out := false, true

	os.Setenv("ENVBOOL", "true")

	if got := Bool("ENVBOOL", in); got != out {
		t.Errorf(`Bool("ENVBOOL", %v) = %v, want %v`, in, got, out)
	}
}

func TestBoolDefault(t *testing.T) {
	in, out := true, true

	if got := Bool("ENVBOOL_DEFAULT", in); got != out {
		t.Errorf(`Bool("ENVBOOL_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestInt(t *testing.T) {
	in, out := 1, 2

	os.Setenv("ENVINT", strconv.Itoa(out))

	if got := Int("ENVINT", in); got != out {
		t.Errorf(`Int("ENVINT", %v) = %v, want %v`, in, got, out)
	}
}

func TestIntDefault(t *testing.T) {
	in, out := 3, 3

	if got := Int("ENVINT_DEFAULT", in); got != out {
		t.Errorf(`Int("ENVINT_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestString(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("ENVSTR", out)

	if got := String("ENVSTR", in); got != out {
		t.Errorf(`String("ENVSTR", "%v") = %v, want %v`, in, got, out)
	}
}

func TestStringDefault(t *testing.T) {
	in, out := "baz", "baz"

	if got := String("ENVSTR_DEFAULT", in); got != out {
		t.Errorf(`String("ENVSTR_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}
