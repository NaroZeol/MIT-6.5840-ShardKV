package shardkv

import "6.5840/shardctrler"

//
// Sharded key/value server.
// Lots of replica groups, each running Raft.
// Shardctrler decides which group serves each shard.
// Shardctrler may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK                  = "OK"
	ERR_NoKey           = "Not such key"
	ERR_WrongGroup      = "Wrong group"
	ERR_WrongLeader     = "Not leader"
	ERR_CommitTimeout   = "Commit timeout"
	ERR_FailedToCommit  = "Failed to commit"
	ERR_HigherConfigNum = "Higer config number"

	ERR_LowerConfigNum = "Lower config number"
)

const (
	OT_GET            = "Get"
	OT_PUT            = "Put"
	OT_APPEND         = "Append"
	OT_ApplyMovement  = "ApplyMovement"
	OT_ChangeConfig   = "ChangeConfig"
	OT_DeleteShards   = "DeleteShards"
	OT_UpdateShardVec = "UpdateShardVec"
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

type ApplyMovementArgs struct {
	Id     int64
	ReqNum int64

	Mp       map[string]string
	Sessions map[string]Session
	Config   shardctrler.Config
	Shards   map[int]bool
}

type ApplyMovementReply struct {
	Err Err
}

type UpdateShardVecArgs struct {
	Id     int64
	ReqNum int64

	ConfigNum int
	NewVec    [shardctrler.NShards]bool
}

type UpdateShardVecReply struct {
	Err Err
}

type ChangeConfigArgs struct {
	Id     int64
	ReqNum int64

	OldNum int
	NewNum int
	Config shardctrler.Config
}

type ChangeConfigReply struct {
	Num int
	Err Err
}

type RequestMapAndSessionArgs struct {
	Gid       int
	Me        int
	Shards    map[int]bool
	ConfigNum int
}

type RequestMapAndSessionReply struct {
	Err      Err
	Mp       map[string]string
	Sessions map[string]Session
}

type DeleteShardsArgs struct {
	Id        int64
	ReqNum    int64
	Shards    map[int]bool
	ConfigNum int
}

type DeleteShardsReply struct {
	Err Err
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

func (args ApplyMovementArgs) getId() int64 {
	return args.Id
}

func (args ApplyMovementArgs) getReqNum() int64 {
	return args.ReqNum
}

// only define for satisfying interface
func (args ApplyMovementArgs) getKey() string {
	return ""
}

func (args DeleteShardsArgs) getId() int64 {
	return args.Id
}

func (args DeleteShardsArgs) getReqNum() int64 {
	return args.ReqNum
}

// only define for satisfying interface
func (args DeleteShardsArgs) getKey() string {
	return ""
}

func (args ChangeConfigArgs) getId() int64 {
	return args.Id
}

func (args ChangeConfigArgs) getReqNum() int64 {
	return args.ReqNum
}

// only define for satisfying interface
func (args ChangeConfigArgs) getKey() string {
	return ""
}

func (args UpdateShardVecArgs) getId() int64 {
	return args.Id
}

func (args UpdateShardVecArgs) getReqNum() int64 {
	return args.ReqNum
}

// only define for satisfying interface
func (args UpdateShardVecArgs) getKey() string {
	return ""
}

func (reply *GetReply) setErr(err Err) {
	reply.Err = err
}

func (reply *PutAppendReply) setErr(err Err) {
	reply.Err = err
}

func (reply *ApplyMovementReply) setErr(err Err) {
	reply.Err = err
}

func (reply *ChangeConfigReply) setErr(err Err) {
	reply.Err = err
}

func (reply *DeleteShardsReply) setErr(err Err) {
	reply.Err = err
}

func (reply *UpdateShardVecReply) setErr(err Err) {
	reply.Err = err
}
