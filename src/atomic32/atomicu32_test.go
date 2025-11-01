package atomic32_test

/*
---------------------------------------
This file has been written by GPT.
---------------------------------------
*/

import (
	"runtime"
	"sync"
	"testing"
	"withoutforget/go_atomic/src/atomic32"
)

func TestBasicStoreLoadU32(t *testing.T) {
	var a atomic32.AtomicU32

	testCases := []struct {
		name  string
		value uint32
	}{
		{"zero", 0},
		{"positive", 42},
		{"large", 4294967295},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a.Store(tc.value)
			if got := a.Load(); got != tc.value {
				t.Errorf("Store(%d) then Load() = %d, want %d", tc.value, got, tc.value)
			}
		})
	}
}

func TestMemoryOrderingU32(t *testing.T) {
	var a atomic32.AtomicU32

	a.StoreRelaxed(100)
	if got := a.LoadRelaxed(); got != 100 {
		t.Errorf("Relaxed: got %d, want 100", got)
	}

	a.StoreRelease(200)
	if got := a.LoadAcquire(); got != 200 {
		t.Errorf("Acquire/Release: got %d, want 200", got)
	}

	a.StoreSeqCst(300)
	if got := a.LoadSeqCst(); got != 300 {
		t.Errorf("SeqCst: got %d, want 300", got)
	}
}

func TestAddU32(t *testing.T) {
	testOpU32(t, "Add", func(a *atomic32.AtomicU32, v uint32) uint32 { return a.Add(v) }, 10, 5, 10, 15)
	testOpU32(t, "Add wrap", func(a *atomic32.AtomicU32, v uint32) uint32 { return a.Add(v) }, 4294967290, 10, 4294967290, 4)
}

func TestAddMemoryOrderingU32(t *testing.T) {
	testOpOrderingU32(t, "AddRelaxed", (*atomic32.AtomicU32).AddRelaxed, 100, 50, 100, 150)
	testOpOrderingU32(t, "AddAcquire", (*atomic32.AtomicU32).AddAcquire, 100, 50, 100, 150)
	testOpOrderingU32(t, "AddRelease", (*atomic32.AtomicU32).AddRelease, 100, 50, 100, 150)
	testOpOrderingU32(t, "AddAcqRel", (*atomic32.AtomicU32).AddAcqRel, 100, 50, 100, 150)
	testOpOrderingU32(t, "AddSeqCst", (*atomic32.AtomicU32).AddSeqCst, 100, 50, 100, 150)
}

func TestOrU32(t *testing.T) {
	testOpU32(t, "Or", func(a *atomic32.AtomicU32, v uint32) uint32 { return a.Or(v) }, 0b1010, 0b0101, 10, 0b1111)
}

