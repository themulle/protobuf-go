package timestamppb

/*
timestampbGormAdapter.go
implements all interfaces for GORM to serialize/deseialize the Timestamp-Type to a date value
values are treated as regular time.Time types
*/

import (
	"encoding/json"
	time "time"

	"database/sql"
	"database/sql/driver"
)

// Scan converts sql-time to Timestamp (used by GORM)
func (x *Timestamp) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*x = *New(nullTime.Time)
	return
}

// Value returns Timestamp as Time for "regular" processing (used by GORM)
func (x *Timestamp) Value() (driver.Value, error) {
	return x.AsTime(), nil
}

// GormDataType gorm common data type
func (x *Timestamp) GormDataType() string {
	return "time"
}

// GobEncode encodes Tiemstamp to byte-slice (used by GORM)
func (x *Timestamp) GobEncode() ([]byte, error) {
	return x.AsTime().GobEncode()
}

// GobDecode decodes the byte-slice to Tiemstamp (used by GORM)
func (x *Timestamp) GobDecode(b []byte) error {
	t := time.Time{}
	err := t.GobDecode(b)
	x.Seconds = int64(t.Unix())
	x.Nanos = int32(t.Nanosecond())
	return err
}

func (x *Timestamp) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
		Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	}{}
	err1 := json.Unmarshal(data, aux)
	if err1 == nil {
		x.Seconds = aux.Seconds
		x.Nanos = aux.Nanos
	} else {
		var t time.Time
		if err := json.Unmarshal(data, &t); err == nil {
			x.Seconds = int64(t.Unix())
			x.Nanos = int32(t.Nanosecond())
			err1 = nil
		}
	}
	return err1
}
