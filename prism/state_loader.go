package prism

type LoaderState string

const (
	LoaderStateIdle    LoaderState = "idle"
	LoaderStateLoading LoaderState = "loading"
	LoaderStateError   LoaderState = "error"
	LoaderStateReady   LoaderState = "ready"
)

type LoaderStateMachine struct {
	State LoaderState
	Error error
}

func NewLoaderState() *LoaderStateMachine {
	return &LoaderStateMachine{
		State: LoaderStateIdle,
	}
}

func (lsm *LoaderStateMachine) IsIdle() bool    { return lsm.State == LoaderStateIdle }
func (lsm *LoaderStateMachine) IsLoading() bool { return lsm.State == LoaderStateLoading }
func (lsm *LoaderStateMachine) IsError() bool   { return lsm.State == LoaderStateError }
func (lsm *LoaderStateMachine) IsReady() bool   { return lsm.State == LoaderStateReady }

func (lsm *LoaderStateMachine) SetIdle() *LoaderStateMachine {
	lsm.State = LoaderStateIdle
	lsm.Error = nil
	return lsm
}

func (lsm *LoaderStateMachine) SetLoading() *LoaderStateMachine {
	lsm.State = LoaderStateLoading
	lsm.Error = nil
	return lsm
}

func (lsm *LoaderStateMachine) SetError(err error) *LoaderStateMachine {
	lsm.State = LoaderStateError
	lsm.Error = err
	return lsm
}

func (lsm *LoaderStateMachine) SetReady() *LoaderStateMachine {
	lsm.State = LoaderStateReady
	lsm.Error = nil
	return lsm
}

func (lsm *LoaderStateMachine) GetError() error {
	return lsm.Error
}

func (lsm *LoaderStateMachine) GetErrorMessage() string {
	if lsm.Error != nil {
		return lsm.Error.Error()
	}
	return ""
}