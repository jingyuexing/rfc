package rfc

type RFC interface {
	ToBytes() []byte
	SetData(data []byte)
}