package parallel

import (
	"sync"
)

type ErrOut struct {
	In  any
	Err error
}

func DoParallel[In any, Out any](ins []In, exec func(In) ([]Out, error)) ([]Out, []ErrOut) {
	if len(ins) == 0 {
		return nil, nil
	}
	group := sync.WaitGroup{}
	group.Add(len(ins))

	s := make([]Out, 0, len(ins))
	sLock := sync.Mutex{}

	var e []ErrOut
	eLock := sync.Mutex{}

	for _, in := range ins {
		go func(i In) {
			defer group.Done()
			out, err := exec(i)
			if err != nil {
				eLock.Lock()
				e = append(e, ErrOut{i, err})
				eLock.Unlock()
			}
			sLock.Lock()
			s = append(s, out...)
			sLock.Unlock()
		}(in)
	}

	group.Wait()

	return s, e
}
