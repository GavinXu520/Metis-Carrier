package types

import "github.com/RosettaFlow/Carrier-Go/common"

type ProposalTask struct {
	ProposalId common.Hash
	*ScheduleTask
	CreateAt uint64
}

type ConsensusTaskWrap struct {
	Task     *ScheduleTask
	ResultCh chan<- *TaskConsResult
}
type ScheduleTaskWrap struct {
	Task     *ScheduleTask
	ResultCh chan<- *ScheduleResult
}

type ScheduleTask struct {
	TaskId                string                        `json:"TaskId"`
	TaskName              string                        `json:"taskName"`
	Owner                 *ScheduleTaskDataSupplier     `json:"owner"`
	Partners              []*ScheduleTaskDataSupplier   `json:"partners"`
	PowerSuppliers        []*ScheduleTaskPowerSupplier  `json:"powerSuppliers"`
	Receivers             []*ScheduleTaskResultReceiver `json:"receivers"`
	CalculateContractCode string                        `json:"calculateContractCode"`
	DataSplitContractCode string                        `json:"dataSplitContractCode"`
	OperationCost         *TaskOperationCost            `json:"spend"`
	CreateAt              uint64                        `json:"createAt"`
}

type ScheduleTaskDataSupplier struct {
	*NodeAlias
	MetaData *SupplierMetaData `json:"metaData"`
}
type ScheduleTaskPowerSupplier struct {
	*NodeAlias
}

type ScheduleSupplierMetaData struct {
	MetaId          string   `json:"metaId"`
	ColumnIndexList []uint64 `json:"columnIndexList"`
}

type ScheduleTaskResultReceiver struct {
	*NodeAlias
	Providers []*NodeAlias `json:"providers"`
}

type TaskConsStatus uint16

const (
	TaskConsensusInterrupt TaskConsStatus = 0x0001
	TaskRunningInterrupt   TaskConsStatus = 0x0100
)

// Task consensus result
type TaskConsResult struct {
	TaskId string
	Status TaskConsStatus
	Done   bool
	Err    error
}

type TaskSchedStatus bool

const (
	TaskSchedOk     TaskSchedStatus = true
	TaskSchedFailed TaskSchedStatus = false
)

type ScheduleResult struct {
	TaskId string
	Status TaskSchedStatus
	Err    error
}

//type ScheduleNodeAlias struct {
//	Name         string `json:"name"`
//	NodeId       string `json:"nodeId"`
//	IdentityId   string `json:"identityId"`
//	PeerIp   string `json:"peerIp"`
//	PeerPort string `json:"peerPort"`
//}
