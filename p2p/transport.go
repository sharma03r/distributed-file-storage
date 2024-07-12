package p2p

// peer denotes the remote node
type Peer interface {
	Close() error
}

// to handle communication between two nodes
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
