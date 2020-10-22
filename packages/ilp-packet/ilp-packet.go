// Package ilp-packet provided utilities to serialize and deserialize interledger packets
package ilpacket


import (
	"regexp"
	"errors"
	"time"
	"bytes"
)

type Address string
type Timestamp string
type IlpPacketType uint8

type IlpPacket struct {
	packetType IlpPacketType
	typeString string
	data bytes.Buffer // Not a buffer, maybe make this an interface ? Needs to be one of 3 packet structs
}

// ILP Packet structures
type IlpPrepare struct {
	amount string
	executionCondition bytes.Buffer
	expiresAt time.Time
	destination string
	data bytes.Buffer
}

type IlpFulfill struct {
	fulfillment bytes.Buffer
	data bytes.Buffer
}

type IlpReject struct {
	code string
	triggeredBy string
	message string
	data bytes.Buffer
}


const (
	// ILP Address requirements.
	ILP_ADDRESS_REGEX = `^(g|private|example|peer|self|test[1-3]?|local)([.][a-zA-Z0-9_~-]+)+$`
	ILP_ADDRESS_MAX_LENGTH = 1023

	// ILP Timestamp requirements
	ILP_TIMESTAMP_LENGTH = 17

	// ILP Packet Types
	ILP_PREPARE = IlpPacketType(12)
	ILP_FULFILL = IlpPacketType(13)
	ILP_REJECT = IlpPacketType(14)

	// Readable names for special character that may appear in ILP addresses
	hyphen = "-"
	period = "."
	underscore = "_"
	tilde = "~"
)

// ILP Packet serialization functions
func serializeIlpPacket(packet IlpPacket) IlpPacket {
	switch packet.packetType {
		case ILP_PREPARE:
			return serializeIlpPrepare(packet.data)
		case ILP_FULFILL:
			return serializeIlpFulfill(packet.data)
		case ILP_REJECT:
			return serializeIlpReject(packet.data)
		default:
			errors.New("object has invalid type")
	}
}

// TODO - Implement
func serializeIlpPrepare(bytes.Buffer) IlpPrepare {

}
// TODO - Implement
func serializeIlpFulfill(bytes.Buffer) IlpFulfill {

}
// TODO - Implement
func serializeIlpReject(bytes.Buffer) IlpReject{

}

// NewAddress returns an ILP Address provided it passes validation.
func NewAddress(a Address) (Address, error) {

	matched, _ := regexp.MatchString(ILP_ADDRESS_REGEX, string(a))
	if matched && len(a) <= ILP_ADDRESS_MAX_LENGTH{
		return a, nil
	}
	
	return "", errors.New("invalid ILP address")
}

// TODO - Implement
func NewTimestamp(t Timestamp) (Timestamp, error) {
	return "", nil
}

// TODO - Implement
func DateToInterledgerTime(t time.Time) time.Time {
	return time.Now()
}

// TODO - Implement
func DateToGeneralTime(t time.Time) string {
	return ""
}