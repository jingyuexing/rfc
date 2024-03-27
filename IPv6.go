package rfc

import "errors"

type IPv6 struct {
	Version uint8 // 4bit
	TrafficClass uint8 // 8bit
	FlowLabel uint32 // 20bit
	PayloadLength uint16 // 16bit
	NextHeader uint8 // 8bit
	HopLimit uint8  // 8bit
	SourceAddress [16]byte // 128bit
	DestinationAddress [16]byte //128bit
	Payload []byte
}

func fill16( target []byte,data [16]byte) []byte{
	for _,item := range data {
		target = append(target, item)
	}
	return target
}

func (v6 IPv6) ToBytes() []byte{
	result := make([]byte,0)
	mask4 := (1 << 4) - 1
	mask20 := (1 << 20) - 1
	var cache uint32

	cache = cache | (uint32(v6.Version) & uint32(mask4))
	cache = (cache << 4) | (uint32(v6.TrafficClass) & 0xff)
	cache = (cache << 8) | v6.FlowLabel & uint32(mask20)
	result = Fill(result,uint(cache))
	cache = 0
	cache = cache | uint32(v6.PayloadLength)
	cache = (cache << 16) | (uint32(v6.NextHeader) << 8 | uint32(v6.HopLimit))
	result = Fill(result,uint(cache))
	result = fill16(result,v6.SourceAddress)
	result = fill16(result,v6.DestinationAddress)
	result = append(result, v6.Payload...)
	return result
}

func (v6 IPv6) SetData(data []byte) error{
	if len(v6.Payload) > ((1 << 16) -1) {
		return errors.New("out of length")
	}
	v6.Payload = data
	v6.PayloadLength = uint16(len(v6.Payload))
	return nil
}