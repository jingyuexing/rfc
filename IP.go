package rfc
type IP struct {
	Version uint8 //4bit
	IHL uint8 //4bit
	Type uint8 // 8bit
	Length uint16 //16bit
	Identification uint16 //16bit
	Flag uint8 // 3bit
	Offset uint16 //13bit
	Time2Live uint8 //8bit
	Protocol uint8 //8bit
	Checksum uint16 //8bit
	SourceAddress uint32 // 32bit
	DestinationAddress uint32 // 32bit
	Options uint16
	Padding uint8
	payload []byte
}

func (ip IP) ToBytes() []byte{
	mask4 := (1 << 4) - 1
	mask3 := (1 << 3) - 1
	mask13 := (1 << 13) - 1
	result := make([]byte,0)
	var cache uint32
	cache = cache | uint32(ip.Version)
	cache = cache << 4 | ((uint32(ip.IHL) & uint32(mask4)))
	cache = cache << 8 | uint32(ip.Type)
	cache = cache << 16 | uint32(ip.Length)
	result = Fill(result,uint(cache))
	cache = 0
	result = Fill(result,uint(ip.Identification))
	result = Fill(result,uint(ip.Flag & uint8(mask3)) << 13 | uint(ip.Offset & uint16(mask13)))
	result = Fill(result,uint(ip.Time2Live))
	result = Fill(result,uint(ip.Protocol))
	result = Fill(result,uint(ip.Checksum))
	result = Fill(result,uint(ip.SourceAddress))
	result = Fill(result,uint(ip.DestinationAddress))
	result = Fill(result,uint(ip.Options))
	result = Fill(result,uint(ip.Padding))
	return result
}

func (ip IP) SetData(data ...byte){
	ip.payload = append(ip.payload, data...)
}

func (ip IP) Clear(){
	ip.payload = []byte{}
}
