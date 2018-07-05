// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/remoteworkers/v1test2/command.proto

package remoteworkers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf5 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes a shell-style task to execute.
type CommandTask struct {
	// The inputs to the task.
	Inputs *CommandTask_Inputs `protobuf:"bytes,1,opt,name=inputs" json:"inputs,omitempty"`
	// The expected outputs from the task.
	ExpectedOutputs *CommandTask_Outputs `protobuf:"bytes,4,opt,name=expected_outputs,json=expectedOutputs" json:"expected_outputs,omitempty"`
	// The timeouts of this task.
	Timeouts *CommandTask_Timeouts `protobuf:"bytes,5,opt,name=timeouts" json:"timeouts,omitempty"`
}

func (m *CommandTask) Reset()                    { *m = CommandTask{} }
func (m *CommandTask) String() string            { return proto.CompactTextString(m) }
func (*CommandTask) ProtoMessage()               {}
func (*CommandTask) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommandTask) GetInputs() *CommandTask_Inputs {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CommandTask) GetExpectedOutputs() *CommandTask_Outputs {
	if m != nil {
		return m.ExpectedOutputs
	}
	return nil
}

func (m *CommandTask) GetTimeouts() *CommandTask_Timeouts {
	if m != nil {
		return m.Timeouts
	}
	return nil
}

// Describes the inputs to a shell-style task.
type CommandTask_Inputs struct {
	// The command itself to run (e.g., argv)
	Arguments []string `protobuf:"bytes,1,rep,name=arguments" json:"arguments,omitempty"`
	// The input filesystem to be set up prior to the task beginning. The
	// contents should be a repeated set of FileMetadata messages though other
	// formats are allowed if better for the implementation (eg, a LUCI-style
	// .isolated file).
	//
	// This field is repeated since implementations might want to cache the
	// metadata, in which case it may be useful to break up portions of the
	// filesystem that change frequently (eg, specific input files) from those
	// that don't (eg, standard header files).
	Files []*Digest `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
	// All environment variables required by the task.
	EnvironmentVariables []*CommandTask_Inputs_EnvironmentVariable `protobuf:"bytes,3,rep,name=environment_variables,json=environmentVariables" json:"environment_variables,omitempty"`
}

func (m *CommandTask_Inputs) Reset()                    { *m = CommandTask_Inputs{} }
func (m *CommandTask_Inputs) String() string            { return proto.CompactTextString(m) }
func (*CommandTask_Inputs) ProtoMessage()               {}
func (*CommandTask_Inputs) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *CommandTask_Inputs) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *CommandTask_Inputs) GetFiles() []*Digest {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CommandTask_Inputs) GetEnvironmentVariables() []*CommandTask_Inputs_EnvironmentVariable {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

// An environment variable required by this task.
type CommandTask_Inputs_EnvironmentVariable struct {
	// The envvar name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The envvar value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *CommandTask_Inputs_EnvironmentVariable) Reset() {
	*m = CommandTask_Inputs_EnvironmentVariable{}
}
func (m *CommandTask_Inputs_EnvironmentVariable) String() string { return proto.CompactTextString(m) }
func (*CommandTask_Inputs_EnvironmentVariable) ProtoMessage()    {}
func (*CommandTask_Inputs_EnvironmentVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0, 0}
}

func (m *CommandTask_Inputs_EnvironmentVariable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommandTask_Inputs_EnvironmentVariable) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Describes the expected outputs of the command.
type CommandTask_Outputs struct {
	// A list of expected files, relative to the execution root.
	Files []string `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	// A list of expected directories, relative to the execution root.
	Directories []string `protobuf:"bytes,2,rep,name=directories" json:"directories,omitempty"`
}

func (m *CommandTask_Outputs) Reset()                    { *m = CommandTask_Outputs{} }
func (m *CommandTask_Outputs) String() string            { return proto.CompactTextString(m) }
func (*CommandTask_Outputs) ProtoMessage()               {}
func (*CommandTask_Outputs) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

func (m *CommandTask_Outputs) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CommandTask_Outputs) GetDirectories() []string {
	if m != nil {
		return m.Directories
	}
	return nil
}

