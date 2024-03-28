package priority

import (
	"cmp"
	"slices"
)

type Priority interface {
	GetPriority() Holder
	SetPriority(p int64)
}

func Sort[T Priority](elems []T) {
	slices.SortFunc(elems, func(a, b T) int {
		return cmp.Compare(a.GetPriority().Priority, b.GetPriority().Priority)
	})
}

func Reorder[T Priority](elems []T) {
	slices.SortFunc(elems, func(a, b T) int {
		ap := a.GetPriority()
		bp := b.GetPriority()
		if ap.Priority != bp.Priority {
			return cmp.Compare(ap.Priority, bp.Priority)
		}
		return cmp.Compare(ap.OldPriority, bp.OldPriority)
	})

	for i, e := range elems {
		e.SetPriority(int64(i))
	}
}
