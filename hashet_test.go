package hashet_test

import (
	"testing"

	. "github.com/marc-gr/hashet"
)

func TestNew(t *testing.T) {
	const length = 1
	want := "00"

	hash := New(length)
	if got := hash.String(); want != got {
		t.Errorf("New() hash = %s, want %s", got, want)
	}
}

func TestNewFromSet(t *testing.T) {
	type args struct {
		l   int
		set [][]byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "hash new from set throws error",
			args:    args{l: 1, set: [][]byte{{0, 2}}},
			wantErr: true,
		},
		{
			name: "hash new from set prints the correct hash",
			args: args{l: 2, set: [][]byte{{2, 1}, {1, 2}}},
			want: "0303",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := NewFromSet(tt.args.l, tt.args.set...)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewFromSet() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got := hash.String(); got != tt.want {
				t.Errorf("NewFromSet() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestRehash(t *testing.T) {
	type args struct {
		l   int
		set [][]byte
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantErr      bool
		wantMismatch int
	}{
		{
			name:         "hash rehash throws correct mismatch shorter error",
			args:         args{l: 1, set: [][]byte{{0, 2}}},
			wantErr:      true,
			wantMismatch: -1,
		},
		{
			name:         "hash rehash throws correct mismatch longer error",
			args:         args{l: 2, set: [][]byte{{0}}},
			wantErr:      true,
			wantMismatch: 1,
		},
		{
			name: "hash rehash prints the correct hash",
			args: args{l: 2, set: [][]byte{{1, 2}, {2, 2}}},
			want: "0300",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := New(tt.args.l)
			err := hash.Rehash(tt.args.set...)
			if err != nil {
				if tt.wantErr {
					if errMsm, ok := err.(Mismatch); !ok {
						t.Errorf("Rehash() error is not a Mismatch error: %T", err)
					} else if got := errMsm.Mismatch(); got != tt.wantMismatch {
						t.Errorf("Rehash() error mismatch = %d, mismatchValue %d", got, tt.wantMismatch)
					}
				} else {
					t.Errorf("Rehash() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got := hash.String(); got != tt.want {
				t.Errorf("Rehash() = %s, want %s", got, tt.want)
			}
		})
	}
}
