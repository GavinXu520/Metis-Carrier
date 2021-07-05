package types

import (
	"bytes"
	"github.com/RosettaFlow/Carrier-Go/common"
	pb "github.com/RosettaFlow/Carrier-Go/lib/api"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"io"
	"sync/atomic"
)

type Task struct {
	data *libTypes.TaskData

	// caches
	hash atomic.Value
	size atomic.Value
}

func NewTask(data *libTypes.TaskData) *Task {
	return &Task{data: data}
}

func (m *Task) EncodePb(w io.Writer) error {
	if m.data == nil {
		m.data = new(libTypes.TaskData)
	}
	data, err := m.data.Marshal()
	if err == nil {
		w.Write(data)
	}
	return err
}

func (m *Task) DecodePb(data []byte) error {
	m.size.Store(common.StorageSize(len(data)))
	return m.data.Unmarshal(data)
}

func (m *Task) Hash() common.Hash {
	if hash := m.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	buffer := new(bytes.Buffer)
	m.EncodePb(buffer)
	v := protoBufHash(buffer.Bytes())
	m.hash.Store(v)
	return v
}

func (m *Task) TaskId() string {
	return m.data.TaskId
}

func (m *Task) TaskData() *libTypes.TaskData {
	return m.data
}

// TaskDataArray is a Transaction slice type for basic sorting.
type TaskDataArray []*Task

// Len returns the length of s.
func (s TaskDataArray) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s TaskDataArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s TaskDataArray) GetPb(i int) []byte {
	buffer := new(bytes.Buffer)
	s[i].EncodePb(buffer)
	return buffer.Bytes()
}

func NewTaskDataArray(metaData []*libTypes.TaskData) TaskDataArray {
	var s TaskDataArray
	for _, v := range metaData {
		s = append(s, NewTask(v))
	}
	return s
}

func (s TaskDataArray) To() []*libTypes.TaskData {
	arr := make([]*libTypes.TaskData, 0, s.Len())
	for _, v := range s {
		arr = append(arr, v.data)
	}
	return arr
}

type TaskDetailShow struct {
	TaskId        string                   `json:"taskId"`
	TaskName      string                   `json:"taskName"`
	Owner         *NodeAlias               `json:"owner"`
	AlgoSupplier  *NodeAlias               `json:"algoSupplier"`
	DataSupplier  []*TaskDataSupplierShow  `json:"dataSupplier"`
	PowerSupplier []*TaskPowerSupplierShow `json:"powerSupplier"`
	Receivers     []*NodeAlias             `json:"receivers"`
	CreateAt      uint64                   `json:"createat"`
	StartAt       uint64                   `json:"startAt"`
	EndAt         uint64                   `json:"endAt"`
	State         string                   `json:"state"`
	OperationCost *TaskOperationCost       `json:"operationCost"`
}

func ConvertTaskDetailShowToPB(task *TaskDetailShow) *pb.TaskDetailShow {
	return &pb.TaskDetailShow{
		TaskId:        task.TaskId,
		TaskName:      task.TaskName,
		Owner:         ConvertNodeAliasToPB(task.Owner),
		AlgoSupplier:  ConvertNodeAliasToPB(task.AlgoSupplier),
		DataSupplier:  ConvertTaskDataSupplierShowArrToPB(task.DataSupplier),
		PowerSupplier: ConvertTaskPowerSupplierShowArrToPB(task.PowerSupplier),
		Receivers:     ConvertNodeAliasArrToPB(task.Receivers),
		CreateAt:      task.CreateAt,
		StartAt:       task.StartAt,
		EndAt:         task.EndAt,
		State:         task.State,
		OperationCost: ConvertTaskOperationCostToPB(task.OperationCost),
	}
}
func ConvertTaskDetailShowFromPB(task *pb.TaskDetailShow) *TaskDetailShow {
	return &TaskDetailShow{
		TaskId:        task.TaskId,
		TaskName:      task.TaskName,
		Owner:         ConvertNodeAliasFromPB(task.Owner),
		AlgoSupplier:  ConvertNodeAliasFromPB(task.AlgoSupplier),
		DataSupplier:  ConvertTaskDataSupplierShowArrFromPB(task.DataSupplier),
		PowerSupplier: ConvertTaskPowerSupplierShowArrFromPB(task.PowerSupplier),
		Receivers:     ConvertNodeAliasArrFromPB(task.Receivers),
		CreateAt:      task.CreateAt,
		StartAt:       task.StartAt,
		EndAt:         task.EndAt,
		State:         task.State,
		OperationCost: ConvertTaskOperationCostFromPB(task.OperationCost),
	}
}

