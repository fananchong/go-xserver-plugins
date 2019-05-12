// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMD_LOGIN_ENUM int32

const (
	CMD_LOGIN_INVALID CMD_LOGIN_ENUM = 0
	CMD_LOGIN_LOGIN   CMD_LOGIN_ENUM = 1
)

var CMD_LOGIN_ENUM_name = map[int32]string{
	0: "INVALID",
	1: "LOGIN",
}
var CMD_LOGIN_ENUM_value = map[string]int32{
	"INVALID": 0,
	"LOGIN":   1,
}

func (x CMD_LOGIN_ENUM) String() string {
	return proto.EnumName(CMD_LOGIN_ENUM_name, int32(x))
}
func (CMD_LOGIN_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorLogin, []int{0, 0} }

type ENUM_LOGIN_MODE_ENUM int32

const (
	ENUM_LOGIN_MODE_DEFAULT ENUM_LOGIN_MODE_ENUM = 0
	ENUM_LOGIN_MODE_CUSTOM  ENUM_LOGIN_MODE_ENUM = 1
)

var ENUM_LOGIN_MODE_ENUM_name = map[int32]string{
	0: "DEFAULT",
	1: "CUSTOM",
}
var ENUM_LOGIN_MODE_ENUM_value = map[string]int32{
	"DEFAULT": 0,
	"CUSTOM":  1,
}

func (x ENUM_LOGIN_MODE_ENUM) String() string {
	return proto.EnumName(ENUM_LOGIN_MODE_ENUM_name, int32(x))
}
func (ENUM_LOGIN_MODE_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorLogin, []int{1, 0} }

type ENUM_LOGIN_ERROR_ENUM int32

const (
	ENUM_LOGIN_ERROR_OK           ENUM_LOGIN_ERROR_ENUM = 0
	ENUM_LOGIN_ERROR_VERIFY_FAIL  ENUM_LOGIN_ERROR_ENUM = 1
	ENUM_LOGIN_ERROR_SYSTEM_ERROR ENUM_LOGIN_ERROR_ENUM = 2
)

var ENUM_LOGIN_ERROR_ENUM_name = map[int32]string{
	0: "OK",
	1: "VERIFY_FAIL",
	2: "SYSTEM_ERROR",
}
var ENUM_LOGIN_ERROR_ENUM_value = map[string]int32{
	"OK":           0,
	"VERIFY_FAIL":  1,
	"SYSTEM_ERROR": 2,
}

func (x ENUM_LOGIN_ERROR_ENUM) String() string {
	return proto.EnumName(ENUM_LOGIN_ERROR_ENUM_name, int32(x))
}
func (ENUM_LOGIN_ERROR_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorLogin, []int{2, 0} }

type CMD_LOGIN struct {
}

func (m *CMD_LOGIN) Reset()                    { *m = CMD_LOGIN{} }
func (m *CMD_LOGIN) String() string            { return proto.CompactTextString(m) }
func (*CMD_LOGIN) ProtoMessage()               {}
func (*CMD_LOGIN) Descriptor() ([]byte, []int) { return fileDescriptorLogin, []int{0} }

// 登录 ( C -> LOGIN )
type ENUM_LOGIN_MODE struct {
}

func (m *ENUM_LOGIN_MODE) Reset()                    { *m = ENUM_LOGIN_MODE{} }
func (m *ENUM_LOGIN_MODE) String() string            { return proto.CompactTextString(m) }
func (*ENUM_LOGIN_MODE) ProtoMessage()               {}
func (*ENUM_LOGIN_MODE) Descriptor() ([]byte, []int) { return fileDescriptorLogin, []int{1} }

type ENUM_LOGIN_ERROR struct {
}

func (m *ENUM_LOGIN_ERROR) Reset()                    { *m = ENUM_LOGIN_ERROR{} }
func (m *ENUM_LOGIN_ERROR) String() string            { return proto.CompactTextString(m) }
func (*ENUM_LOGIN_ERROR) ProtoMessage()               {}
func (*ENUM_LOGIN_ERROR) Descriptor() ([]byte, []int) { return fileDescriptorLogin, []int{2} }

