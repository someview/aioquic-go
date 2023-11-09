package aioquic

import (
	"context"
	"net"
	"time"
)

type Config struct {
	// GetConfigForClient is called for incoming connections.
	// If the error is not nil, the connection attempt is refused.
	GetConfigForClient func(info *ClientHelloInfo) (*Config, error)
	// The QUIC versions that can be negotiated.
	// If not set, it uses all versions available.
	Versions []VersionNumber
	// HandshakeIdleTimeout is the idle timeout before completion of the handshake.
	// If we don't receive any packet from the peer within this time, the connection attempt is aborted.
	// Additionally, if the handshake doesn't complete in twice this time, the connection attempt is also aborted.
	// If this value is zero, the timeout is set to 5 seconds.
	HandshakeIdleTimeout time.Duration
	// MaxIdleTimeout is the maximum duration that may pass without any incoming network activity.
	// The actual value for the idle timeout is the minimum of this value and the peer's.
	// This value only applies after the handshake has completed.
	// If the timeout is exceeded, the connection is closed.
	// If this value is zero, the timeout is set to 30 seconds.
	MaxIdleTimeout time.Duration
	// RequireAddressValidation determines if a QUIC Retry packet is sent.
	// This allows the server to verify the client's address, at the cost of increasing the handshake latency by 1 RTT.
	// See https://datatracker.ietf.org/doc/html/rfc9000#section-8 for details.
	// If not set, every client is forced to prove its remote address.
	RequireAddressValidation func(net.Addr) bool
	// The TokenStore stores tokens received from the server.
	// Tokens are used to skip address validation on future connection attempts.
	// The key used to store tokens is the ServerName from the tls.Config, if set
	// otherwise the token is associated with the server's IP address.
	TokenStore TokenStore
	// InitialStreamReceiveWindow is the initial size of the stream-level flow control window for receiving data.
	// If the application is consuming data quickly enough, the flow control auto-tuning algorithm
	// will increase the window up to MaxStreamReceiveWindow.
	// If this value is zero, it will default to 512 KB.
	// Values larger than the maximum varint (quicvarint.Max) will be clipped to that value.
	InitialStreamReceiveWindow uint64
	// MaxStreamReceiveWindow is the maximum stream-level flow control window for receiving data.
	// If this value is zero, it will default to 6 MB.
	// Values larger than the maximum varint (quicvarint.Max) will be clipped to that value.
	MaxStreamReceiveWindow uint64
	// InitialConnectionReceiveWindow is the initial size of the stream-level flow control window for receiving data.
	// If the application is consuming data quickly enough, the flow control auto-tuning algorithm
	// will increase the window up to MaxConnectionReceiveWindow.
	// If this value is zero, it will default to 512 KB.
	// Values larger than the maximum varint (quicvarint.Max) will be clipped to that value.
	InitialConnectionReceiveWindow uint64
	// MaxConnectionReceiveWindow is the connection-level flow control window for receiving data.
	// If this value is zero, it will default to 15 MB.
	// Values larger than the maximum varint (quicvarint.Max) will be clipped to that value.
	MaxConnectionReceiveWindow uint64
	// AllowConnectionWindowIncrease is called every time the connection flow controller attempts
	// to increase the connection flow control window.
	// If set, the caller can prevent an increase of the window. Typically, it would do so to
	// limit the memory usage.
	// To avoid deadlocks, it is not valid to call other functions on the connection or on streams
	// in this callback.
	AllowConnectionWindowIncrease func(conn Connection, delta uint64) bool
	// MaxIncomingStreams is the maximum number of concurrent bidirectional streams that a peer is allowed to open.
	// If not set, it will default to 100.
	// If set to a negative value, it doesn't allow any bidirectional streams.
	// Values larger than 2^60 will be clipped to that value.
	MaxIncomingStreams int64
	// MaxIncomingUniStreams is the maximum number of concurrent unidirectional streams that a peer is allowed to open.
	// If not set, it will default to 100.
	// If set to a negative value, it doesn't allow any unidirectional streams.
	// Values larger than 2^60 will be clipped to that value.
	MaxIncomingUniStreams int64
	// KeepAlivePeriod defines whether this peer will periodically send a packet to keep the connection alive.
	// If set to 0, then no keep alive is sent. Otherwise, the keep alive is sent on that period (or at most
	// every half of MaxIdleTimeout, whichever is smaller).
	KeepAlivePeriod time.Duration
	// DisablePathMTUDiscovery disables Path MTU Discovery (RFC 8899).
	// This allows the sending of QUIC packets that fully utilize the available MTU of the path.
	// Path MTU discovery is only available on systems that allow setting of the Don't Fragment (DF) bit.
	// If unavailable or disabled, packets will be at most 1252 (IPv4) / 1232 (IPv6) bytes in size.
	DisablePathMTUDiscovery bool
	// Allow0RTT allows the application to decide if a 0-RTT connection attempt should be accepted.
	// Only valid for the server.
	Allow0RTT bool
	// Enable QUIC datagram support (RFC 9221).
	EnableDatagrams bool
	Tracer          func(context.Context, logging.Perspective, ConnectionID) *logging.ConnectionTracer
}

type ClientHelloInfo struct {
	RemoteAddr net.Addr
}