func TestOrMemoryOrderingU32(t *testing.T) {
	testOpOrderingU32(t, "OrRelaxed", (*atomic32.AtomicU32).OrRelaxed, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU32(t, "OrAcquire", (*atomic32.AtomicU32).OrAcquire, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU32(t, "OrRelease", (*atomic32.AtomicU32).OrRelease, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU32(t, "OrAcqRel", (*atomic32.AtomicU32).OrAcqRel, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU32(t, "OrSeqCst", (*atomic32.AtomicU32).OrSeqCst, 0b1100, 0b0011, 0b1100, 0b1111)
}

func TestAndU32(t *testing.T) {
	testOpU32(t, "And", func(a *atomic32.AtomicU32, v uint32) uint32 { return a.And(v) }, 0b1111, 0b1010, 15, 0b1010)
}

func TestAndMemoryOrderingU32(t *testing.T) {
	testOpOrderingU32(t, "AndRelaxed", (*atomic32.AtomicU32).AndRelaxed, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU32(t, "AndAcquire", (*atomic32.AtomicU32).AndAcquire, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU32(t, "AndRelease", (*atomic32.AtomicU32).AndRelease, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU32(t, "AndAcqRel", (*atomic32.AtomicU32).AndAcqRel, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU32(t, "AndSeqCst", (*atomic32.AtomicU32).AndSeqCst, 0b1111, 0b0101, 0b1111, 0b0101)
}

func TestCASU32(t *testing.T) {
	var a atomic32.AtomicU32
	a.Store(42)

	if !a.CAS(42, 100) {
		t.Error("CAS(42, 100) should succeed")
	}
	if got := a.Load(); got != 100 {
		t.Errorf("After successful CAS: got %d, want 100", got)
	}

	if a.CAS(42, 200) {
		t.Error("CAS(42, 200) should fail")
	}
	if got := a.Load(); got != 100 {
		t.Errorf("After failed CAS: got %d, want 100", got)
	}
}

func TestConcurrentAddU32(t *testing.T) {
	var a atomic32.AtomicU32
	testConcurrentU32(t, &a, 100, 1000, func(a *atomic32.AtomicU32) { a.Add(1) }, 100000)
}

func TestConcurrentCASU32(t *testing.T) {
	var a atomic32.AtomicU32
	a.Store(0)

	const goroutines = 100
	var successCounter atomic32.AtomicU32

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				old := a.Load()
				if old >= 1000 {
					break
				}
				if a.CAS(old, old+1) {
					successCounter.Add(1)
				}
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()

	if got := a.Load(); got != 1000 {
		t.Errorf("Final value: got %d, want at least 1000", got)
	}

	successCount := successCounter.Load()
	if successCount < 1 {
		t.Error("No successful CAS operations")
	}
}

func TestConcurrentMixedU32(t *testing.T) {
	var a atomic32.AtomicU32
	a.Store(0)

	const goroutines = 50
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines * 4)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				a.Add(1)
			}
		}()
	}

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				a.Or(0b0001)
			}
		}()
	}

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				a.And(0xFFFFFFFF)
			}
		}()
	}

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				val := a.Load()
				a.Store(val)
			}
		}()
	}

	wg.Wait()

	final := a.Load()
	t.Logf("Final value after mixed operations: %d", final)
}

func TestFunctionPointersU32(t *testing.T) {
	var a atomic32.AtomicU32

	atomic32.AtomicStoreU32(&a, 42)
	if got := atomic32.AtomicLoadU32(&a); got != 42 {
		t.Errorf("Function call: got %d, want 42", got)
	}

	old := atomic32.AtomicAddU32(&a, 8)
	if old != 42 {
		t.Errorf("AtomicAddU32 returned %d, want 42", old)
	}
	if got := atomic32.AtomicLoadU32(&a); got != 50 {
		t.Errorf("After AtomicAddU32: got %d, want 50", got)
	}
}

func TestOverflowU32(t *testing.T) {
	var a atomic32.AtomicU32
	a.Store(4294967295)

	a.Add(1)
	if got := a.Load(); got != 0 {
		t.Errorf("Overflow: got %d, want 0", got)
	}
}

func TestZeroValueU32(t *testing.T) {
	var a atomic32.AtomicU32

	if got := a.Load(); got != 0 {
		t.Errorf("Zero value Load: got %d, want 0", got)
	}

	a.Add(5)
	if got := a.Load(); got != 5 {
		t.Errorf("Zero value after Add: got %d, want 5", got)
	}
}

func testOpU32(t *testing.T, name string, op func(*atomic32.AtomicU32, uint32) uint32, initial, delta, wantOld, wantNew uint32) {
	t.Run(name, func(t *testing.T) {
		var a atomic32.AtomicU32
		a.Store(initial)
		old := op(&a, delta)
		if old != wantOld {
			t.Errorf("returned %d, want %d", old, wantOld)
		}
		if got := a.Load(); got != wantNew {
			t.Errorf("result %d, want %d", got, wantNew)
		}
	})
}

func testOpOrderingU32(t *testing.T, name string, fn func(*atomic32.AtomicU32, uint32) uint32, initial, delta, wantOld, wantNew uint32) {
	t.Run(name, func(t *testing.T) {
		var a atomic32.AtomicU32
		a.Store(initial)
		old := fn(&a, delta)
		if old != wantOld {
			t.Errorf("%s: returned %d, want %d", name, old, wantOld)
		}
		if got := a.Load(); got != wantNew {
			t.Errorf("%s: result %d, want %d", name, got, wantNew)
		}
	})
}

