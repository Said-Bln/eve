// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fw.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ACEMatch struct {
	// FIXME: We should convert this to enum
	Type  string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ACEMatch) Reset()                    { *m = ACEMatch{} }
func (m *ACEMatch) String() string            { return proto.CompactTextString(m) }
func (*ACEMatch) ProtoMessage()               {}
func (*ACEMatch) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ACEMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ACEMatch) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ACEAction struct {
	Drop bool `protobuf:"varint,1,opt,name=drop" json:"drop,omitempty"`
	// limit action, and its associated parameter
	Limit      bool   `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Limitrate  uint32 `protobuf:"varint,3,opt,name=limitrate" json:"limitrate,omitempty"`
	Limitunit  string `protobuf:"bytes,4,opt,name=limitunit" json:"limitunit,omitempty"`
	Limitburst uint32 `protobuf:"varint,5,opt,name=limitburst" json:"limitburst,omitempty"`
	// port map action, and its assoicated paramtere
	Portmap bool   `protobuf:"varint,6,opt,name=portmap" json:"portmap,omitempty"`
	AppPort uint32 `protobuf:"varint,7,opt,name=appPort" json:"appPort,omitempty"`
}

func (m *ACEAction) Reset()                    { *m = ACEAction{} }
func (m *ACEAction) String() string            { return proto.CompactTextString(m) }
func (*ACEAction) ProtoMessage()               {}
func (*ACEAction) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ACEAction) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *ACEAction) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *ACEAction) GetLimitrate() uint32 {
	if m != nil {
		return m.Limitrate
	}
	return 0
}

func (m *ACEAction) GetLimitunit() string {
	if m != nil {
		return m.Limitunit
	}
	return ""
}

func (m *ACEAction) GetLimitburst() uint32 {
	if m != nil {
		return m.Limitburst
	}
	return 0
}

func (m *ACEAction) GetPortmap() bool {
	if m != nil {
		return m.Portmap
	}
	return false
}

func (m *ACEAction) GetAppPort() uint32 {
	if m != nil {
		return m.AppPort
	}
	return 0
}

type ACE struct {
	Matches []*ACEMatch `protobuf:"bytes,1,rep,name=matches" json:"matches,omitempty"`
	// Expect only single action...repeated here is
	// for future work.
	Actions []*ACEAction `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *ACE) Reset()                    { *m = ACE{} }
func (m *ACE) String() string            { return proto.CompactTextString(m) }
func (*ACE) ProtoMessage()               {}
func (*ACE) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ACE) GetMatches() []*ACEMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *ACE) GetActions() []*ACEAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ACEMatch)(nil), "ACEMatch")
	proto.RegisterType((*ACEAction)(nil), "ACEAction")
	proto.RegisterType((*ACE)(nil), "ACE")
}

func init() { proto.RegisterFile("fw.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbb, 0x6a, 0xf3, 0x30,
	0x14, 0xc7, 0x71, 0x6e, 0xb6, 0xcf, 0xc7, 0xb7, 0x88, 0x0e, 0x1a, 0x7a, 0x09, 0x69, 0x87, 0x4c,
	0x32, 0xb4, 0x7d, 0x80, 0xba, 0xc6, 0x63, 0x21, 0x68, 0xec, 0x26, 0xcb, 0x4a, 0x22, 0xb0, 0x2d,
	0x21, 0x4b, 0x29, 0xcd, 0xfb, 0xf5, 0xbd, 0x8a, 0x8f, 0x71, 0xd2, 0xed, 0xfc, 0x2f, 0x3f, 0x84,
	0xfe, 0x90, 0xec, 0xbf, 0x98, 0x75, 0xc6, 0x9b, 0xcd, 0x2b, 0x24, 0x79, 0x51, 0x7e, 0x08, 0x2f,
	0x8f, 0x84, 0xc0, 0xc2, 0x7f, 0x5b, 0x45, 0xa3, 0x75, 0xb4, 0x4d, 0x39, 0xde, 0xe4, 0x06, 0x96,
	0x27, 0xd1, 0x04, 0x45, 0x67, 0x68, 0x8e, 0x62, 0xf3, 0x13, 0x41, 0x9a, 0x17, 0x65, 0x2e, 0xbd,
	0x36, 0xdd, 0xc0, 0xd5, 0xce, 0x58, 0xe4, 0x12, 0x8e, 0xf7, 0xc0, 0x35, 0xba, 0xd5, 0x1e, 0xb9,
	0x84, 0x8f, 0x82, 0xdc, 0x42, 0x8a, 0x87, 0x13, 0x5e, 0xd1, 0xf9, 0x3a, 0xda, 0xfe, 0xe7, 0x57,
	0xe3, 0x92, 0x86, 0x4e, 0x7b, 0xba, 0xc0, 0xf7, 0xae, 0x06, 0xb9, 0x07, 0x40, 0x51, 0x05, 0xd7,
	0x7b, 0xba, 0x44, 0xf8, 0x8f, 0x43, 0x28, 0xc4, 0xd6, 0x38, 0xdf, 0x0a, 0x4b, 0x57, 0xf8, 0xe6,
	0x24, 0x87, 0x44, 0x58, 0xbb, 0x33, 0xce, 0xd3, 0x18, 0xb1, 0x49, 0x6e, 0x76, 0x30, 0xcf, 0x8b,
	0x92, 0x3c, 0x42, 0xdc, 0x0e, 0x0b, 0xa8, 0x9e, 0x46, 0xeb, 0xf9, 0xf6, 0xdf, 0x73, 0xca, 0xa6,
	0x51, 0xf8, 0x94, 0x90, 0x27, 0x88, 0x05, 0xfe, 0xb7, 0xa7, 0x33, 0x2c, 0x01, 0xbb, 0x4c, 0xc0,
	0xa7, 0xe8, 0xfd, 0x0d, 0x1e, 0xa4, 0x69, 0xd9, 0x59, 0xd5, 0xaa, 0x16, 0x4c, 0x36, 0x26, 0xd4,
	0x2c, 0xf4, 0xca, 0x9d, 0xb4, 0x54, 0xe3, 0xe4, 0x9f, 0x77, 0x07, 0xed, 0x8f, 0xa1, 0x62, 0xd2,
	0xb4, 0xd9, 0xd8, 0xcb, 0x84, 0xd5, 0xd9, 0x59, 0x9a, 0x6e, 0xaf, 0x0f, 0xd5, 0x0a, 0x5b, 0x2f,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x29, 0x45, 0x9f, 0xa4, 0x01, 0x00, 0x00,
}
