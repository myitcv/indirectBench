package main

import (
	"testing"
	"unsafe"
)

const (
	Large = 1000000
)

func BenchmarkRawSpecific(b *testing.B) {
	t1 := &T1{
		s1: &S{},
		s2: &S{},
		s3: &S{},
		s4: &S{},
		s5: &S{},
	}
	for n := 0; n < b.N; n++ {
		_ = t1.s1
		_ = t1.s2
		_ = t1.s3
		_ = t1.s4
		_ = t1.s5
	}
}

func BenchmarkPointersSpecific(b *testing.B) {
	t1 := &T1{
		s1: &S{},
		s2: &S{},
		s3: &S{},
		s4: &S{},
		s5: &S{},
	}
	for n := 0; n < b.N; n++ {
		_ = t1.S1()
		_ = t1.S2()
		_ = t1.S3()
		_ = t1.S4()
		_ = t1.S5()
	}
}

func BenchmarkIndirectSpecific(b *testing.B) {
	lookup := []unsafe.Pointer{
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
	}
	t2 := &T2{
		s1: 0,
		s2: 1,
		s3: 2,
		s4: 3,
		s5: 4,
	}
	for n := 0; n < b.N; n++ {
		_ = t2.S1(lookup)
		_ = t2.S2(lookup)
		_ = t2.S3(lookup)
		_ = t2.S4(lookup)
		_ = t2.S5(lookup)
	}
}

func BenchmarkGoRoutineSpecific(b *testing.B) {
	lookup := []unsafe.Pointer{
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
		unsafe.Pointer(&S{}),
	}
	Set(lookup)
	t3 := &T3{
		s1: 0,
		s2: 1,
		s3: 2,
		s4: 3,
		s5: 4,
	}
	for n := 0; n < b.N; n++ {
		_ = t3.S1()
		_ = t3.S2()
		_ = t3.S3()
		_ = t3.S4()
		_ = t3.S5()
	}
}

func BenchmarkIndirectMapSpecific(b *testing.B) {
	lookup := map[uint]unsafe.Pointer{
		0: unsafe.Pointer(&S{}),
		1: unsafe.Pointer(&S{}),
		2: unsafe.Pointer(&S{}),
		3: unsafe.Pointer(&S{}),
		4: unsafe.Pointer(&S{}),
	}
	t4 := &T4{
		s1: 0,
		s2: 1,
		s3: 2,
		s4: 3,
		s5: 4,
	}
	for n := 0; n < b.N; n++ {
		_ = t4.S1(lookup)
		_ = t4.S2(lookup)
		_ = t4.S3(lookup)
		_ = t4.S4(lookup)
		_ = t4.S5(lookup)
	}
}

func BenchmarkPointersMultiple(b *testing.B) {
	t1 := &T1{
		ss: make([]*S, Large),
	}

	for i := 0; i < Large; i++ {
		t1.ss[i] = &S{}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < Large; i++ {
			_ = t1.S(uint(i))
		}
	}
}

func BenchmarkIndirectMultiple(b *testing.B) {
	t2 := &T2{
		ss: make([]uint, Large),
	}

	context := make([]unsafe.Pointer, Large)

	for i := 0; i < Large; i++ {
		context[uint(i)] = unsafe.Pointer(&S{})
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < Large; i++ {
			_ = t2.S(context, uint(i))
		}
	}
}

func BenchmarkGoRoutineMultiple(b *testing.B) {
	t3 := &T3{
		ss: make([]uint, Large),
	}

	context := make([]unsafe.Pointer, Large)

	for i := 0; i < Large; i++ {
		context[uint(i)] = unsafe.Pointer(&S{})
	}

	Set(context)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < Large; i++ {
			_ = t3.S(uint(i))
		}
	}
}

func BenchmarkIndirectMapMultiple(b *testing.B) {
	t4 := &T4{
		ss: make([]uint, Large),
	}

	context := make(map[uint]unsafe.Pointer, Large)

	for i := 0; i < Large; i++ {
		context[uint(i)] = unsafe.Pointer(&S{})
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < Large; i++ {
			_ = t4.S(context, uint(i))
		}
	}
}
