package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running Raft.
// Shardctrler decides which group serves each shard.
// Shardctrler may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK                 = "OK"
	ERR_NoKey          = "Not such key"
	ERR_WrongGroup     = "Wrong group"
	ERR_WrongLeader    = "Not leader"
	ERR_CommitTimeout  = "Commit timeout"
	ERR_FailedToCommit = "Failed to commit"
)

const (
	OT_GET    = "Get"
	OT_PUT    = "Put"
	OT_APPEND = "Append"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	// You'll have to add definitions here.
	Key   string
	Value string
	Op    string // "Put" or "Append"

	Id     int64
	ReqNum int64
}

type PutAppendReply struct {
	Err Err
}

type GetArgs struct {
	Key    string
	Id     int64
	ReqNum int64
}

type GetReply struct {
	Err   Err
	Value string
}

type GenericArgs interface {
	getId() int64
	getReqNum() int64
	getKey() string
}

type GenericReply interface {
	setErr(str Err)
}

func (args GetArgs) getId() int64 {
	return args.Id
}

func (args GetArgs) getReqNum() int64 {
	return args.ReqNum
}

func (args GetArgs) getKey() string {
	return args.Key
}

func (args PutAppendArgs) getId() int64 {
	return args.Id
}

func (args PutAppendArgs) getReqNum() int64 {
	return args.ReqNum
}

func (args PutAppendArgs) getKey() string {
	return args.Key
}

func (reply *GetReply) setErr(err Err) {
	reply.Err = err
}

func (reply *PutAppendReply) setErr(err Err) {
	reply.Err = err
}
