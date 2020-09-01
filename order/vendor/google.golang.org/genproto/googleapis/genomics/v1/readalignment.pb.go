// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/genomics/v1/readalignment.proto

package genomics

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A linear alignment can be represented by one CIGAR string. Describes the
// mapped position and local alignment of the read to the reference.
type LinearAlignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The position of this alignment.
	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// The mapping quality of this alignment. Represents how likely
	// the read maps to this position as opposed to other locations.
	//
	// Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to
	// the nearest integer.
	MappingQuality int32 `protobuf:"varint,2,opt,name=mapping_quality,json=mappingQuality,proto3" json:"mapping_quality,omitempty"`
	// Represents the local alignment of this sequence (alignment matches, indels,
	// etc) against the reference.
	Cigar []*CigarUnit `protobuf:"bytes,3,rep,name=cigar,proto3" json:"cigar,omitempty"`
}

func (x *LinearAlignment) Reset() {
	*x = LinearAlignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_genomics_v1_readalignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinearAlignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinearAlignment) ProtoMessage() {}

func (x *LinearAlignment) ProtoReflect() protoreflect.Message {
	mi := &file_google_genomics_v1_readalignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinearAlignment.ProtoReflect.Descriptor instead.
func (*LinearAlignment) Descriptor() ([]byte, []int) {
	return file_google_genomics_v1_readalignment_proto_rawDescGZIP(), []int{0}
}

func (x *LinearAlignment) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *LinearAlignment) GetMappingQuality() int32 {
	if x != nil {
		return x.MappingQuality
	}
	return 0
}

func (x *LinearAlignment) GetCigar() []*CigarUnit {
	if x != nil {
		return x.Cigar
	}
	return nil
}

