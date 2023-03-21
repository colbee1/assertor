package assertor

import (
	"errors"
	"strings"
	"testing"
)

func TestAssertor_Assert(t *testing.T) {

	type args struct {
		ok     bool
		errMsg string
		args   []any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{ok: 2 == 2, errMsg: "right", args: []interface{}{}},
			want: true,
		},
		{
			args: args{ok: 2 == 3, errMsg: "wrong", args: []interface{}{}},
			want: false,
		},
	}
	a := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := a.Assert(tt.args.ok, tt.args.errMsg, tt.args.args...); got != tt.want {
				t.Errorf("Assertor.Assert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssertor_Validate(t *testing.T) {
	a := New()
	a.Assert(2 == 2, "ok")
	err := a.Validate()
	if err != nil {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, nil)
	}

	w1 := "wrong1"
	w2 := "wrong2"
	sep := "; "
	expected := ErrValidate.Error() + ": 2 unsatisfied requirement(s): " + strings.Join([]string{w1, w2}, sep)

	a.Assert(2 == 3, w1)
	a.Assert(2 == 4, w2)
	err = a.Validate()
	if !errors.Is(err, ErrValidate) {
		t.Errorf("errors.Is(ErrValidate) should be true")
	}
	if err.Error() != expected {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, expected)
	}

	sep = "-"
	a.Separator = sep
	expected = ErrValidate.Error() + ": 2 unsatisfied requirement(s): " + strings.Join([]string{w1, w2}, sep)
	err = a.Validate()
	if err.Error() != expected {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, expected)
	}

}
