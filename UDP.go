package rfc

type UDP struct {
	SourcePort uint16
	DestinationPort uint16
	Length uint16
	Checksum uint16
	payload []byte
}

func (udp UDP) ToBytes() []byte{
	result := make([]byte,0)
	result = Fill(result,uint(udp.SourcePort))
	result = Fill(result,uint(udp.DestinationPort))
	result = Fill(result,uint(udp.Length))
	result = Fill(result,uint(udp.Checksum))
	result = append(result, udp.payload...)
	return result
}

func (udp UDP) SetData(data ...byte){
	udp.payload = data
}