// A read alignment describes a linear alignment of a string of DNA to a
// [reference sequence][google.genomics.v1.Reference], in addition to metadata
// about the fragment (the molecule of DNA sequenced) and the read (the bases
// which were read by the sequencer). A read is equivalent to a line in a SAM
// file. A read belongs to exactly one read group and exactly one
// [read group set][google.genomics.v1.ReadGroupSet].
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
//
// ### Reverse-stranded reads
//
// Mapped reads (reads having a non-null `alignment`) can be aligned to either
// the forward or the reverse strand of their associated reference. Strandedness
// of a mapped read is encoded by `alignment.position.reverseStrand`.
//
// If we consider the reference to be a forward-stranded coordinate space of
// `[0, reference.length)` with `0` as the left-most position and
// `reference.length` as the right-most position, reads are always aligned left
// to right. That is, `alignment.position.position` always refers to the
// left-most reference coordinate and `alignment.cigar` describes the alignment
// of this read to the reference from left to right. All per-base fields such as
// `alignedSequence` and `alignedQuality` share this same left-to-right
// orientation; this is true of reads which are aligned to either strand. For
// reverse-stranded reads, this means that `alignedSequence` is the reverse
// complement of the bases that were originally reported by the sequencing
// machine.
//
// ### Generating a reference-aligned sequence string
//
// When interacting with mapped reads, it's often useful to produce a string
// representing the local alignment of the read to reference. The following
// pseudocode demonstrates one way of doing this:
//
//     out = ""
//     offset = 0
//     for c in read.alignment.cigar {
//       switch c.operation {
//       case "ALIGNMENT_MATCH", "SEQUENCE_MATCH", "SEQUENCE_MISMATCH":
//         out += read.alignedSequence[offset:offset+c.operationLength]
//         offset += c.operationLength
//         break
//       case "CLIP_SOFT", "INSERT":
//         offset += c.operationLength
//         break
//       case "PAD":
//         out += repeat("*", c.operationLength)
//         break
//       case "DELETE":
//         out += repeat("-", c.operationLength)
//         break
//       case "SKIP":
//         out += repeat(" ", c.operationLength)
//         break
//       case "CLIP_HARD":
//         break
//       }
//     }
//     return out
//
// ### Converting to SAM's CIGAR string
//
// The following pseudocode generates a SAM CIGAR string from the
// `cigar` field. Note that this is a lossy conversion
// (`cigar.referenceSequence` is lost).
//
//     cigarMap = {
//       "ALIGNMENT_MATCH": "M",
//       "INSERT": "I",
//       "DELETE": "D",
//       "SKIP": "N",
//       "CLIP_SOFT": "S",
//       "CLIP_HARD": "H",
//       "PAD": "P",
//       "SEQUENCE_MATCH": "=",
//       "SEQUENCE_MISMATCH": "X",
//     }
//     cigarStr = ""
//     for c in read.alignment.cigar {
//       cigarStr += c.operationLength + cigarMap[c.operation]
//     }
//     return cigarStr
type Read struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-generated read ID, unique across all reads. This is different
	// from the `fragmentName`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the read group this read belongs to. A read belongs to exactly
	// one read group. This is a server-generated ID which is distinct from SAM's
	// RG tag (for that value, see
	// [ReadGroup.name][google.genomics.v1.ReadGroup.name]).
	ReadGroupId string `protobuf:"bytes,2,opt,name=read_group_id,json=readGroupId,proto3" json:"read_group_id,omitempty"`
	// The ID of the read group set this read belongs to. A read belongs to
	// exactly one read group set.
	ReadGroupSetId string `protobuf:"bytes,3,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName string `protobuf:"bytes,4,opt,name=fragment_name,json=fragmentName,proto3" json:"fragment_name,omitempty"`
	// The orientation and the distance between reads from the fragment are
	// consistent with the sequencing protocol (SAM flag 0x2).
	ProperPlacement bool `protobuf:"varint,5,opt,name=proper_placement,json=properPlacement,proto3" json:"proper_placement,omitempty"`
	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment bool `protobuf:"varint,6,opt,name=duplicate_fragment,json=duplicateFragment,proto3" json:"duplicate_fragment,omitempty"`
	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength int32 `protobuf:"varint,7,opt,name=fragment_length,json=fragmentLength,proto3" json:"fragment_length,omitempty"`
	// The read number in sequencing. 0-based and less than numberReads. This
	// field replaces SAM flag 0x40 and 0x80.
	ReadNumber int32 `protobuf:"varint,8,opt,name=read_number,json=readNumber,proto3" json:"read_number,omitempty"`
	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads int32 `protobuf:"varint,9,opt,name=number_reads,json=numberReads,proto3" json:"number_reads,omitempty"`
	// Whether this read did not pass filters, such as platform or vendor quality
	// controls (SAM flag 0x200).
	FailedVendorQualityChecks bool `protobuf:"varint,10,opt,name=failed_vendor_quality_checks,json=failedVendorQualityChecks,proto3" json:"failed_vendor_quality_checks,omitempty"`
	// The linear alignment for this alignment record. This field is null for
	// unmapped reads.
	Alignment *LinearAlignment `protobuf:"bytes,11,opt,name=alignment,proto3" json:"alignment,omitempty"`
	// Whether this alignment is secondary. Equivalent to SAM flag 0x100.
	// A secondary alignment represents an alternative to the primary alignment
	// for this read. Aligners may return secondary alignments if a read can map
	// ambiguously to multiple coordinates in the genome. By convention, each read
	// has one and only one alignment where both `secondaryAlignment`
	// and `supplementaryAlignment` are false.
	SecondaryAlignment bool `protobuf:"varint,12,opt,name=secondary_alignment,json=secondaryAlignment,proto3" json:"secondary_alignment,omitempty"`
	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800.
	// Supplementary alignments are used in the representation of a chimeric
	// alignment. In a chimeric alignment, a read is split into multiple
	// linear alignments that map to different reference contigs. The first
	// linear alignment in the read will be designated as the representative
	// alignment; the remaining linear alignments will be designated as
	// supplementary alignments. These alignments may have different mapping
	// quality scores. In each linear alignment in a chimeric alignment, the read
	// will be hard clipped. The `alignedSequence` and
	// `alignedQuality` fields in the alignment record will only
	// represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `protobuf:"varint,13,opt,name=supplementary_alignment,json=supplementaryAlignment,proto3" json:"supplementary_alignment,omitempty"`
	// The bases of the read sequence contained in this alignment record,
	// **without CIGAR operations applied** (equivalent to SEQ in SAM).
	// `alignedSequence` and `alignedQuality` may be
	// shorter than the full read sequence and quality. This will occur if the
	// alignment is part of a chimeric alignment, or if the read was trimmed. When
	// this occurs, the CIGAR for this read will begin/end with a hard clip
	// operator that will indicate the length of the excised sequence.
	AlignedSequence string `protobuf:"bytes,14,opt,name=aligned_sequence,json=alignedSequence,proto3" json:"aligned_sequence,omitempty"`
	// The quality of the read sequence contained in this alignment record
	// (equivalent to QUAL in SAM).
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedQuality []int32 `protobuf:"varint,15,rep,packed,name=aligned_quality,json=alignedQuality,proto3" json:"aligned_quality,omitempty"`
	// The mapping of the primary alignment of the
	// `(readNumber+1)%numberReads` read in the fragment. It replaces
	// mate position and mate strand in SAM.
	NextMatePosition *Position `protobuf:"bytes,16,opt,name=next_mate_position,json=nextMatePosition,proto3" json:"next_mate_position,omitempty"`
	// A map of additional read alignment information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info map[string]*_struct.ListValue `protobuf:"bytes,17,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Read) Reset() {
	*x = Read{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_genomics_v1_readalignment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Read) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Read) ProtoMessage() {}

func (x *Read) ProtoReflect() protoreflect.Message {
	mi := &file_google_genomics_v1_readalignment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Read.ProtoReflect.Descriptor instead.
func (*Read) Descriptor() ([]byte, []int) {
	return file_google_genomics_v1_readalignment_proto_rawDescGZIP(), []int{1}
}

func (x *Read) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Read) GetReadGroupId() string {
	if x != nil {
		return x.ReadGroupId
	}
	return ""
}

func (x *Read) GetReadGroupSetId() string {
	if x != nil {
		return x.ReadGroupSetId
	}
	return ""
}

func (x *Read) GetFragmentName() string {
	if x != nil {
		return x.FragmentName
	}
	return ""
}

