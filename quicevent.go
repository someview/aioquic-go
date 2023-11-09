package aioquic

// QuicEvent BasicConstraint
type QuicEvent interface {
}

type ConnectionIdIssued struct {
	ConnectionID []byte
}

type ConnectionIdRetired struct {
	ConnectionID []byte
}

// The ConnectionTerminated event is fired when the QUIC connection is terminated.
type ConnectionTerminated struct {
	// The error code which was specified when closing the connection
	ErrorCode int
	// The frame type which caused the connection to be closed, or `None`
	FrameType int
	// The human-readable reason for which the connection was closed.
	ReasonPhrase string
}

type DatagramFrameReceived struct {
	//  The DatagramFrameReceived event is fired when a DATAGRAM frame is received
	Data []byte
}

// HandshakeCompleted The HandshakeCompleted event is fired when the TLS handshake completes.
type HandshakeCompleted struct {
	// The protocol which was negotiated using ALPN, or `None`
	AlpnProtocol string
	// Whether early (0-RTT) data was accepted by the remote peer
	EarlyDataAccepted bool
	// Whether a TLS session was resumed.
	SessionResumed bool
}

// The PingAcknowledged event is fired when a PING frame is acknowledged.
type PingAcknowledged struct {
	UID int
}

type ProtocolNegotiated struct {
	AlpnProtocol string
}

// The StopSendingReceived event is fired when the remote peer requests
//
//	stopping data transmission on a stream.
type StopSendingReceived struct {
	ErrorCode int
	StreamID  int
}

// The StreamDataReceived event is fired whenever data is received on a
//
//	stream.
type StreamDataReceived struct {
	Data []byte
	// "Whether the STREAM frame had the FIN bit set
	EndStream bool
	StreamID  int
}

// The StreamReset event is fired when the remote peer resets a stream.
type StreamReset struct {
	ErrorCode int
	StreamID  int
}