type MSG_LOGIN struct {
	Account  string               `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string               `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Mode     ENUM_LOGIN_MODE_ENUM `protobuf:"varint,3,opt,name=Mode,proto3,enum=protocol.ENUM_LOGIN_MODE_ENUM" json:"Mode,omitempty"`
	Userdata []byte               `protobuf:"bytes,4,opt,name=Userdata,proto3" json:"Userdata,omitempty"`
}

func (m *MSG_LOGIN) Reset()                    { *m = MSG_LOGIN{} }
func (m *MSG_LOGIN) String() string            { return proto.CompactTextString(m) }
func (*MSG_LOGIN) ProtoMessage()               {}
func (*MSG_LOGIN) Descriptor() ([]byte, []int) { return fileDescriptorLogin, []int{3} }

func (m *MSG_LOGIN) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MSG_LOGIN) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MSG_LOGIN) GetMode() ENUM_LOGIN_MODE_ENUM {
	if m != nil {
		return m.Mode
	}
	return ENUM_LOGIN_MODE_DEFAULT
}

func (m *MSG_LOGIN) GetUserdata() []byte {
	if m != nil {
		return m.Userdata
	}
	return nil
}

type MSG_LOGIN_RESULT struct {
	Err      ENUM_LOGIN_ERROR_ENUM `protobuf:"varint,1,opt,name=Err,proto3,enum=protocol.ENUM_LOGIN_ERROR_ENUM" json:"Err,omitempty"`
	Token    string                `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Address  []string              `protobuf:"bytes,3,rep,name=Address" json:"Address,omitempty"`
	Port     []int32               `protobuf:"varint,4,rep,packed,name=Port" json:"Port,omitempty"`
	NodeTyps []int32               `protobuf:"varint,5,rep,packed,name=NodeTyps" json:"NodeTyps,omitempty"`
	NodeIDs  []int32               `protobuf:"varint,6,rep,packed,name=NodeIDs" json:"NodeIDs,omitempty"`
}

func (m *MSG_LOGIN_RESULT) Reset()                    { *m = MSG_LOGIN_RESULT{} }
func (m *MSG_LOGIN_RESULT) String() string            { return proto.CompactTextString(m) }
func (*MSG_LOGIN_RESULT) ProtoMessage()               {}
func (*MSG_LOGIN_RESULT) Descriptor() ([]byte, []int) { return fileDescriptorLogin, []int{4} }

func (m *MSG_LOGIN_RESULT) GetErr() ENUM_LOGIN_ERROR_ENUM {
	if m != nil {
		return m.Err
	}
	return ENUM_LOGIN_ERROR_OK
}

func (m *MSG_LOGIN_RESULT) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MSG_LOGIN_RESULT) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MSG_LOGIN_RESULT) GetPort() []int32 {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *MSG_LOGIN_RESULT) GetNodeTyps() []int32 {
	if m != nil {
		return m.NodeTyps
	}
	return nil
}

func (m *MSG_LOGIN_RESULT) GetNodeIDs() []int32 {
	if m != nil {
		return m.NodeIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*CMD_LOGIN)(nil), "protocol.CMD_LOGIN")
	proto.RegisterType((*ENUM_LOGIN_MODE)(nil), "protocol.ENUM_LOGIN_MODE")
	proto.RegisterType((*ENUM_LOGIN_ERROR)(nil), "protocol.ENUM_LOGIN_ERROR")
	proto.RegisterType((*MSG_LOGIN)(nil), "protocol.MSG_LOGIN")
	proto.RegisterType((*MSG_LOGIN_RESULT)(nil), "protocol.MSG_LOGIN_RESULT")
	proto.RegisterEnum("protocol.CMD_LOGIN_ENUM", CMD_LOGIN_ENUM_name, CMD_LOGIN_ENUM_value)
	proto.RegisterEnum("protocol.ENUM_LOGIN_MODE_ENUM", ENUM_LOGIN_MODE_ENUM_name, ENUM_LOGIN_MODE_ENUM_value)
	proto.RegisterEnum("protocol.ENUM_LOGIN_ERROR_ENUM", ENUM_LOGIN_ERROR_ENUM_name, ENUM_LOGIN_ERROR_ENUM_value)
}
func (m *CMD_LOGIN) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CMD_LOGIN) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ENUM_LOGIN_MODE) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ENUM_LOGIN_MODE) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ENUM_LOGIN_ERROR) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ENUM_LOGIN_ERROR) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MSG_LOGIN) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_LOGIN) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if m.Mode != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogin(dAtA, i, uint64(m.Mode))
	}
	if len(m.Userdata) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Userdata)))
		i += copy(dAtA[i:], m.Userdata)
	}
	return i, nil
}

func (m *MSG_LOGIN_RESULT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_LOGIN_RESULT) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogin(dAtA, i, uint64(m.Err))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Port) > 0 {
		dAtA2 := make([]byte, len(m.Port)*10)
		var j1 int
		for _, num1 := range m.Port {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogin(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.NodeTyps) > 0 {
		dAtA4 := make([]byte, len(m.NodeTyps)*10)
		var j3 int
		for _, num1 := range m.NodeTyps {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLogin(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if len(m.NodeIDs) > 0 {
		dAtA6 := make([]byte, len(m.NodeIDs)*10)
		var j5 int
		for _, num1 := range m.NodeIDs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintLogin(dAtA, i, uint64(j5))
		i += copy(dAtA[i:], dAtA6[:j5])
	}
	return i, nil
}

func encodeFixed64Login(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Login(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CMD_LOGIN) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ENUM_LOGIN_MODE) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ENUM_LOGIN_ERROR) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MSG_LOGIN) Size() (n int) {
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.Mode != 0 {
		n += 1 + sovLogin(uint64(m.Mode))
	}
	l = len(m.Userdata)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	return n
}

func (m *MSG_LOGIN_RESULT) Size() (n int) {
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovLogin(uint64(m.Err))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			l = len(s)
			n += 1 + l + sovLogin(uint64(l))
		}
	}
	if len(m.Port) > 0 {
		l = 0
		for _, e := range m.Port {
			l += sovLogin(uint64(e))
		}
		n += 1 + sovLogin(uint64(l)) + l
	}
	if len(m.NodeTyps) > 0 {
		l = 0
		for _, e := range m.NodeTyps {
			l += sovLogin(uint64(e))
		}
		n += 1 + sovLogin(uint64(l)) + l
	}
	if len(m.NodeIDs) > 0 {
		l = 0
		for _, e := range m.NodeIDs {
			l += sovLogin(uint64(e))
		}
		n += 1 + sovLogin(uint64(l)) + l
	}
	return n
}

func sovLogin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CMD_LOGIN) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CMD_LOGIN: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CMD_LOGIN: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ENUM_LOGIN_MODE) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ENUM_LOGIN_MODE: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ENUM_LOGIN_MODE: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ENUM_LOGIN_ERROR) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ENUM_LOGIN_ERROR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ENUM_LOGIN_ERROR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MSG_LOGIN) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MSG_LOGIN: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_LOGIN: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= (ENUM_LOGIN_MODE_ENUM(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userdata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userdata = append(m.Userdata[:0], dAtA[iNdEx:postIndex]...)
			if m.Userdata == nil {
				m.Userdata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MSG_LOGIN_RESULT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MSG_LOGIN_RESULT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_LOGIN_RESULT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= (ENUM_LOGIN_ERROR_ENUM(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Port = append(m.Port, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLogin
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogin
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Port = append(m.Port, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
		case 5:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.NodeTyps = append(m.NodeTyps, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLogin
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogin
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.NodeTyps = append(m.NodeTyps, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeTyps", wireType)
			}
		case 6:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.NodeIDs = append(m.NodeIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLogin
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogin
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.NodeIDs = append(m.NodeIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeIDs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLogin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLogin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLogin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("login.proto", fileDescriptorLogin) }

var fileDescriptorLogin = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xdf, 0x8a, 0x9b, 0x40,
	0x18, 0xc5, 0x9d, 0xf8, 0x67, 0xd7, 0x2f, 0xcb, 0xee, 0x30, 0xf4, 0x62, 0xe8, 0x85, 0x2b, 0x73,
	0x25, 0x14, 0x02, 0x6b, 0x9f, 0xc0, 0xae, 0x93, 0x45, 0xea, 0x9f, 0x30, 0x6a, 0x20, 0x57, 0x92,
	0x46, 0x29, 0xa5, 0x21, 0x13, 0xd4, 0x52, 0xfa, 0x16, 0xed, 0x13, 0xf5, 0xb6, 0x97, 0x7d, 0x84,
	0x92, 0xbe, 0x48, 0x71, 0x4c, 0x6c, 0x29, 0x7b, 0xe5, 0xfc, 0x3c, 0xdf, 0xf9, 0xe6, 0x9c, 0x81,
	0xf9, 0x5e, 0xbe, 0xff, 0x70, 0x58, 0x1c, 0x5b, 0xd9, 0x4b, 0x72, 0xad, 0x3e, 0x3b, 0xb9, 0x67,
	0xaf, 0xc0, 0x7e, 0x4c, 0xc2, 0x2a, 0xce, 0x9e, 0xa2, 0x94, 0x39, 0x60, 0xf0, 0xb4, 0x4c, 0xc8,
	0x1c, 0xae, 0xa2, 0x74, 0x1d, 0xc4, 0x51, 0x88, 0x35, 0x62, 0x83, 0xa9, 0x54, 0x8c, 0x98, 0x0f,
	0x77, 0x83, 0x3e, 0x4e, 0x57, 0x49, 0x16, 0x72, 0x76, 0xff, 0xd7, 0x12, 0xf2, 0x65, 0x50, 0xc6,
	0x05, 0xd6, 0x08, 0x80, 0xf5, 0x58, 0xe6, 0x45, 0x96, 0x60, 0xc4, 0x38, 0xe0, 0x7f, 0x3c, 0x5c,
	0x88, 0x4c, 0xb0, 0x87, 0xb3, 0xc9, 0x82, 0x59, 0xf6, 0x16, 0x6b, 0xe4, 0x0e, 0xe6, 0x6b, 0x2e,
	0xa2, 0xe5, 0xa6, 0x5a, 0x06, 0x51, 0x8c, 0x11, 0xc1, 0x70, 0x93, 0x6f, 0xf2, 0x82, 0x27, 0xa3,
	0x01, 0xcf, 0xd8, 0x37, 0x04, 0x76, 0x92, 0x3f, 0x8d, 0x6b, 0x08, 0x85, 0xab, 0x60, 0xb7, 0x93,
	0x9f, 0x0e, 0x3d, 0x45, 0x2e, 0xf2, 0x6c, 0x71, 0x41, 0xf2, 0x12, 0xae, 0x57, 0xdb, 0xae, 0xfb,
	0x2c, 0xdb, 0x9a, 0xce, 0x94, 0x34, 0x31, 0xf1, 0xc1, 0x48, 0x64, 0xdd, 0x50, 0xdd, 0x45, 0xde,
	0xad, 0xef, 0x2c, 0x2e, 0x8f, 0xb0, 0xf8, 0xaf, 0x94, 0x62, 0xa1, 0x66, 0x87, 0x7d, 0x65, 0xd7,
	0xb4, 0xf5, 0xb6, 0xdf, 0x52, 0xc3, 0x45, 0xde, 0x8d, 0x98, 0x98, 0x7d, 0x47, 0x80, 0xa7, 0x4c,
	0x95, 0xe0, 0x79, 0x19, 0x17, 0xe4, 0x01, 0x74, 0xde, 0xb6, 0x2a, 0xd6, 0xad, 0x7f, 0xff, 0xec,
	0x1d, 0xaa, 0xd3, 0x78, 0xc9, 0x30, 0x4b, 0x5e, 0x80, 0x59, 0xc8, 0x8f, 0xcd, 0xe1, 0x1c, 0x78,
	0x04, 0xd5, 0xb1, 0xae, 0xdb, 0xa6, 0xeb, 0xa8, 0xee, 0xea, 0xaa, 0xe3, 0x88, 0x84, 0x80, 0xb1,
	0x92, 0x6d, 0x4f, 0x0d, 0x57, 0xf7, 0x4c, 0xa1, 0xce, 0x43, 0xce, 0x54, 0xd6, 0x4d, 0xf1, 0xe5,
	0xd8, 0x51, 0x53, 0xfd, 0x9f, 0x78, 0xd8, 0x34, 0x9c, 0xa3, 0xb0, 0xa3, 0x96, 0x92, 0x2e, 0xf8,
	0x06, 0xff, 0x38, 0x39, 0xe8, 0xe7, 0xc9, 0x41, 0xbf, 0x4e, 0x0e, 0xfa, 0xfa, 0xdb, 0xd1, 0xde,
	0x59, 0x2a, 0xf0, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x3d, 0x1d, 0xba, 0x2f, 0x02,
	0x00, 0x00,
}
