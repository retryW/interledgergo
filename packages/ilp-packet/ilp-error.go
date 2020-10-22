package ilpacket

type IlpError string

// ILP Reject error codes
const (
	// Final errors
	BAD_REQUEST = IlpError("F00")
	INVALID_PACKET = IlpError("F01")
	UNREACHABLE = IlpError("F02")
	INVALID_AMOUNT = IlpError("F03")
	INSUFFICIENT_DESTINATION_AMOUNT = IlpError("F04")
	WRONG_CONDITION = IlpError("F05")
	UNEXPECTED_PAYMENT = IlpError("F06")
	CANNOT_RECEIVE = IlpError("F07")
	AMOUNT_TOO_LARGE = IlpError("F08")
	F_APPLICATION_ERROR = IlpError("F99")

	// Temporary errors
	INTERNAL_ERROR = IlpError("T00")
	PEER_UNREACHABLE = IlpError("T01")
	PEER_BUSY = IlpError("T02")
	CONNECTOR_BUSY = IlpError("T03")
	INSUFFICIENT_LIQUIDITY = IlpError("T04")
	RATE_LIMITED = IlpError("T05")
	T_APPLICATION_ERROR = IlpError("T99")

	// Relative errors
	TRANSFER_TIMED_OUT = IlpError("R00")
	INSUFFICIENT_SOURCE_AMOUNT = IlpError("R01")
	INSUFFICIENT_TIMEOUT = IlpError("R02")
	R_APPLICATION_ERROR = IlpError("R99")
)