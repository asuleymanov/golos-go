package operations

//WorkerStateOperation represents worker_state operation data.
type WorkerStateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	State    int    `json:"state"`
}

//Type function that defines the type of operation WorkerStateOperation.
func (op *WorkerStateOperation) Type() OpType {
	return TypeWorkerState
}

//Data returns the operation data WorkerStateOperation.
func (op *WorkerStateOperation) Data() interface{} {
	return op
}
