package durationpb

/*
durationpbGormAdapter.go
implements all interfaces for GORM to serialize/deseialize the Duration-Type to a bigint value
value is represented as milliseconds
*/

import (
	"math/big"
	time "time"

	"database/sql"
	"database/sql/driver"
)

//Scan converts sql-time to Duration (used by GORM)
func (x *Duration) Scan(value interface{}) (err error) {
	nullInt := &sql.NullInt64{}
	err = nullInt.Scan(value)
	*x = *New(time.Duration(nullInt.Int64))
	return
}

//Value returns Duration in nanoseconds (used by GORM)
func (x *Duration) Value() (driver.Value, error) {
	return x.GetSeconds()*int64(time.Second) + int64(x.GetNanos())*int64(time.Nanosecond), nil
}

//GormDataType gorm common data type
func (x *Duration) GormDataType() string {
	return "bigint"
}

//GobEncode encodes Duration to byte-slice (used by GORM)
func (x *Duration) GobEncode() ([]byte, error) {
	tmp := big.NewInt(x.GetSeconds()*int64(time.Second) + int64(x.GetNanos())*int64(time.Nanosecond))
	return tmp.GobEncode()
}

//GobDecode decodes the byte-slice to Duration (used by GORM)
func (x *Duration) GobDecode(b []byte) error {
	tmp := big.Int{}
	err := tmp.GobDecode(b)
	x = New(time.Duration(tmp.Int64()))
	return err
}
