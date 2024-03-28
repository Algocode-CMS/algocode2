package priority

type Embed struct {
	Priority Holder
}

func (e *Embed) GetPriority() Holder {
	return e.Priority
}

func (e *Embed) SetPriority(p int64) {
	e.Priority.Priority = p
}
