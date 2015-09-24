// Code generated by protoc-gen-go.
// source: bacs/problem/problem.proto
// DO NOT EDIT!

/*
Package bacs_problem is a generated protocol buffer package.

It is generated from these files:
	bacs/problem/problem.proto
	bacs/problem/result.proto

It has these top-level messages:
	Problem
	System
	Revision
	Info
	Statement
	Profile
	Utility
*/
package bacs_problem

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Problem struct {
	System    *System              `protobuf:"bytes,1,opt,name=system" json:"system,omitempty"`
	Info      *Info                `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Statement *Statement           `protobuf:"bytes,3,opt,name=statement" json:"statement,omitempty"`
	Profile   []*Profile           `protobuf:"bytes,4,rep,name=profile" json:"profile,omitempty"`
	Utility   map[string]*Utility  `protobuf:"bytes,5,rep,name=utility" json:"utility,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Extension *google_protobuf.Any `protobuf:"bytes,1000,opt,name=extension" json:"extension,omitempty"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}

func (m *Problem) GetSystem() *System {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *Problem) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Problem) GetStatement() *Statement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *Problem) GetProfile() []*Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *Problem) GetUtility() map[string]*Utility {
	if m != nil {
		return m.Utility
	}
	return nil
}

func (m *Problem) GetExtension() *google_protobuf.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

type System struct {
	ProblemType string    `protobuf:"bytes,1,opt,name=problem_type" json:"problem_type,omitempty"`
	Package     string    `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	Revision    *Revision `protobuf:"bytes,3,opt,name=revision" json:"revision,omitempty"`
}

func (m *System) Reset()         { *m = System{} }
func (m *System) String() string { return proto.CompactTextString(m) }
func (*System) ProtoMessage()    {}

func (m *System) GetRevision() *Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

type Revision struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Revision) Reset()         { *m = Revision{} }
func (m *Revision) String() string { return proto.CompactTextString(m) }
func (*Revision) ProtoMessage()    {}

type Info struct {
	Name       []*Info_Name `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
	Author     []string     `protobuf:"bytes,2,rep,name=author" json:"author,omitempty"`
	Source     string       `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Maintainer []string     `protobuf:"bytes,4,rep,name=maintainer" json:"maintainer,omitempty"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}

func (m *Info) GetName() []*Info_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

type Info_Name struct {
	Language string `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Info_Name) Reset()         { *m = Info_Name{} }
func (m *Info_Name) String() string { return proto.CompactTextString(m) }
func (*Info_Name) ProtoMessage()    {}

type Statement struct {
	Version []*Statement_Version `protobuf:"bytes,1,rep,name=version" json:"version,omitempty"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}

func (m *Statement) GetVersion() []*Statement_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Statement_Version struct {
	Language string `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Format   string `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	Package  string `protobuf:"bytes,3,opt,name=package" json:"package,omitempty"`
}

func (m *Statement_Version) Reset()         { *m = Statement_Version{} }
func (m *Statement_Version) String() string { return proto.CompactTextString(m) }
func (*Statement_Version) ProtoMessage()    {}

type Statement_Version_Manifest struct {
	Version  *Statement_Version               `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Revision *Revision                        `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Data     *Statement_Version_Manifest_Data `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *Statement_Version_Manifest) Reset()         { *m = Statement_Version_Manifest{} }
func (m *Statement_Version_Manifest) String() string { return proto.CompactTextString(m) }
func (*Statement_Version_Manifest) ProtoMessage()    {}

func (m *Statement_Version_Manifest) GetVersion() *Statement_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Statement_Version_Manifest) GetRevision() *Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Statement_Version_Manifest) GetData() *Statement_Version_Manifest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Statement_Version_Manifest_Data struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *Statement_Version_Manifest_Data) Reset()         { *m = Statement_Version_Manifest_Data{} }
func (m *Statement_Version_Manifest_Data) String() string { return proto.CompactTextString(m) }
func (*Statement_Version_Manifest_Data) ProtoMessage()    {}

type Profile struct {
	Description string               `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Extension   *google_protobuf.Any `protobuf:"bytes,1000,opt,name=extension" json:"extension,omitempty"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}

func (m *Profile) GetExtension() *google_protobuf.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

type Utility struct {
	BuilderName string               `protobuf:"bytes,1,opt,name=builder_name" json:"builder_name,omitempty"`
	Nested      []*Utility           `protobuf:"bytes,100,rep,name=nested" json:"nested,omitempty"`
	Extension   *google_protobuf.Any `protobuf:"bytes,1000,opt,name=extension" json:"extension,omitempty"`
}

func (m *Utility) Reset()         { *m = Utility{} }
func (m *Utility) String() string { return proto.CompactTextString(m) }
func (*Utility) ProtoMessage()    {}

func (m *Utility) GetNested() []*Utility {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Utility) GetExtension() *google_protobuf.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}
