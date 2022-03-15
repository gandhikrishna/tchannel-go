// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"fmt"
	"github.com/uber/tchannel-go/thirdparty/github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Healthy
//  - Msg
type HealthCheckRes struct {
	Healthy bool   `thrift:"healthy,1" db:"healthy" json:"healthy"`
	Msg     string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewHealthCheckRes() *HealthCheckRes {
	return &HealthCheckRes{}
}

func (p *HealthCheckRes) GetHealthy() bool {
	return p.Healthy
}

func (p *HealthCheckRes) GetMsg() string {
	return p.Msg
}
func (p *HealthCheckRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *HealthCheckRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Healthy = v
	}
	return nil
}

func (p *HealthCheckRes) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *HealthCheckRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HealthCheckRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HealthCheckRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("healthy", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:healthy: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Healthy)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.healthy (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:healthy: ", p), err)
	}
	return err
}

func (p *HealthCheckRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err)
	}
	if err := oprot.WriteString(string(p.Msg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err)
	}
	return err
}

func (p *HealthCheckRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HealthCheckRes(%+v)", *p)
}
