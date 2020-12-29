package pulid

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/oklog/ulid"
)

func TestID_String(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{Prefix: "TT"}, "TT00000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			if got := id.String(); got != tt.want {
				t.Errorf("ID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		pulid string
	}
	tests := []struct {
		name    string
		args    args
		wantId  ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := Parse(tt.args.pulid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotId, tt.wantId) {
				t.Errorf("Parse() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestID_MarshalText(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			got, err := id.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_UnmarshalText(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	type args struct {
		v []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			if err := id.UnmarshalText(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ID.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_MarshalBinary(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			got, err := id.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_MarshalBinaryTo(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	type args struct {
		dst []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			if err := id.MarshalBinaryTo(tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("ID.MarshalBinaryTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_UnmarshalBinary(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			if err := id.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ID.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_Scan(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			if err := id.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("ID.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_Value(t *testing.T) {
	type fields struct {
		ULID   ulid.ULID
		Prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := ID{
				ULID:   tt.fields.ULID,
				Prefix: tt.fields.Prefix,
			}
			got, err := id.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalID(t *testing.T) {
	type args struct {
		id ID
	}
	tests := []struct {
		name string
		args args
		want graphql.Marshaler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarshalID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalID(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantId  ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := UnmarshalID(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotId, tt.wantId) {
				t.Errorf("UnmarshalID() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
