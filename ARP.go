package rfc

// ether_type$XEROX_PUP,
// ether_type$DOD_INTERNET,
// ether_type$CHAOS,
// ether_type$ADDRESS_RESOLUTION.

type ETH uint

const (
	ETH_TYPE_XEROX_PUP ETH = iota
	ETH_TYPE_DOD_INTERNET
	ETH_TYPE_CHAOS
	ETH_TYPE_ADDRESS_RESOLUTION
)

const (
	OP_REPLY = iota
	OP_REQUEST
)

type ARP struct {
	// 48bit
	DestinationAddress uint64
	// 48bit
	SenderAddress uint64
	// 16bit
	ProtocolType uint16
	// mac address
	// 16bit
	HardwareAddress uint16
	// protocol address space
	// 16bit
	ProtocolAddress uint16

	HardwareLength uint8

	ProtocolLength uint8
	// opcode (ares_op$REQUEST = 1 | ares_op$REPLY = 2)
	Opcode uint16
	// Hardware address of sender of this packet, length from the `HardwareLength` field.
	SenderHardwareAddress uint32

	// Protocol address of sender of this packet, length from the `ProtocolLength` field.
	SenderIPAddress uint64
	// Hardware address of target of this packet (if known).
	TargetHardwareAddress uint32
	// Protocol address of target.
	TargetIPAddress uint64
}

func (arp ARP) ToBytes() []byte{
	return nil
}

func (arp ARP) SetData(data ...byte){
	
}
