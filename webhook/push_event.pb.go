// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_event.proto

package vcs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PushCommit struct {
	Sha       string        `protobuf:"bytes,1,opt,name=sha" json:"sha,omitempty"`
	Id        string        `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	TreeId    string        `protobuf:"bytes,3,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	Distinct  bool          `protobuf:"varint,4,opt,name=distinct" json:"distinct,omitempty"`
	Message   string        `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Timestamp string        `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Url       string        `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,8,opt,name=author" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,9,opt,name=committer" json:"committer,omitempty"`
	Added     []string      `protobuf:"bytes,10,rep,name=added" json:"added,omitempty"`
	Removed   []string      `protobuf:"bytes,11,rep,name=removed" json:"removed,omitempty"`
	Modified  []string      `protobuf:"bytes,12,rep,name=modified" json:"modified,omitempty"`
}

func (m *PushCommit) Reset()                    { *m = PushCommit{} }
func (m *PushCommit) String() string            { return proto.CompactTextString(m) }
func (*PushCommit) ProtoMessage()               {}
func (*PushCommit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PushCommit) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *PushCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PushCommit) GetTreeId() string {
	if m != nil {
		return m.TreeId
	}
	return ""
}

func (m *PushCommit) GetDistinct() bool {
	if m != nil {
		return m.Distinct
	}
	return false
}

func (m *PushCommit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PushCommit) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PushCommit) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PushCommit) GetAuthor() *CommitAuthor {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *PushCommit) GetCommitter() *CommitAuthor {
	if m != nil {
		return m.Committer
	}
	return nil
}

func (m *PushCommit) GetAdded() []string {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *PushCommit) GetRemoved() []string {
	if m != nil {
		return m.Removed
	}
	return nil
}

func (m *PushCommit) GetModified() []string {
	if m != nil {
		return m.Modified
	}
	return nil
}

