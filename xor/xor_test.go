// Package xor does
package xor

import (
	"bytes"
	"crypto/sha256"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/penguingovernor/goxor/protocol"
)

func TestGenerateKey(t *testing.T) {
	shaData := sha256.Sum256([]byte("goxor"))
	testData := shaData[:]

	type args struct {
		key       []byte
		signature []byte
	}
	tests := []struct {
		name string
		args args
		want *protocol.Key
	}{
		{
			name: "Valid Generate Key",
			args: args{
				key:       []byte("key"),
				signature: []byte("goxor"),
			},
			want: &protocol.Key{
				PayLoad:   []byte("key"),
				Signature: testData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateKey(tt.args.key, tt.args.signature); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	shaData := sha256.Sum256([]byte("goxor"))
	testData := shaData[:]
	type args struct {
		data      []byte
		signature []byte
	}
	tests := []struct {
		name string
		args args
		want *protocol.Data
	}{
		{
			name: "Valid Generate Data",
			args: args{
				data:      []byte("key"),
				signature: []byte("goxor"),
			},
			want: &protocol.Data{
				PayLoad:   []byte("key"),
				Signature: testData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateData(tt.args.data, tt.args.signature); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {

	data := GenerateData([]byte("a"), []byte("b"))
	key := GenerateKey([]byte("b"), []byte("b"))
	badKey := GenerateKey([]byte("b"), []byte("c"))

	// "a" ^ "b" = 97 ^ 98 = 3
	dataXorKey := GenerateData([]byte(string(rune(3))), []byte("b"))

	type args struct {
		data *protocol.Data
		key  *protocol.Key
	}
	tests := []struct {
		name    string
		args    args
		want    *protocol.Data
		wantErr bool
	}{
		{
			name:    "Valid Xor",
			args:    args{data, key},
			want:    dataXorKey,
			wantErr: false,
		},
		{
			name:    "Invalid Xor",
			args:    args{data, badKey},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.data, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	decryptedData := GenerateData([]byte("a"), []byte("b"))
	key := GenerateKey([]byte("b"), []byte("b"))
	badKey := GenerateKey([]byte("b"), []byte("c"))

	// "a" ^ "b" = 97 ^ 98 = 3
	data := GenerateData([]byte(string(rune(3))), []byte("b"))
	type args struct {
		data *protocol.Data
		key  *protocol.Key
	}
	tests := []struct {
		name    string
		args    args
		want    *protocol.Data
		wantErr bool
	}{
		{
			name:    "Valid decrypt",
			args:    args{data, key},
			want:    decryptedData,
			wantErr: false,
		},
		{
			name:    "Invalid decrypt",
			args:    args{data, badKey},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.data, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		x []byte
		y []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2 Nil slices",
			args: args{nil, nil},
			want: true,
		},
		{
			name: "1 nil, 1 valid",
			args: args{nil, []byte("h")},
			want: false,
		},
		{
			name: "Different length slices",
			args: args{[]byte("Hello"), []byte("World!")},
			want: false,
		},
		{
			name: "Same length slices, diff content",
			args: args{[]byte("four"), []byte("five")},
			want: false,
		},
		{
			name: "Smae length slices, same content",
			args: args{[]byte("four"), []byte("four")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	message := GenerateKey([]byte("Test"), []byte("signature"))
	slice, err := proto.Marshal(message)
	if err != nil {
		t.Error("Could not marshall data")
	}

	type args struct {
		message proto.Message
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "Successful Write",
			args: args{
				message: message,
			},
			wantW:   string(slice),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Write(w, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Write() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestWriteKey(t *testing.T) {
	message := GenerateKey([]byte("Test"), []byte("signature"))
	slice, err := proto.Marshal(message)
	if err != nil {
		t.Error("Could not marshall data")
	}

	type args struct {
		key *protocol.Key
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "Successful Write",
			args: args{
				key: message,
			},
			wantW:   string(slice),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := WriteKey(w, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("WriteKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteKey() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	message := GenerateData([]byte("Test"), []byte("signature"))
	slice, err := proto.Marshal(message)
	if err != nil {
		t.Error("Could not marshall data")
	}
	type args struct {
		data *protocol.Data
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "Successful Write",
			args: args{
				data: message,
			},
			wantW:   string(slice),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := WriteData(w, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteData() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
