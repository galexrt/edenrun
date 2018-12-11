// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto

package v1alpha // import "github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha"

import testing "testing"
import math_rand "math/rand"
import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
import _ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func BenchmarkVariableProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Variable, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedVariable(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVariableProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedVariable(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &Variable{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVariableSpecProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VariableSpec, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedVariableSpec(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVariableSpecProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedVariableSpec(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &VariableSpec{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*GetRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedGetRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedGetRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &GetRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*GetResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedGetResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedGetResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &GetResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ListRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedListRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedListRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &ListRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ListResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedListResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedListResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &ListResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*AddRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedAddRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedAddRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &AddRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*AddResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedAddResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedAddResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &AddResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*UpdateRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedUpdateRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedUpdateRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &UpdateRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*UpdateResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedUpdateResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedUpdateResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &UpdateResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*DeleteRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedDeleteRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedDeleteRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &DeleteRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*DeleteResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedDeleteResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedDeleteResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &DeleteResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchRequestProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*WatchRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedWatchRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchRequestProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedWatchRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &WatchRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchResponseProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*WatchResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedWatchResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchResponseProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_golang_protobuf_proto.Marshal(NewPopulatedWatchResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &WatchResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_golang_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVariableSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Variable, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedVariable(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVariableSpecSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VariableSpec, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedVariableSpec(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*GetRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedGetRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*GetResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedGetResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ListRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedListRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkListResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ListResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedListResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*AddRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedAddRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAddResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*AddResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedAddResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*UpdateRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedUpdateRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUpdateResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*UpdateResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedUpdateResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*DeleteRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedDeleteRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDeleteResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*DeleteResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedDeleteResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchRequestSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*WatchRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedWatchRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkWatchResponseSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*WatchResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedWatchResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
