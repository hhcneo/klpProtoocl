// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package klp

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString


const (
  MessageType_kNtfLogin MessageType = 1
  MessageType_kNtfLogout MessageType = 2
  MessageType_kNtfChangeLocation MessageType = 3
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kNtfLogin: return "kNtfLogin"
  case MessageType_kNtfLogout: return "kNtfLogout"
  case MessageType_kNtfChangeLocation: return "kNtfChangeLocation"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kNtfLogin": return MessageType_kNtfLogin, nil 
  case "kNtfLogout": return MessageType_kNtfLogout, nil 
  case "kNtfChangeLocation": return MessageType_kNtfChangeLocation, nil 
  }
  return MessageType(0), fmt.Errorf("not a valid MessageType string")
}


func MessageTypePtr(v MessageType) *MessageType { return &v }

func (p MessageType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *MessageType) UnmarshalText(text []byte) error {
q, err := MessageTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *MessageType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = MessageType(v)
return nil
}

func (p * MessageType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - Ssn
//  - Gsn
//  - TraceID
type NtfLogin struct {
  Ssn Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Gsn Gsn `thrift:"gsn,2,required" db:"gsn" json:"gsn"`
  TraceID Buffer `thrift:"trace_id,3,required" db:"trace_id" json:"trace_id"`
}

func NewNtfLogin() *NtfLogin {
  return &NtfLogin{}
}


func (p *NtfLogin) GetSsn() Ssn {
  return p.Ssn
}

func (p *NtfLogin) GetGsn() Gsn {
  return p.Gsn
}

func (p *NtfLogin) GetTraceID() Buffer {
  return p.TraceID
}
func (p *NtfLogin) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetGsn bool = false;
  var issetTraceID bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetSsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetGsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetTraceID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetSsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ssn is not set"));
  }
  if !issetGsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Gsn is not set"));
  }
  if !issetTraceID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TraceID is not set"));
  }
  return nil
}

func (p *NtfLogin)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *NtfLogin)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := Gsn(v)
  p.Gsn = temp
}
  return nil
}

func (p *NtfLogin)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := Buffer(v)
  p.TraceID = temp
}
  return nil
}

func (p *NtfLogin) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "NtfLogin"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NtfLogin) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *NtfLogin) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "gsn", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gsn: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Gsn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gsn (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gsn: ", p), err) }
  return err
}

func (p *NtfLogin) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "trace_id", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:trace_id: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.TraceID); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.trace_id (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:trace_id: ", p), err) }
  return err
}

func (p *NtfLogin) Equals(other *NtfLogin) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if p.Gsn != other.Gsn { return false }
  if bytes.Compare(p.TraceID, other.TraceID) != 0 { return false }
  return true
}

func (p *NtfLogin) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NtfLogin(%+v)", *p)
}

func (p *NtfLogin) Validate() error {
  return nil
}
// Attributes:
//  - Ssn
//  - Gsn
//  - TraceID
type NtfLogout struct {
  Ssn Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Gsn Gsn `thrift:"gsn,2,required" db:"gsn" json:"gsn"`
  TraceID Buffer `thrift:"trace_id,3,required" db:"trace_id" json:"trace_id"`
}

func NewNtfLogout() *NtfLogout {
  return &NtfLogout{}
}


func (p *NtfLogout) GetSsn() Ssn {
  return p.Ssn
}

func (p *NtfLogout) GetGsn() Gsn {
  return p.Gsn
}

func (p *NtfLogout) GetTraceID() Buffer {
  return p.TraceID
}
func (p *NtfLogout) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetGsn bool = false;
  var issetTraceID bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetSsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetGsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetTraceID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetSsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ssn is not set"));
  }
  if !issetGsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Gsn is not set"));
  }
  if !issetTraceID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TraceID is not set"));
  }
  return nil
}

func (p *NtfLogout)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *NtfLogout)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := Gsn(v)
  p.Gsn = temp
}
  return nil
}

func (p *NtfLogout)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := Buffer(v)
  p.TraceID = temp
}
  return nil
}

func (p *NtfLogout) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "NtfLogout"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NtfLogout) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *NtfLogout) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "gsn", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gsn: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Gsn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gsn (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gsn: ", p), err) }
  return err
}

func (p *NtfLogout) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "trace_id", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:trace_id: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.TraceID); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.trace_id (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:trace_id: ", p), err) }
  return err
}

func (p *NtfLogout) Equals(other *NtfLogout) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if p.Gsn != other.Gsn { return false }
  if bytes.Compare(p.TraceID, other.TraceID) != 0 { return false }
  return true
}

func (p *NtfLogout) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NtfLogout(%+v)", *p)
}

func (p *NtfLogout) Validate() error {
  return nil
}
// Attributes:
//  - Ssn
//  - Gsn
//  - Location
//  - TraceID
type NtfChangeLocation struct {
  Ssn Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Gsn Gsn `thrift:"gsn,2,required" db:"gsn" json:"gsn"`
  Location *UserLocation `thrift:"location,3,required" db:"location" json:"location"`
  TraceID Buffer `thrift:"trace_id,4,required" db:"trace_id" json:"trace_id"`
}

func NewNtfChangeLocation() *NtfChangeLocation {
  return &NtfChangeLocation{}
}


func (p *NtfChangeLocation) GetSsn() Ssn {
  return p.Ssn
}

func (p *NtfChangeLocation) GetGsn() Gsn {
  return p.Gsn
}
var NtfChangeLocation_Location_DEFAULT *UserLocation
func (p *NtfChangeLocation) GetLocation() *UserLocation {
  if !p.IsSetLocation() {
    return NtfChangeLocation_Location_DEFAULT
  }
return p.Location
}

func (p *NtfChangeLocation) GetTraceID() Buffer {
  return p.TraceID
}
func (p *NtfChangeLocation) IsSetLocation() bool {
  return p.Location != nil
}

func (p *NtfChangeLocation) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetGsn bool = false;
  var issetLocation bool = false;
  var issetTraceID bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetSsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetGsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetLocation = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetTraceID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetSsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ssn is not set"));
  }
  if !issetGsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Gsn is not set"));
  }
  if !issetLocation{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Location is not set"));
  }
  if !issetTraceID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TraceID is not set"));
  }
  return nil
}

func (p *NtfChangeLocation)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *NtfChangeLocation)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := Gsn(v)
  p.Gsn = temp
}
  return nil
}

func (p *NtfChangeLocation)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  p.Location = &UserLocation{}
  if err := p.Location.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Location), err)
  }
  return nil
}

func (p *NtfChangeLocation)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := Buffer(v)
  p.TraceID = temp
}
  return nil
}

func (p *NtfChangeLocation) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "NtfChangeLocation"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NtfChangeLocation) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *NtfChangeLocation) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "gsn", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gsn: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Gsn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gsn (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gsn: ", p), err) }
  return err
}

func (p *NtfChangeLocation) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "location", thrift.STRUCT, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:location: ", p), err) }
  if err := p.Location.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Location), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:location: ", p), err) }
  return err
}

func (p *NtfChangeLocation) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "trace_id", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:trace_id: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.TraceID); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.trace_id (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:trace_id: ", p), err) }
  return err
}

func (p *NtfChangeLocation) Equals(other *NtfChangeLocation) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if p.Gsn != other.Gsn { return false }
  if !p.Location.Equals(other.Location) { return false }
  if bytes.Compare(p.TraceID, other.TraceID) != 0 { return false }
  return true
}

func (p *NtfChangeLocation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NtfChangeLocation(%+v)", *p)
}

func (p *NtfChangeLocation) Validate() error {
  return nil
}