func testConcurrentU32(t *testing.T, a *atomic32.AtomicU32, goroutines, iterations int, op func(*atomic32.AtomicU32), expected uint32) {
	a.Store(0)
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				op(a)
			}
		}()
	}

	wg.Wait()

	if got := a.Load(); got != expected {
		t.Errorf("Concurrent: got %d, want %d", got, expected)
	}
}

func BenchmarkStoreMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreRelaxed(uint32(i))
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreRelease(uint32(i))
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreSeqCst(uint32(i))
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.Store(uint32(i))
		}
	})
}

func BenchmarkLoadMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	a.Store(42)
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a.LoadRelaxed()
		}
	})
	b.Run("Acquire", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a.LoadAcquire()
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a.LoadSeqCst()
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a.Load()
		}
	})
}

func BenchmarkAddMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AddRelaxed(1)
		}
	})
	b.Run("Acquire", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AddAcquire(1)
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AddRelease(1)
		}
	})
	b.Run("AcqRel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AddAcqRel(1)
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AddSeqCst(1)
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.Add(1)
		}
	})
}

func BenchmarkOrMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.OrRelaxed(0b0001)
		}
	})
	b.Run("Acquire", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.OrAcquire(0b0001)
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.OrRelease(0b0001)
		}
	})
	b.Run("AcqRel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.OrAcqRel(0b0001)
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.OrSeqCst(0b0001)
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.Or(0b0001)
		}
	})
}

func BenchmarkAndMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndRelaxed(0xFFFFFFFF)
		}
	})
	b.Run("Acquire", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndAcquire(0xFFFFFFFF)
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndRelease(0xFFFFFFFF)
		}
	})
	b.Run("AcqRel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndAcqRel(0xFFFFFFFF)
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndSeqCst(0xFFFFFFFF)
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.And(0xFFFFFFFF)
		}
	})
}

func BenchmarkConcurrentStoreMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.StoreRelaxed(42)
			}
		})
	})
	b.Run("Release", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.StoreRelease(42)
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.StoreSeqCst(42)
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.Store(42)
			}
		})
	})
}

func BenchmarkConcurrentLoadMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	a.Store(42)
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = a.LoadRelaxed()
			}
		})
	})
	b.Run("Acquire", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = a.LoadAcquire()
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = a.LoadSeqCst()
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = a.Load()
			}
		})
	})
}

func BenchmarkConcurrentAddMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AddRelaxed(1)
			}
		})
	})
	b.Run("Acquire", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AddAcquire(1)
			}
		})
	})
	b.Run("Release", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AddRelease(1)
			}
		})
	})
	b.Run("AcqRel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AddAcqRel(1)
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AddSeqCst(1)
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.Add(1)
			}
		})
	})
}

func BenchmarkConcurrentOrMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.OrRelaxed(0b0001)
			}
		})
	})
	b.Run("Acquire", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.OrAcquire(0b0001)
			}
		})
	})
	b.Run("Release", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.OrRelease(0b0001)
			}
		})
	})
	b.Run("AcqRel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.OrAcqRel(0b0001)
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.OrSeqCst(0b0001)
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.Or(0b0001)
			}
		})
	})
}

func BenchmarkConcurrentAndMemoryOrderingU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndRelaxed(0xFFFFFFFF)
			}
		})
	})
	b.Run("Acquire", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndAcquire(0xFFFFFFFF)
			}
		})
	})
	b.Run("Release", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndRelease(0xFFFFFFFF)
			}
		})
	})
	b.Run("AcqRel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndAcqRel(0xFFFFFFFF)
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndSeqCst(0xFFFFFFFF)
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.And(0xFFFFFFFF)
			}
		})
	})
}

func BenchmarkCASU32(b *testing.B) {
	var a atomic32.AtomicU32
	for i := 0; i < b.N; i++ {
		a.Store(uint32(i))
		a.CAS(uint32(i), uint32(i+1))
	}
}

func BenchmarkConcurrentCASU32(b *testing.B) {
	var a atomic32.AtomicU32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			old := a.Load()
			a.CAS(old, old+1)
		}
	})
}
