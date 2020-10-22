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
	data bytes.Buffer
}
type IlpPrepare struct {
	amount string
	executionCondition bytes.Buffer
	expiresAt time.Time
	destination string
	data bytes.Buffer
}

const (
	// ILP Address requirements.
	ILP_ADDRESS_REGEX = `^(g|private|example|peer|self|test[1-3]?|local)([.][a-zA-Z0-9_~-]+)+$`
	ILP_ADDRESS_MAX_LENGTH = 1023

	// ILP Timestamp requirements
	ILP_TIMESTAMP_LENGTH = 17

	// Readable names for special character that may appear in ILP addresses
	hyphen = "-"
	period = "."
	underscore = "_"
	tilde = "~"
)

// NewAddress returns an ILP Address provided it passes validation.
func NewAddress(a Address) (Address, error) {
	matched, _ := regexp.MatchString(ILP_ADDRESS_REGEX, string(a))
	if matched && len(a) <= ILP_ADDRESS_MAX_LENGTH{
		return a, nil
	} else {
		return "", errors.New("invalid ILP address")
	}
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