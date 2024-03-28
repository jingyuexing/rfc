package rfc


type TCP struct {
	SourcePort uint16 //16bit
	DestinationPort uint16 //16bit
	Acknowledgment uint32 //32bit
	Sequence uint32 //32bit
	Offset uint8 //4bit
	Reserved uint8 // 6bit
	URG bool // 1bit
	ACK bool // 1bit
	PSH bool // 1bit
	RST bool // 1bit
	SYN bool // 1bit
	FIN bool // 1bit
	Window uint16 //16bit
	CheckSum uint16 // 16bit
	UrgentPointer uint16 // 16bit
	Options uint32 // max 25bit
	Padding uint8 // max 8bit 
	payload []byte
}

func (tcp TCP) ToBytes() []byte{
	mask4 := (1 << 4) - 1
	mask6 := (1 << 6) - 1
	mask8 := (1 << 8) - 1

	result := make([]byte,0)

	result = Fill(result,uint(tcp.SourcePort))

	result = Fill(result,uint(tcp.DestinationPort))

	result = Fill(result,uint(tcp.Sequence))
	result = Fill(result,uint(tcp.Acknowledgment))

	var cache uint

	cache = cache | uint(tcp.Offset) & uint(mask4)
	cache = (cache << 4) | uint(tcp.Reserved) & uint(mask6)
	ControlBits := boolToInt(tcp.FIN) | (boolToInt(tcp.SYN) << 1) | (boolToInt(tcp.RST) << 2) | (boolToInt(tcp.PSH) << 3) | (boolToInt(tcp.ACK) << 4) | (boolToInt(tcp.URG) << 5)
	cache = cache | uint(tcp.Offset) & uint(mask4)
	cache = (cache << 4) | uint(tcp.Reserved) & uint(mask6)
	cache = (cache << 6) | uint(ControlBits) & uint(mask6)
	cache = (cache << 6) | uint(tcp.Window)
	result = Fill(result,cache)
	cache = 0
	result = Fill(result,uint(tcp.CheckSum))
	result = Fill(result,uint(tcp.UrgentPointer))
	result = Fill(result,(uint(tcp.Options) << 8) | uint(tcp.Padding) & uint(mask8))
	result = append(result, tcp.payload...)
	return result

}

func (tcp TCP) SetData(data ...byte){
	
}