// Describes the timeouts associated with this task.
type CommandTask_Timeouts struct {
	// This specifies the maximum time that the task can run, excluding the
	// time required to download inputs or upload outputs. That is, the worker
	// will terminate the task if it runs longer than this.
	Execution *google_protobuf5.Duration `protobuf:"bytes,1,opt,name=execution" json:"execution,omitempty"`
	// This specifies the maximum amount of time the task can be idle - that is,
	// go without generating some output in either stdout or stderr. If the
	// process is silent for more than the specified time, the worker will
	// terminate the task.
	Idle *google_protobuf5.Duration `protobuf:"bytes,2,opt,name=idle" json:"idle,omitempty"`
	// If the execution or IO timeouts are exceeded, the worker will try to
	// gracefully terminate the task and return any existing logs. However,
	// tasks may be hard-frozen in which case this process will fail. This
	// timeout specifies how long to wait for a terminated task to shut down
	// gracefully (e.g. via SIGTERM) before we bring down the hammer (e.g.
	// SIGKILL on *nix, CTRL_BREAK_EVENT on Windows).
	Shutdown *google_protobuf5.Duration `protobuf:"bytes,3,opt,name=shutdown" json:"shutdown,omitempty"`
}

func (m *CommandTask_Timeouts) Reset()                    { *m = CommandTask_Timeouts{} }
func (m *CommandTask_Timeouts) String() string            { return proto.CompactTextString(m) }
func (*CommandTask_Timeouts) ProtoMessage()               {}
func (*CommandTask_Timeouts) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func (m *CommandTask_Timeouts) GetExecution() *google_protobuf5.Duration {
	if m != nil {
		return m.Execution
	}
	return nil
}

func (m *CommandTask_Timeouts) GetIdle() *google_protobuf5.Duration {
	if m != nil {
		return m.Idle
	}
	return nil
}

func (m *CommandTask_Timeouts) GetShutdown() *google_protobuf5.Duration {
	if m != nil {
		return m.Shutdown
	}
	return nil
}

// Describes the actual outputs from the task.
type CommandOutputs struct {
	// exit_code is only fully reliable if the status' code is OK. If the task
	// exceeded its deadline or was cancelled, the process may still produce an
	// exit code as it is cancelled, and this will be populated, but a successful
	// (zero) is unlikely to be correct unless the status code is OK.
	ExitCode int32 `protobuf:"varint,1,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	// The output files. The blob referenced by the digest should contain
	// one of the following (implementation-dependent):
	//    * A marshalled DirectoryMetadata of the returned filesystem
	//    * A LUCI-style .isolated file
	Outputs *Digest `protobuf:"bytes,2,opt,name=outputs" json:"outputs,omitempty"`
}

func (m *CommandOutputs) Reset()                    { *m = CommandOutputs{} }
func (m *CommandOutputs) String() string            { return proto.CompactTextString(m) }
func (*CommandOutputs) ProtoMessage()               {}
func (*CommandOutputs) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CommandOutputs) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *CommandOutputs) GetOutputs() *Digest {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Can be used as part of CompleteRequest.metadata, or are part of a more
// sophisticated message.
type CommandOverhead struct {
	// The elapsed time between calling Accept and Complete. The server will also
	// have its own idea of what this should be, but this excludes the overhead of
	// the RPCs and the bot response time.
	Duration *google_protobuf5.Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	// The amount of time *not* spent executing the command (ie
	// uploading/downloading files).
	Overhead *google_protobuf5.Duration `protobuf:"bytes,2,opt,name=overhead" json:"overhead,omitempty"`
}

func (m *CommandOverhead) Reset()                    { *m = CommandOverhead{} }
func (m *CommandOverhead) String() string            { return proto.CompactTextString(m) }
func (*CommandOverhead) ProtoMessage()               {}
func (*CommandOverhead) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CommandOverhead) GetDuration() *google_protobuf5.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *CommandOverhead) GetOverhead() *google_protobuf5.Duration {
	if m != nil {
		return m.Overhead
	}
	return nil
}

// The metadata for a file. Similar to the equivalent message in the Remote
// Execution API.
type FileMetadata struct {
	// The path of this file. If this message is part of the
	// CommandResult.output_files fields, the path is relative to the execution
	// root and must correspond to an entry in CommandTask.outputs.files. If this
	// message is part of a Directory message, then the path is relative to the
	// root of that directory.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// A pointer to the contents of the file. The method by which a client
	// retrieves the contents from a CAS system is not defined here.
	Digest *Digest `protobuf:"bytes,2,opt,name=digest" json:"digest,omitempty"`
	// If the file is small enough, its contents may also or alternatively be
	// listed here.
	Contents []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	// Properties of the file
	IsExecutable bool `protobuf:"varint,4,opt,name=is_executable,json=isExecutable" json:"is_executable,omitempty"`
}

func (m *FileMetadata) Reset()                    { *m = FileMetadata{} }
func (m *FileMetadata) String() string            { return proto.CompactTextString(m) }
func (*FileMetadata) ProtoMessage()               {}
func (*FileMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *FileMetadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileMetadata) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *FileMetadata) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *FileMetadata) GetIsExecutable() bool {
	if m != nil {
		return m.IsExecutable
	}
	return false
}

// The metadata for a directory. Similar to the equivalent message in the Remote
// Execution API.
type DirectoryMetadata struct {
	// The path of the directory, as in [FileMetadata.path][google.devtools.remoteworkers.v1test2.FileMetadata.path].
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// A pointer to the contents of the directory, in the form of a marshalled
	// Directory message.
	Digest *Digest `protobuf:"bytes,2,opt,name=digest" json:"digest,omitempty"`
}

func (m *DirectoryMetadata) Reset()                    { *m = DirectoryMetadata{} }
func (m *DirectoryMetadata) String() string            { return proto.CompactTextString(m) }
func (*DirectoryMetadata) ProtoMessage()               {}
func (*DirectoryMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DirectoryMetadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DirectoryMetadata) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

// A reference to the contents of a file or a directory. If the latter, the has
// refers to the hash of a marshalled Directory message. Similar to the
// equivalent message in the Remote Execution API.
type Digest struct {
	// A string-encoded hash (eg "1a2b3c", not the byte array [0x1a, 0x2b, 0x3c])
	// using an implementation-defined hash algorithm (eg SHA-256).
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	// The size of the contents. While this is not strictly required as part of an
	// identifier (after all, any given hash will have exactly one canonical
	// size), it's useful in almost all cases when one might want to send or
	// retrieve blobs of content and is included here for this reason.
	SizeBytes int64 `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes" json:"size_bytes,omitempty"`
}