type TaskDataSupplierShow struct {
	MemberInfo   *NodeAlias `json:"memberInfo"`
	MetaDataId   string     `json:"metaId"`
	MetaDataName string     `json:"metaName"`
}

func ConvertTaskDataSupplierShowToPB(dataSupplier *TaskDataSupplierShow) *pb.TaskDataSupplierShow {
	return &pb.TaskDataSupplierShow{
		MemberInfo:   ConvertNodeAliasToPB(dataSupplier.MemberInfo),
		MetaDataId:   dataSupplier.MetaDataId,
		MetaDataName: dataSupplier.MetaDataName,
	}
}
func ConvertTaskDataSupplierShowFromPB(dataSupplier *pb.TaskDataSupplierShow) *TaskDataSupplierShow {
	return &TaskDataSupplierShow{
		MemberInfo:   ConvertNodeAliasFromPB(dataSupplier.MemberInfo),
		MetaDataId:   dataSupplier.MetaDataId,
		MetaDataName: dataSupplier.MetaDataName,
	}
}
func ConvertTaskDataSupplierShowArrToPB(dataSuppliers []*TaskDataSupplierShow) []*pb.TaskDataSupplierShow {
	arr := make([]*pb.TaskDataSupplierShow, len(dataSuppliers))
	for i, dataSupplier := range dataSuppliers {

		supplier := &pb.TaskDataSupplierShow{
			MemberInfo:   ConvertNodeAliasToPB(dataSupplier.MemberInfo),
			MetaDataId:   dataSupplier.MetaDataId,
			MetaDataName: dataSupplier.MetaDataName,
		}
		arr[i] = supplier
	}

	return arr
}
func ConvertTaskDataSupplierShowArrFromPB(dataSuppliers []*pb.TaskDataSupplierShow) []*TaskDataSupplierShow {
	arr := make([]*TaskDataSupplierShow, len(dataSuppliers))
	for i, dataSupplier := range dataSuppliers {

		supplier := &TaskDataSupplierShow{
			MemberInfo:   ConvertNodeAliasFromPB(dataSupplier.MemberInfo),
			MetaDataId:   dataSupplier.MetaDataId,
			MetaDataName: dataSupplier.MetaDataName,
		}
		arr[i] = supplier
	}

	return arr
}

type TaskPowerSupplierShow struct {
	MemberInfo    *NodeAlias     `json:"memberInfo"`
	ResourceUsage *ResourceUsage `json:"resourceUsage"`
}

func ConvertTaskPowerSupplierShowShowToPB(powerSupplier *TaskPowerSupplierShow) *pb.TaskPowerSupplierShow {
	return &pb.TaskPowerSupplierShow{
		MemberInfo: ConvertNodeAliasToPB(powerSupplier.MemberInfo),
		PowerInfo:  ConvertResourceUsageToPB(powerSupplier.ResourceUsage),
	}
}
func ConvertTaskPowerSupplierShowFromPB(powerSupplier *pb.TaskPowerSupplierShow) *TaskPowerSupplierShow {
	return &TaskPowerSupplierShow{
		MemberInfo:    ConvertNodeAliasFromPB(powerSupplier.MemberInfo),
		ResourceUsage: ConvertResourceUsageFromPB(powerSupplier.PowerInfo),
	}
}
func ConvertTaskPowerSupplierShowArrToPB(powerSuppliers []*TaskPowerSupplierShow) []*pb.TaskPowerSupplierShow {
	arr := make([]*pb.TaskPowerSupplierShow, len(powerSuppliers))
	for i, powerSupplier := range powerSuppliers {

		supplier := &pb.TaskPowerSupplierShow{
			MemberInfo: ConvertNodeAliasToPB(powerSupplier.MemberInfo),
			PowerInfo:  ConvertResourceUsageToPB(powerSupplier.ResourceUsage),
		}
		arr[i] = supplier
	}

	return arr
}
func ConvertTaskPowerSupplierShowArrFromPB(powerSuppliers []*pb.TaskPowerSupplierShow) []*TaskPowerSupplierShow {
	arr := make([]*TaskPowerSupplierShow, len(powerSuppliers))
	for i, powerSupplier := range powerSuppliers {

		supplier := &TaskPowerSupplierShow{
			MemberInfo:    ConvertNodeAliasFromPB(powerSupplier.MemberInfo),
			ResourceUsage: ConvertResourceUsageFromPB(powerSupplier.PowerInfo),
		}
		arr[i] = supplier
	}

	return arr
}
