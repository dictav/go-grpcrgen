// automatically generated by the FlatBuffers compiler, do not modify

package myservice

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Request struct {
	_tab flatbuffers.Table
}

func GetRootAsRequest(buf []byte, offset flatbuffers.UOffsetT) *Request {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Request{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Request) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Request) Table() flatbuffers.Table {
	return rcv._tab
}

func RequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func RequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
