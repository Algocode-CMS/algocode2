package priority

import (
	"database/sql/driver"
	"errors"
)

type Holder struct {
	Priority    int64
	OldPriority int64
}

var ErrWrongPriority = errors.New("wrong type for priority field")

func (p *Holder) Scan(src any) error {
	if src == nil {
		p.Priority = 0
		p.OldPriority = 0
		return nil
	}

	priority, ok := src.(int64)
	if ok {
		p.Priority = priority
		p.OldPriority = priority
		return nil
	}

	return ErrWrongPriority
}

func (p *Holder) Value() (driver.Value, error) {
	return p.Priority, nil
}