type VcsPushEvent struct {
	Ref          string        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Created      bool          `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Deleted      bool          `protobuf:"varint,5,opt,name=deleted" json:"deleted,omitempty"`
	Forced       bool          `protobuf:"varint,6,opt,name=forced" json:"forced,omitempty"`
	Commits      []*PushCommit `protobuf:"bytes,9,rep,name=commits" json:"commits,omitempty"`
	HeadCommit   *PushCommit   `protobuf:"bytes,10,opt,name=head_commit,json=headCommit" json:"head_commit,omitempty"`
	Repository   *Repository   `protobuf:"bytes,11,opt,name=repository" json:"repository,omitempty"`
	Pusher       *CommitAuthor `protobuf:"bytes,12,opt,name=pusher" json:"pusher,omitempty"`
	Sender       *User         `protobuf:"bytes,13,opt,name=sender" json:"sender,omitempty"`
	Organization *User         `protobuf:"bytes,15,opt,name=organization" json:"organization,omitempty"`
}

func (m *VcsPushEvent) Reset()                    { *m = VcsPushEvent{} }
func (m *VcsPushEvent) String() string            { return proto.CompactTextString(m) }
func (*VcsPushEvent) ProtoMessage()               {}
func (*VcsPushEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *VcsPushEvent) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *VcsPushEvent) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *VcsPushEvent) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *VcsPushEvent) GetForced() bool {
	if m != nil {
		return m.Forced
	}
	return false
}

func (m *VcsPushEvent) GetCommits() []*PushCommit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *VcsPushEvent) GetHeadCommit() *PushCommit {
	if m != nil {
		return m.HeadCommit
	}
	return nil
}

func (m *VcsPushEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *VcsPushEvent) GetPusher() *CommitAuthor {
	if m != nil {
		return m.Pusher
	}
	return nil
}

func (m *VcsPushEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *VcsPushEvent) GetOrganization() *User {
	if m != nil {
		return m.Organization
	}
	return nil
}

func init() {
	proto.RegisterType((*PushCommit)(nil), "vcs.PushCommit")
	proto.RegisterType((*VcsPushEvent)(nil), "vcs.VcsPushEvent")
}

func init() { proto.RegisterFile("push_event.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x13, 0x9a, 0x36, 0x93, 0xc2, 0x16, 0x83, 0x60, 0x54, 0x71, 0x08, 0x7b, 0xca, 0x1e,
	0xe8, 0xa2, 0xe5, 0x09, 0x10, 0xe2, 0xc0, 0x0d, 0x59, 0x82, 0x6b, 0x65, 0xe2, 0xe9, 0xc6, 0x52,
	0x13, 0x57, 0xb6, 0x53, 0x09, 0xde, 0x93, 0x67, 0xe1, 0x8a, 0x6c, 0x27, 0xcd, 0xc2, 0xaa, 0x37,
	0x7f, 0x3f, 0x63, 0xcd, 0x7c, 0x33, 0xb0, 0x3e, 0xf6, 0xb6, 0xd9, 0xd1, 0x89, 0x3a, 0xb7, 0x3d,
	0x1a, 0xed, 0x34, 0x4b, 0x4f, 0xb5, 0xdd, 0xbc, 0xa8, 0x75, 0xdb, 0x2a, 0xb7, 0x13, 0xbd, 0x6b,
	0xb4, 0x89, 0xca, 0x06, 0x7a, 0x4b, 0xe3, 0x7b, 0x6d, 0xe8, 0xa8, 0xad, 0x72, 0xda, 0xfc, 0x8c,
	0xcc, 0xf5, 0xef, 0x04, 0xe0, 0x6b, 0x6f, 0x9b, 0x4f, 0xa1, 0x92, 0xad, 0x21, 0xb5, 0x8d, 0xc0,
	0x59, 0x39, 0xab, 0x72, 0xee, 0x9f, 0xec, 0x19, 0x24, 0x4a, 0x62, 0x12, 0x88, 0x44, 0x49, 0xf6,
	0x1a, 0x16, 0xce, 0x10, 0xed, 0x94, 0xc4, 0x34, 0x90, 0x99, 0x87, 0x5f, 0x24, 0xdb, 0xc0, 0x52,
	0x2a, 0xeb, 0x54, 0x57, 0x3b, 0x7c, 0x52, 0xce, 0xaa, 0x25, 0x3f, 0x63, 0x86, 0xb0, 0x68, 0xc9,
	0x5a, 0x71, 0x4f, 0x38, 0x0f, 0x45, 0x23, 0x64, 0x6f, 0x20, 0x77, 0xaa, 0x25, 0xeb, 0x44, 0x7b,
	0xc4, 0x2c, 0x68, 0x13, 0xe1, 0xdb, 0xe9, 0xcd, 0x01, 0x17, 0xb1, 0x9d, 0xde, 0x1c, 0xd8, 0x0d,
	0x64, 0x71, 0x3a, 0x5c, 0x96, 0xb3, 0xaa, 0xb8, 0x7b, 0xbe, 0x3d, 0xd5, 0x76, 0x1b, 0xbb, 0xff,
	0x18, 0x04, 0x3e, 0x18, 0xd8, 0x2d, 0xe4, 0x31, 0x0f, 0x47, 0x06, 0xf3, 0x4b, 0xee, 0xc9, 0xc3,
	0x5e, 0xc2, 0x5c, 0x48, 0x49, 0x12, 0xa1, 0x4c, 0xab, 0x9c, 0x47, 0xe0, 0x7b, 0x37, 0xd4, 0xea,
	0x13, 0x49, 0x2c, 0x02, 0x3f, 0x42, 0x3f, 0x71, 0xab, 0xa5, 0xda, 0x2b, 0x92, 0xb8, 0x0a, 0xd2,
	0x19, 0x5f, 0xff, 0x49, 0x60, 0xf5, 0xbd, 0xb6, 0x3e, 0xda, 0xcf, 0x7e, 0x4d, 0x7e, 0x14, 0x43,
	0xfb, 0x31, 0x59, 0x43, 0x7b, 0xff, 0x71, 0x6d, 0x48, 0x38, 0x92, 0x43, 0x5e, 0x23, 0xf4, 0x8a,
	0xa4, 0x03, 0x79, 0x65, 0x1e, 0x95, 0x01, 0xb2, 0x57, 0x90, 0xed, 0xb5, 0xa9, 0x49, 0x86, 0xac,
	0x96, 0x7c, 0x40, 0xec, 0x06, 0x16, 0x71, 0x0e, 0x8b, 0x79, 0x99, 0x56, 0xc5, 0xdd, 0x55, 0x98,
	0x74, 0xda, 0x2c, 0x1f, 0x75, 0xf6, 0x1e, 0x8a, 0x86, 0x84, 0xdc, 0x45, 0x8c, 0x10, 0x82, 0x79,
	0x64, 0x07, 0xef, 0x19, 0x8e, 0xe2, 0x16, 0x60, 0xba, 0x1b, 0x2c, 0x1e, 0x14, 0xf0, 0x33, 0xcd,
	0x1f, 0x58, 0xfc, 0x92, 0xfc, 0x81, 0x92, 0xc1, 0xd5, 0xc5, 0x25, 0x45, 0x03, 0x7b, 0x0b, 0x99,
	0xa5, 0x4e, 0x92, 0xc1, 0xa7, 0xc1, 0x9a, 0x07, 0xeb, 0x37, 0x4b, 0x86, 0x0f, 0x02, 0x7b, 0x07,
	0x2b, 0x6d, 0xee, 0x45, 0xa7, 0x7e, 0x09, 0xa7, 0x74, 0x87, 0x57, 0xff, 0x1b, 0xff, 0x91, 0x7f,
	0x64, 0xe1, 0xb0, 0x3f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x78, 0xb0, 0x95, 0xd3, 0x24, 0x03,
	0x00, 0x00,
}
