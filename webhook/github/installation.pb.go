// Code generated by protoc-gen-go. DO NOT EDIT.
// source: installation.proto

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Installation struct {
	Id              int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Account         *User  `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	AccessTokensUrl string `protobuf:"bytes,3,opt,name=access_tokens_url,json=accessTokensUrl" json:"access_tokens_url,omitempty"`
	RepositoriesUrl string `protobuf:"bytes,4,opt,name=repositories_url,json=repositoriesUrl" json:"repositories_url,omitempty"`
	HtmlUrl         string `protobuf:"bytes,5,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
}

func (m *Installation) Reset()                    { *m = Installation{} }
func (m *Installation) String() string            { return proto.CompactTextString(m) }
func (*Installation) ProtoMessage()               {}
func (*Installation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Installation) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Installation) GetAccount() *User {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Installation) GetAccessTokensUrl() string {
	if m != nil {
		return m.AccessTokensUrl
	}
	return ""
}

func (m *Installation) GetRepositoriesUrl() string {
	if m != nil {
		return m.RepositoriesUrl
	}
	return ""
}

func (m *Installation) GetHtmlUrl() string {
	if m != nil {
		return m.HtmlUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Installation)(nil), "github.Installation")
}

func init() { proto.RegisterFile("installation.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xcf, 0x6a, 0x84, 0x30,
	0x10, 0xc7, 0x71, 0x62, 0xab, 0xb6, 0x53, 0xe9, 0x9f, 0x9c, 0x6c, 0x4f, 0xd2, 0x43, 0xb1, 0x3d,
	0x78, 0x68, 0x9f, 0xa2, 0x57, 0x59, 0xcf, 0x12, 0x63, 0x58, 0xc3, 0x66, 0x13, 0x99, 0x4c, 0xde,
	0x6d, 0x1f, 0x6f, 0x31, 0x61, 0xc1, 0x63, 0xbe, 0xbf, 0x0f, 0x81, 0x01, 0xae, 0xad, 0x27, 0x61,
	0x8c, 0x20, 0xed, 0x6c, 0xb7, 0xa2, 0x23, 0xc7, 0x8b, 0xa3, 0xa6, 0x25, 0x4c, 0x1f, 0x10, 0xbc,
	0xc2, 0xd4, 0x3e, 0x2f, 0x0c, 0xaa, 0xff, 0x1d, 0xe5, 0xcf, 0x90, 0xe9, 0xb9, 0x66, 0x0d, 0x6b,
	0xf3, 0x3e, 0xd3, 0x33, 0xff, 0x82, 0x52, 0x48, 0xe9, 0x82, 0xa5, 0x3a, 0x6b, 0x58, 0xfb, 0xf4,
	0x5b, 0x75, 0xe9, 0x9b, 0x6e, 0xf0, 0x0a, 0xfb, 0xdb, 0xc8, 0x7f, 0xe0, 0x4d, 0x48, 0xa9, 0xbc,
	0x1f, 0xc9, 0x9d, 0x94, 0xf5, 0x63, 0x40, 0x53, 0xdf, 0x35, 0xac, 0x7d, 0xec, 0x5f, 0xd2, 0x70,
	0x88, 0x7d, 0x40, 0xc3, 0xbf, 0xe1, 0x15, 0xd5, 0xea, 0xbc, 0x26, 0x87, 0x5a, 0x25, 0x7a, 0x9f,
	0xe8, 0xbe, 0x6f, 0xf4, 0x1d, 0x1e, 0x16, 0x3a, 0x9b, 0x48, 0xf2, 0x48, 0xca, 0xed, 0x3d, 0xa0,
	0x99, 0x8a, 0x78, 0xc1, 0xdf, 0x35, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x16, 0x8d, 0xff, 0xeb, 0x00,
	0x00, 0x00,
}