func (x *Read) GetProperPlacement() bool {
	if x != nil {
		return x.ProperPlacement
	}
	return false
}

func (x *Read) GetDuplicateFragment() bool {
	if x != nil {
		return x.DuplicateFragment
	}
	return false
}

func (x *Read) GetFragmentLength() int32 {
	if x != nil {
		return x.FragmentLength
	}
	return 0
}

func (x *Read) GetReadNumber() int32 {
	if x != nil {
		return x.ReadNumber
	}
	return 0
}

func (x *Read) GetNumberReads() int32 {
	if x != nil {
		return x.NumberReads
	}
	return 0
}

func (x *Read) GetFailedVendorQualityChecks() bool {
	if x != nil {
		return x.FailedVendorQualityChecks
	}
	return false
}

func (x *Read) GetAlignment() *LinearAlignment {
	if x != nil {
		return x.Alignment
	}
	return nil
}

func (x *Read) GetSecondaryAlignment() bool {
	if x != nil {
		return x.SecondaryAlignment
	}
	return false
}

func (x *Read) GetSupplementaryAlignment() bool {
	if x != nil {
		return x.SupplementaryAlignment
	}
	return false
}

func (x *Read) GetAlignedSequence() string {
	if x != nil {
		return x.AlignedSequence
	}
	return ""
}

func (x *Read) GetAlignedQuality() []int32 {
	if x != nil {
		return x.AlignedQuality
	}
	return nil
}

func (x *Read) GetNextMatePosition() *Position {
	if x != nil {
		return x.NextMatePosition
	}
	return nil
}

func (x *Read) GetInfo() map[string]*_struct.ListValue {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_google_genomics_v1_readalignment_proto protoreflect.FileDescriptor

var file_google_genomics_v1_readalignment_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x69, 0x67, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x38, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x69, 0x67, 0x61, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x67, 0x61, 0x72, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x05, 0x63, 0x69, 0x67, 0x61, 0x72, 0x22, 0xec, 0x06, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x11, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2d, 0x0a, 0x12, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x75, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x19, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x09,
	0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x6c, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x37, 0x0a, 0x17, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x79, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x16, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79,
	0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x61,
	0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x4a, 0x0a,
	0x12, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x53, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6d, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x63, 0x73, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_genomics_v1_readalignment_proto_rawDescOnce sync.Once
	file_google_genomics_v1_readalignment_proto_rawDescData = file_google_genomics_v1_readalignment_proto_rawDesc
)

func file_google_genomics_v1_readalignment_proto_rawDescGZIP() []byte {
	file_google_genomics_v1_readalignment_proto_rawDescOnce.Do(func() {
		file_google_genomics_v1_readalignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_genomics_v1_readalignment_proto_rawDescData)
	})
	return file_google_genomics_v1_readalignment_proto_rawDescData
}

var file_google_genomics_v1_readalignment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_genomics_v1_readalignment_proto_goTypes = []interface{}{
	(*LinearAlignment)(nil),   // 0: google.genomics.v1.LinearAlignment
	(*Read)(nil),              // 1: google.genomics.v1.Read
	nil,                       // 2: google.genomics.v1.Read.InfoEntry
	(*Position)(nil),          // 3: google.genomics.v1.Position
	(*CigarUnit)(nil),         // 4: google.genomics.v1.CigarUnit
	(*_struct.ListValue)(nil), // 5: google.protobuf.ListValue
}
var file_google_genomics_v1_readalignment_proto_depIdxs = []int32{
	3, // 0: google.genomics.v1.LinearAlignment.position:type_name -> google.genomics.v1.Position
	4, // 1: google.genomics.v1.LinearAlignment.cigar:type_name -> google.genomics.v1.CigarUnit
	0, // 2: google.genomics.v1.Read.alignment:type_name -> google.genomics.v1.LinearAlignment
	3, // 3: google.genomics.v1.Read.next_mate_position:type_name -> google.genomics.v1.Position
	2, // 4: google.genomics.v1.Read.info:type_name -> google.genomics.v1.Read.InfoEntry
	5, // 5: google.genomics.v1.Read.InfoEntry.value:type_name -> google.protobuf.ListValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_genomics_v1_readalignment_proto_init() }
func file_google_genomics_v1_readalignment_proto_init() {
	if File_google_genomics_v1_readalignment_proto != nil {
		return
	}
	file_google_genomics_v1_cigar_proto_init()
	file_google_genomics_v1_position_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_genomics_v1_readalignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinearAlignment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_genomics_v1_readalignment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Read); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_genomics_v1_readalignment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_genomics_v1_readalignment_proto_goTypes,
		DependencyIndexes: file_google_genomics_v1_readalignment_proto_depIdxs,
		MessageInfos:      file_google_genomics_v1_readalignment_proto_msgTypes,
	}.Build()
	File_google_genomics_v1_readalignment_proto = out.File
	file_google_genomics_v1_readalignment_proto_rawDesc = nil
	file_google_genomics_v1_readalignment_proto_goTypes = nil
	file_google_genomics_v1_readalignment_proto_depIdxs = nil
}
