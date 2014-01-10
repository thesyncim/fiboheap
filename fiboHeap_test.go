package fiboHeap

import "github.com/thesyncim/fiboheap/node"
import "testing"

// import "math/rand"
// import "time"
//import "fmt"

func TestMap(t *testing.T) {
	ranks := make(map[int]*node.Node)
	if len(ranks) != 0 {
		t.Error("Shouldn't a new array be empty??")
	}

	if ranks[0] != nil {
		t.Error("Shouldn't empty entries to pointers be nil?")
	}
}

func BenchmarkPush(b *testing.B) {
	f := NewHeap()

	for i := 0; i < b.N; i++ {
		f.Insert(float64(i), nil)

	}
}

func AllocN(n int) *FiboHeap {
	f := NewHeap()
	for i := 0; i < n; i++ {
		f.Insert(float64(i), nil)
	}
	return f
}

func BenchmarkPop(b *testing.B) {
	f := AllocN(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f.ExtractMin()

	}
}

func TestConsole(t *testing.T) {
	// fmt.Println("Created heap")
	f := NewHeap()
	f.Insert(1.0, "a")
	// fmt.Printf("inserted 1")
	f.Insert(2.0, "b")
	f.Insert(3.0, "c")
	f.Insert(-1.0, "d")
	// fmt.Println(f.String())
	//fmt.Println(f.ExtractMin())
	// fmt.Println(f.String())
	// f.Insert(4.0)
	// f.Insert(5.0)
	// f.Insert(6.0)
	// f.Insert(-1.0)
	// fmt.Println(f.String())
	// f.ExtractMin()
	// fmt.Println(f.String())
	if false {
		t.Error("WTF?")
	}
}

// func TestConsoleB(t *testing.T) {
// 	f := NewHeap()

// 	f.Insert(158.0)

// 	n := f.Insert(72.0)
// 	n.AddChild(node.NewNode(144.0, nil))

// 	n1 := f.Insert(81.1)
// 	n1.AddChild(node.NewNode(116, nil))
// 	c := node.NewNode(90.0, nil)
// 	c.AddChild(node.NewNode(123.0, nil))
// 	n1.AddChild(c)

// 	fmt.Println(f)

// 	f.ExtractMin()

// 	fmt.Println(f)
// }
/*
func TestConsoleB(t *testing.T) {
	g := NewHeap()
	g.Insert(58336.00)

	n1 := g.Insert(57420.00)
	n1.AddChild(node.NewNode(61358.00, nil))

	n2 := g.Insert(57225.00)
	n2.AddChild(node.NewNode(61414.00, nil))
	n3 := node.NewNode(57230.00, nil)
	n3.AddChild(node.NewNode(58278.00, nil))
	n2.AddChild(n3)

	n4 := g.Insert(58057.00)

	n4.AddChild(node.NewNode(58712.00, nil))

	n5 := node.NewNode(58744.00, nil)
	n5.AddChild(node.NewNode(59433.00, nil))
	n4.AddChild(n5)

	n6 := node.NewNode(58389.00, nil)
	n6.AddChild(node.NewNode(61718.00, nil))
	n7 := node.NewNode(58742.00, nil)
	n7.AddChild(node.NewNode(59275.00, nil))
	n6.AddChild(n7)
	n4.AddChild(n6)

	// fmt.Printf("Heap: %s\n", g)
	if g.String() != "(58336.00), (57420.00: (61358.00)), (57225.00: (61414.00), (57230.00: (58278.00))), (58057.00: (58712.00), (58744.00: (59433.00)), (58389.00: (61718.00), (58742.00: (59275.00))))" {
		t.Error("Heap did not build correctly")
	}

	m := g.ExtractMin()
	if m != n2.Value {
		t.Errorf("Expected min = %f, got %f.\n", n2.Value, m)
	}
	if !g.CheckHeap() {
		t.Errorf("Heap did not consolidate correctly. %d/%d != %s", g.Count(), g.Len(), g)
	}
}
*/
func TestConsoleC(t *testing.T) {
	h := NewHeap()
	// rand.Seed(time.Now().Unix())
	n := 11
	for i := 0; i < n; i++ {
		h.Insert(float64(n^2-i), "a")
	}

	if h.Count() != n {
		t.Error("Random heap not built correctly!")
	}

	if !h.CheckHeap() {
		t.Errorf("Heap corrupted: %d/%d", h.Count(), h.Len())
	}

	for h.Count() != 0 && h.CheckHeap() {
		h.ExtractMin()
	}

	if h.String() != "[Empty Heap]" {
		t.Errorf("\nHeap corrupted: %d/%d;\n%s", h.Count(), h.Len(), h)
	}
}