func (m *Digest) Reset()                    { *m = Digest{} }
func (m *Digest) String() string            { return proto.CompactTextString(m) }
func (*Digest) ProtoMessage()               {}
func (*Digest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Digest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Digest) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

// The contents of a directory. Similar to the equivalent message in the Remote
// Execution API.
type Directory struct {
	// The files in this directory
	Files []*FileMetadata `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	// Any subdirectories
	Directories []*DirectoryMetadata `protobuf:"bytes,2,rep,name=directories" json:"directories,omitempty"`
}

func (m *Directory) Reset()                    { *m = Directory{} }
func (m *Directory) String() string            { return proto.CompactTextString(m) }
func (*Directory) ProtoMessage()               {}
func (*Directory) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Directory) GetFiles() []*FileMetadata {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Directory) GetDirectories() []*DirectoryMetadata {
	if m != nil {
		return m.Directories
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandTask)(nil), "google.devtools.remoteworkers.v1test2.CommandTask")
	proto.RegisterType((*CommandTask_Inputs)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Inputs")
	proto.RegisterType((*CommandTask_Inputs_EnvironmentVariable)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Inputs.EnvironmentVariable")
	proto.RegisterType((*CommandTask_Outputs)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Outputs")
	proto.RegisterType((*CommandTask_Timeouts)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Timeouts")
	proto.RegisterType((*CommandOutputs)(nil), "google.devtools.remoteworkers.v1test2.CommandOutputs")
	proto.RegisterType((*CommandOverhead)(nil), "google.devtools.remoteworkers.v1test2.CommandOverhead")
	proto.RegisterType((*FileMetadata)(nil), "google.devtools.remoteworkers.v1test2.FileMetadata")
	proto.RegisterType((*DirectoryMetadata)(nil), "google.devtools.remoteworkers.v1test2.DirectoryMetadata")
	proto.RegisterType((*Digest)(nil), "google.devtools.remoteworkers.v1test2.Digest")
	proto.RegisterType((*Directory)(nil), "google.devtools.remoteworkers.v1test2.Directory")
}

func init() {
	proto.RegisterFile("google/devtools/remoteworkers/v1test2/command.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcb, 0x6e, 0x13, 0x49,
	0x14, 0x86, 0xe5, 0x38, 0x76, 0xec, 0xe3, 0xcc, 0x64, 0xa6, 0x26, 0x23, 0x19, 0x73, 0x51, 0x64,
	0x84, 0x94, 0x4d, 0xba, 0x85, 0x23, 0xc4, 0x25, 0x0b, 0x44, 0x2e, 0xa0, 0x2c, 0x22, 0x44, 0x29,
	0x02, 0x29, 0x1b, 0xab, 0xec, 0x3e, 0x69, 0x97, 0xd2, 0xee, 0xb2, 0xba, 0xaa, 0x3b, 0x09, 0x1b,
	0x24, 0xde, 0x06, 0x76, 0xf0, 0x0a, 0xbc, 0x18, 0xaa, 0x5b, 0xc7, 0x21, 0x08, 0x1b, 0x2f, 0xd8,
	0x55, 0x9f, 0xd4, 0xff, 0xd5, 0xb9, 0xfc, 0x27, 0x86, 0xed, 0x58, 0x88, 0x38, 0xc1, 0x30, 0xc2,
	0x42, 0x09, 0x91, 0xc8, 0x30, 0xc3, 0xb1, 0x50, 0x78, 0x2e, 0xb2, 0x33, 0xcc, 0x64, 0x58, 0x3c,
	0x54, 0x28, 0x55, 0x2f, 0x1c, 0x8a, 0xf1, 0x98, 0xa5, 0x51, 0x30, 0xc9, 0x84, 0x12, 0xe4, 0x81,
	0x15, 0x05, 0x5e, 0x14, 0x5c, 0x13, 0x05, 0x4e, 0xd4, 0xb9, 0xe7, 0xd8, 0x46, 0x34, 0xc8, 0x4f,
	0xc3, 0x28, 0xcf, 0x98, 0xe2, 0x22, 0xb5, 0x98, 0xee, 0xb7, 0x3a, 0xb4, 0xf6, 0x2c, 0xf8, 0x98,
	0xc9, 0x33, 0xf2, 0x06, 0xea, 0x3c, 0x9d, 0xe4, 0x4a, 0xb6, 0x2b, 0x1b, 0x95, 0xcd, 0x56, 0xef,
	0x69, 0x30, 0xd7, 0x3b, 0xc1, 0x14, 0x23, 0x38, 0x34, 0x00, 0xea, 0x40, 0x04, 0xe1, 0x1f, 0xbc,
	0x98, 0xe0, 0x50, 0x61, 0xd4, 0x17, 0xb9, 0x32, 0xf0, 0x65, 0x03, 0x7f, 0xb6, 0x00, 0xfc, 0xb5,
	0x25, 0xd0, 0x35, 0xcf, 0x74, 0x01, 0xf2, 0x0e, 0x1a, 0x8a, 0x8f, 0x51, 0x68, 0x7c, 0xcd, 0xe0,
	0x77, 0x16, 0xc0, 0x1f, 0x3b, 0x04, 0x2d, 0x61, 0x9d, 0x2f, 0x4b, 0x50, 0xb7, 0x25, 0x91, 0x3b,
	0xd0, 0x64, 0x59, 0x9c, 0x8f, 0x31, 0x35, 0x0d, 0xaa, 0x6e, 0x36, 0xe9, 0x55, 0x80, 0xec, 0x41,
	0xed, 0x94, 0x27, 0x28, 0xdb, 0x4b, 0x1b, 0xd5, 0xcd, 0x56, 0x6f, 0x6b, 0xce, 0xe7, 0xf7, 0x79,
	0x8c, 0x52, 0x51, 0xab, 0x25, 0x1f, 0x2b, 0xf0, 0x3f, 0xa6, 0x05, 0xcf, 0x44, 0xaa, 0xa9, 0xfd,
	0x82, 0x65, 0x9c, 0x0d, 0x34, 0xb5, 0x6a, 0xa8, 0x47, 0x0b, 0x0f, 0x24, 0x38, 0xb8, 0xc2, 0xbe,
	0x75, 0x54, 0xba, 0x8e, 0x37, 0x83, 0xb2, 0xf3, 0x1c, 0xfe, 0xfb, 0xc9, 0x65, 0x42, 0x60, 0x39,
	0x65, 0x63, 0x34, 0xd6, 0x68, 0x52, 0x73, 0x26, 0xeb, 0x50, 0x2b, 0x58, 0x92, 0x63, 0x7b, 0xc9,
	0x04, 0xed, 0x47, 0xe7, 0x05, 0xac, 0xf8, 0xb9, 0xac, 0xfb, 0xae, 0xd8, 0x7e, 0xb9, 0x32, 0x37,
	0xa0, 0x15, 0xf1, 0x0c, 0x87, 0x4a, 0x64, 0xdc, 0x75, 0xac, 0x49, 0xa7, 0x43, 0x9d, 0x4f, 0x15,
	0x68, 0xf8, 0x69, 0x90, 0xc7, 0xd0, 0xc4, 0x0b, 0x1c, 0xe6, 0xda, 0xb9, 0xce, 0x99, 0xb7, 0x7c,
	0x23, 0xbc, 0xb5, 0x83, 0x7d, 0x67, 0x6d, 0x7a, 0x75, 0x97, 0x6c, 0xc1, 0x32, 0x8f, 0x12, 0x9b,
	0xdd, 0x2f, 0x35, 0xe6, 0x1a, 0x79, 0x04, 0x0d, 0x39, 0xca, 0x55, 0x24, 0xce, 0xd3, 0x76, 0x75,
	0x96, 0xa4, 0xbc, 0xda, 0x2d, 0xe0, 0x6f, 0xd7, 0x6f, 0x5f, 0xf5, 0x6d, 0x9d, 0x30, 0x57, 0xfd,
	0xa1, 0x88, 0x6c, 0xbf, 0x6a, 0xb4, 0xa1, 0x03, 0x7b, 0x22, 0x42, 0xf2, 0x0a, 0x56, 0xfc, 0x22,
	0xd8, 0xbc, 0x7e, 0xd3, 0x2a, 0x5e, 0xdd, 0xfd, 0x00, 0x6b, 0xfe, 0xdd, 0x02, 0xb3, 0x11, 0xb2,
	0x48, 0x57, 0xe0, 0x57, 0x7c, 0x76, 0xa3, 0xca, 0xab, 0x5a, 0x26, 0x1c, 0x62, 0x76, 0xaf, 0xca,
	0xab, 0xdd, 0xcf, 0x15, 0x58, 0x7d, 0xc9, 0x13, 0x3c, 0x42, 0xc5, 0x22, 0xa6, 0x98, 0xb6, 0xc8,
	0x84, 0xa9, 0x91, 0xb7, 0x88, 0x3e, 0x93, 0x03, 0xa8, 0x47, 0x26, 0xf1, 0xc5, 0xaa, 0x75, 0x62,
	0xd2, 0x81, 0xc6, 0x50, 0xa4, 0xca, 0xec, 0x9e, 0x9e, 0xcd, 0x2a, 0x2d, 0xbf, 0xc9, 0x7d, 0xf8,
	0x8b, 0xcb, 0xbe, 0x1d, 0xbb, 0xb6, 0xaa, 0xf9, 0x07, 0xd3, 0xa0, 0xab, 0x5c, 0x1e, 0x94, 0xb1,
	0x6e, 0x0a, 0xff, 0xee, 0x3b, 0x83, 0x5d, 0xfe, 0x81, 0x84, 0xbb, 0x3b, 0x50, 0xb7, 0x11, 0xfd,
	0xc8, 0x88, 0xc9, 0xf2, 0x11, 0x7d, 0x26, 0x77, 0x01, 0x24, 0x7f, 0x8f, 0xfd, 0xc1, 0xa5, 0x42,
	0xeb, 0x83, 0x2a, 0x6d, 0xea, 0xc8, 0xae, 0x0e, 0x74, 0xbf, 0x56, 0xa0, 0x59, 0x66, 0x4b, 0x0e,
	0xa7, 0x97, 0xa8, 0xd5, 0xdb, 0x9e, 0x33, 0xa1, 0xe9, 0xd1, 0xf8, 0xcd, 0x3b, 0xb9, 0xb9, 0x79,
	0xad, 0xde, 0x93, 0xb9, 0x2b, 0xfc, 0xa1, 0x7f, 0xd7, 0x76, 0x76, 0xf7, 0xf8, 0x84, 0x3a, 0x4e,
	0x2c, 0x12, 0x96, 0xc6, 0x81, 0xc8, 0xe2, 0x30, 0xc6, 0xd4, 0x58, 0x28, 0xb4, 0x7f, 0x62, 0x13,
	0x2e, 0x67, 0xfc, 0xd4, 0xed, 0x5c, 0x8b, 0x0e, 0xea, 0x46, 0xbe, 0xfd, 0x3d, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0x77, 0x2c, 0xa3, 0x28, 0x07, 0x00, 0x00,
}
