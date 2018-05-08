package vhd

import (
	"encoding/binary"
	"reflect"
	"regexp"
	"testing"
)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
func Test_uuidgen(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uuidgen(); IsValidUUID(got) {
				t.Errorf("uuidgen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fmtField(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Hello World",
			args: args{
				name:  "Name",
				value: "Hello World",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmtField(tt.args.name, tt.args.value)
		})
	}
}

func Test_hexs(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args:args{
				a: []byte("test"),

			},
			want: "0x74657374",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexs(tt.args.a); got != tt.want {
				t.Errorf("hexs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uuid(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "default",
			args: args{
				a: []byte("820e9200-a5d2-7d8e-38f9-aa6b6938f9f7"),
			},
			want: "38323065-3932-3030-2d61-3564322d3764",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uuid(tt.args.a); got != tt.want {
				t.Errorf("uuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uuidToBytes(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{
			name: "Good",
			args: args{
				uuid: "38323065-3932-3030-2d61-3564322d3764",
			},
			wantOut: []byte{56, 50, 48, 101, 57, 50, 48, 48, 45, 97, 53, 100, 50, 45, 55, 100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := uuidToBytes(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("uuidToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("uuidToBytes() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_utf16BytesToString(t *testing.T) {
	type args struct {
		b []byte
		o binary.ByteOrder
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Done",
			args: args{
				b: []byte(""),
				o: binary.BigEndian,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utf16BytesToString(tt.args.b, tt.args.o); got != tt.want {
				t.Errorf("utf16BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
