package mvalue

import "C"

type MValueType = uint8

const (
	None MValueType = iota
	Nil
	Bool
	Int
	Uint
	Double
	String
	List
	Dict
	BaseObject
	Function
	Vector3
	RGBA
	ByteArray
	Vector2
)

type Convertable interface {
	OnRead(reader *Reader)
	OnWrite(writer *Writer)
}

type MValue struct {
	Type  MValueType
	Value interface{}
}

func (m *MValue) ToC() C.struct_goValue {
	return C.struct_goValue{
		typ: C.uchar(m.Type),
	}
}
