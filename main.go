package main

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"gopkg.in/qml.v1/cdata"
)

var uuid uint

type S struct {
	uid uint

	name string
}

type T1 struct {
	s1 *S
	s2 *S
	s3 *S
	s4 *S
	s5 *S

	ss []*S
}

func (t *T1) S1() *S {
	return t.s1
}

func (t *T1) S2() *S {
	return t.s2
}

func (t *T1) S3() *S {
	return t.s3
}

func (t *T1) S4() *S {
	return t.s4
}

func (t *T1) S5() *S {
	return t.s5
}

func (t *T1) S(i uint) *S {
	return t.ss[i]
}

type T2 struct {
	s1 uint
	s2 uint
	s3 uint
	s4 uint
	s5 uint

	ss []uint
}

func (t *T2) S1(context []unsafe.Pointer) *S {
	return (*S)(context[t.s1])
}

func (t *T2) S2(context []unsafe.Pointer) *S {
	return (*S)(context[t.s2])
}

func (t *T2) S3(context []unsafe.Pointer) *S {
	return (*S)(context[t.s3])
}

func (t *T2) S4(context []unsafe.Pointer) *S {
	return (*S)(context[t.s4])
}

func (t *T2) S5(context []unsafe.Pointer) *S {
	return (*S)(context[t.s5])
}

func (t *T2) S(context []unsafe.Pointer, i uint) *S {
	return (*S)(context[t.ss[i]])
}

var goRoutineMap atomic.Value
var mu sync.Mutex

type Map map[uintptr][]unsafe.Pointer

func Get() []unsafe.Pointer {
	m := goRoutineMap.Load().(Map)
	gr := cdata.Ref()

	return m[gr]
}

func Set(v []unsafe.Pointer) {
	gr := cdata.Ref()
	mu.Lock()

	var mc Map
	m := goRoutineMap.Load()

	if m != nil {
		m := m.(Map)

		mc = make(Map, len(m))

		for k, v := range m {
			mc[k] = v
		}
	} else {
		mc = make(Map)
	}

	mc[gr] = v

	goRoutineMap.Store(mc)

	mu.Unlock()
}

type T3 struct {
	s1 uint
	s2 uint
	s3 uint
	s4 uint
	s5 uint

	ss []uint
}

func (t *T3) S1() *S {
	context := Get()
	return (*S)(context[t.s1])
}

func (t *T3) S2() *S {
	context := Get()
	return (*S)(context[t.s2])
}

func (t *T3) S3() *S {
	context := Get()
	return (*S)(context[t.s3])
}

func (t *T3) S4() *S {
	context := Get()
	return (*S)(context[t.s4])
}

func (t *T3) S5() *S {
	context := Get()
	return (*S)(context[t.s5])
}

func (t *T3) S(i uint) *S {
	context := Get()
	return (*S)(context[t.ss[i]])
}

func main() {
}
