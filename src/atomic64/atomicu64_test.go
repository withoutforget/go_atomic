package atomic64_test

/*
---------------------------------------
This file has been written by GPT.
---------------------------------------
*/

import (
	"runtime"
	"sync"
	"testing"
	"withoutforget/go_atomic/src/atomic64"
)

func TestBasicStoreLoadU64(t *testing.T) {
	var a atomic64.AtomicU64

	testCases := []struct {
		name  string
		value uint64
	}{
		{"zero", 0},
		{"positive", 42},
		{"large", 18446744073709551615},
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

func TestMemoryOrderingU64(t *testing.T) {
	var a atomic64.AtomicU64

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

func TestAddU64(t *testing.T) {
	testOpU64(t, "Add", func(a *atomic64.AtomicU64, v uint64) uint64 { return a.Add(v) }, 10, 5, 10, 15)
	testOpU64(t, "Add wrap", func(a *atomic64.AtomicU64, v uint64) uint64 { return a.Add(v) }, 18446744073709551610, 10, 18446744073709551610, 4)
}

func TestAddMemoryOrderingU64(t *testing.T) {
	testOpOrderingU64(t, "AddRelaxed", (*atomic64.AtomicU64).AddRelaxed, 100, 50, 100, 150)
	testOpOrderingU64(t, "AddAcquire", (*atomic64.AtomicU64).AddAcquire, 100, 50, 100, 150)
	testOpOrderingU64(t, "AddRelease", (*atomic64.AtomicU64).AddRelease, 100, 50, 100, 150)
	testOpOrderingU64(t, "AddAcqRel", (*atomic64.AtomicU64).AddAcqRel, 100, 50, 100, 150)
	testOpOrderingU64(t, "AddSeqCst", (*atomic64.AtomicU64).AddSeqCst, 100, 50, 100, 150)
}

func TestOrU64(t *testing.T) {
	testOpU64(t, "Or", func(a *atomic64.AtomicU64, v uint64) uint64 { return a.Or(v) }, 0b1010, 0b0101, 10, 0b1111)
}

func TestOrMemoryOrderingU64(t *testing.T) {
	testOpOrderingU64(t, "OrRelaxed", (*atomic64.AtomicU64).OrRelaxed, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU64(t, "OrAcquire", (*atomic64.AtomicU64).OrAcquire, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU64(t, "OrRelease", (*atomic64.AtomicU64).OrRelease, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU64(t, "OrAcqRel", (*atomic64.AtomicU64).OrAcqRel, 0b1100, 0b0011, 0b1100, 0b1111)
	testOpOrderingU64(t, "OrSeqCst", (*atomic64.AtomicU64).OrSeqCst, 0b1100, 0b0011, 0b1100, 0b1111)
}

func TestAndU64(t *testing.T) {
	testOpU64(t, "And", func(a *atomic64.AtomicU64, v uint64) uint64 { return a.And(v) }, 0b1111, 0b1010, 15, 0b1010)
}

func TestAndMemoryOrderingU64(t *testing.T) {
	testOpOrderingU64(t, "AndRelaxed", (*atomic64.AtomicU64).AndRelaxed, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU64(t, "AndAcquire", (*atomic64.AtomicU64).AndAcquire, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU64(t, "AndRelease", (*atomic64.AtomicU64).AndRelease, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU64(t, "AndAcqRel", (*atomic64.AtomicU64).AndAcqRel, 0b1111, 0b0101, 0b1111, 0b0101)
	testOpOrderingU64(t, "AndSeqCst", (*atomic64.AtomicU64).AndSeqCst, 0b1111, 0b0101, 0b1111, 0b0101)
}

func TestCASU64(t *testing.T) {
	var a atomic64.AtomicU64
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

func TestConcurrentAddU64(t *testing.T) {
	var a atomic64.AtomicU64
	testConcurrentU64(t, &a, 100, 1000, func(a *atomic64.AtomicU64) { a.Add(1) }, 100000)
}

func TestConcurrentCASU64(t *testing.T) {
	var a atomic64.AtomicU64
	a.Store(0)

	const goroutines = 100
	var successCounter atomic64.AtomicU64

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

func TestConcurrentMixedU64(t *testing.T) {
	var a atomic64.AtomicU64
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
				a.And(0xFFFFFFFFFFFFFFFF)
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

func TestFunctionPointersU64(t *testing.T) {
	var a atomic64.AtomicU64

	atomic64.AtomicStoreU64(&a, 42)
	if got := atomic64.AtomicLoadU64(&a); got != 42 {
		t.Errorf("Function call: got %d, want 42", got)
	}

	old := atomic64.AtomicAddU64(&a, 8)
	if old != 42 {
		t.Errorf("AtomicAddU64 returned %d, want 42", old)
	}
	if got := atomic64.AtomicLoadU64(&a); got != 50 {
		t.Errorf("After AtomicAddU64: got %d, want 50", got)
	}
}

func TestOverflowU64(t *testing.T) {
	var a atomic64.AtomicU64
	a.Store(18446744073709551615)

	a.Add(1)
	if got := a.Load(); got != 0 {
		t.Errorf("Overflow: got %d, want 0", got)
	}
}

func TestZeroValueU64(t *testing.T) {
	var a atomic64.AtomicU64

	if got := a.Load(); got != 0 {
		t.Errorf("Zero value Load: got %d, want 0", got)
	}

	a.Add(5)
	if got := a.Load(); got != 5 {
		t.Errorf("Zero value after Add: got %d, want 5", got)
	}
}

func testOpU64(t *testing.T, name string, op func(*atomic64.AtomicU64, uint64) uint64, initial, delta, wantOld, wantNew uint64) {
	t.Run(name, func(t *testing.T) {
		var a atomic64.AtomicU64
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

func testOpOrderingU64(t *testing.T, name string, fn func(*atomic64.AtomicU64, uint64) uint64, initial, delta, wantOld, wantNew uint64) {
	t.Run(name, func(t *testing.T) {
		var a atomic64.AtomicU64
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

func testConcurrentU64(t *testing.T, a *atomic64.AtomicU64, goroutines, iterations int, op func(*atomic64.AtomicU64), expected uint64) {
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

func BenchmarkStoreMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreRelaxed(uint64(i))
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreRelease(uint64(i))
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.StoreSeqCst(uint64(i))
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.Store(uint64(i))
		}
	})
}

func BenchmarkLoadMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkAddMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkOrMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkAndMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
	b.Run("Relaxed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndRelaxed(0xFFFFFFFFFFFFFFFF)
		}
	})
	b.Run("Acquire", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndAcquire(0xFFFFFFFFFFFFFFFF)
		}
	})
	b.Run("Release", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndRelease(0xFFFFFFFFFFFFFFFF)
		}
	})
	b.Run("AcqRel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndAcqRel(0xFFFFFFFFFFFFFFFF)
		}
	})
	b.Run("SeqCst", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.AndSeqCst(0xFFFFFFFFFFFFFFFF)
		}
	})
	b.Run("Default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a.And(0xFFFFFFFFFFFFFFFF)
		}
	})
}

func BenchmarkConcurrentStoreMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkConcurrentLoadMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkConcurrentAddMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkConcurrentOrMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
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

func BenchmarkConcurrentAndMemoryOrderingU64(b *testing.B) {
	var a atomic64.AtomicU64
	b.Run("Relaxed", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndRelaxed(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
	b.Run("Acquire", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndAcquire(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
	b.Run("Release", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndRelease(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
	b.Run("AcqRel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndAcqRel(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
	b.Run("SeqCst", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.AndSeqCst(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
	b.Run("Default", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				a.And(0xFFFFFFFFFFFFFFFF)
			}
		})
	})
}

func BenchmarkCASU64(b *testing.B) {
	var a atomic64.AtomicU64
	for i := 0; i < b.N; i++ {
		a.Store(uint64(i))
		a.CAS(uint64(i), uint64(i+1))
	}
}

func BenchmarkConcurrentCASU64(b *testing.B) {
	var a atomic64.AtomicU64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			old := a.Load()
			a.CAS(old, old+1)
		}
	})
}
