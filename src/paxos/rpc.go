package paxos

const (
	OK          = "OK"
	ErrRejected = "ErrRejected"
)

type Err string

type PrepareArgs struct {
	Instance int
	Proposal int
}

type PrepareReply struct {
	Err      Err
	Instance int
	Proposal int
	Value    interface{}
}

type AcceptArgs struct {
	Instance int
	Proposal int
	Value    interface{}
}

type AcceptReply struct {
	Err Err
}

type DecidedArgs struct {
	Sender  int
	DoneIns int

	Instance int
	Value    interface{}
}

type DecidedReply struct {
}
