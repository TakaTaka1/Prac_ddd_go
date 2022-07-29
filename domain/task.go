package repository

import (
	"time"
)

type task struct {
	id	string
	taskStatus	TaskStatus
	name	string
	dueDate time.Time
	priority	PriorityStatus
	postPoneStatus int64
}



const POSTPONE_MAX_COUNT = 5

type TaskStatus string

const (
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone	TaskStatus = "done"
)

type PriorityStatus string

const (
	PriorityStatusHigh PriorityStatus = "high"
	PriorityStatusMiddle PriorityStatus = "middle"
	PriorityStatusLow	PriorityStatus = "low"
)

func NewTask(name string , dueDate time.Time, priority PriorityStatus) (*task, error) {
	if name == "" || dueDate.IsZero() {
		return nil, errors.New("必須項目が設定されていない")
	}
	
	return &task{
		taskStatus : TaskStatusDoing,
		name:	name,
		dueDate:	dueDate,
		priority:	priority,
		postponeCount:	0,
	},nil
}

func (t *task) Postpone() (*task, error) {
	if !t.CanPostpone() {
		return nil, errors.New("最大延長回数を超過しています。")
	}
	t.dueDate.Add(24 * time.Hour)
	t.postponeCount++
	return t, nil
}

func (t *task) Done() {
	t.taskStatus = TaskStatusDone
}

// name,priorityのsetterは存在しないので、name,priorityを変更することはできない

// getter
func (t *task) GetID() string {
	return t.id
}
func (t *task) GetName() string {
	return t.name
}
func (t *task) GetDueDate() time.Time {
	return t.dueDate
}
func (t *task) GetDueDate() PriorityStatus {
	return t.priority
}

func (t *task) IsDoing() bool {
	return t.taskStatus == TaskStatusDoing
}
func (t *task) CanPostpone() bool {
	return t.postponeCount < POSTPONE_MAX_COUNT
}

// getter
func (t *task) GetID() string {
	return t.id
}
func (t *task) GetName() string {
	return t.name
}
func (t *task) GetDueDate() time.Time {
	return t.dueDate
}
func (t *task) GetDueDate() PriorityStatus {
	return t.priority
}

func (t *task) IsDoing() bool {
	return t.taskStatus == TaskStatusDoing
}
func (t *task) CanPostpone() bool {
	return t.postponeCount < POSTPONE_MAX_COUNT
}