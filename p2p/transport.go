package p2p

// peer denotes the remote node
type Peer interface {
}

// to handle communication between two nodes
type Transport interface {
	ListenAndAccept() error
}
