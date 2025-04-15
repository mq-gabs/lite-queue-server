package protocol

const (
	Seperator byte = 0x10

	RequestNewQueue byte = 0x30
	RequestPush     byte = 0x31
	RequestPop      byte = 0x32

	ResponseSuccessEmpty   byte = 0x60
	ResponseSuccessContent byte = 0x61
	ResponseErrorEmpty     byte = 0x62
	ResponseErrorContent   byte = 0x63
)
