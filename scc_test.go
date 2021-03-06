package scc

import (
    "testing"
)

func TestStack(t *testing.T) {
    const input = 8
    s := new(Stack)
    s.Push(input)
    a, err := s.Pop()
    if a != input || err != nil {
        t.Logf("%v != %v or err: %v", input, a, err)
        t.Fail()
    }
    b, err := s.Pop()
    if b != 0 || err == nil {
        t.Logf("%v != %v or err is nil: %v", 0, a, err)
        t.Fail()
    }
}

func TestPreventsSecondVisitToV(t *testing.T) {
    g := NewGraph([]Edge{
                            {1, 2},
                            {2, 3},
                        })
    first := g.Visit(3, 1)
    second := g.Visit(3, 2)
    if !first || second {
        t.Log("1st visit: %v, 2nd visit: %v", first, second)
        t.Fail()
    }
}

var e = []Edge{
            {1, 2},
            {1, 3},
            {1, 4},
            {1, 5},
            {1, 6},
            {2, 3},
            {2, 4},
            {2, 5},
            {2, 6},
            {3, 4},
            {3, 5},
            {3, 6},
            {4, 5},
            {4, 6},
            {5, 6},
}

func TestTraverseExhaustsVertices(t *testing.T) {
    var finish = []int{6, 5, 4, 3, 2, 1}
    g := NewGraph(e)
    s := g.Traverse()
    for i, v := range finish {
        a, err := s.Pop()
        if err != nil && i != len(finish) - 1 {
            t.Logf("e: %v. %v == %v", err, i, len(finish))
            t.Fail()
        } else if err != nil {
            // do nothing
        } else if a != v {
                t.Logf("%v == %v", a, v)
                t.Fail()
        }
    }
}

func TestRingIs6SCC(t *testing.T) {
    const expected = 6
    g := NewGraph(e)
    result := g.CountSCC()
    if result != expected {
        t.Log(g.vertices)
        t.Logf("%v == %v", result, expected)
        t.Fail()
    }
}

var f = []Edge{
            {9, 7},
            {9, 3},
            {8, 6},
            {8, 5},
            {7, 1},
            {4, 7},
            {1, 4},
            {3, 6},
            {6, 9},
            {5, 2},
            {2, 8},
        }

func TestLargest3are3(t *testing.T) {
    const exp = 3
    g := NewGraph(f)
    g.CountSCC()
    threes := g.LargestSizes(3)
    for _, v := range threes {
        if v != exp {
            t.Logf("%v == %v", v, exp)
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(g.Components())
    }
}

func TestFinds3SCC(t *testing.T) {
    const expected = 3
    g := NewGraph(f)
    result := g.CountSCC()
    if result != expected {
        t.Logf("%v == %v", result, expected)
        t.Fail()
    }
}

func TestAdjacencyBuild(t *testing.T) {
    g := NewGraph(e)
    g.BuildAdjacencyList(false)
    if len(g.adjacent[1]) != 5 {
        t.Logf("%v == %v", g.adjacent[1], 5)
        t.Log(g.adjacent)
        t.Fail()
    }
}

func TestJasonSemkosCase(t *testing.T) {
    expected := []int{3, 3, 3, 0, 0}
    g := NewGraphFromFile("jsemko.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestJasonSemkosCase2(t *testing.T) {
    expected := []int{3, 3, 2, 0, 0}
    g := NewGraphFromFile("js2.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestJasonSemkosCase3(t *testing.T) {
    expected := []int{3, 3, 1, 1, 0}
    g := NewGraphFromFile("js3.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestJasonSemkosCase4(t *testing.T) {
    expected := []int{7, 1, 0, 0, 0}
    g := NewGraphFromFile("js4.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestJasonSemkosCase5(t *testing.T) {
    expected := []int{6, 3, 2, 1, 0}
    g := NewGraphFromFile("js5.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestRougherCase(t *testing.T) {
    expected := []int{6, 1, 1, 0, 0}
    g := NewGraphFromFile("rough.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}

func TestLastCase(t *testing.T) {
    expected := []int{3, 2, 2, 2, 1}
    g := NewGraphFromFile("last.txt")
    g.CountSCC()
    ls := g.LargestSizes(5)
    for i, l := range ls {
        if expected[i] != l {
            t.Fail()
        }
    }
    if t.Failed() {
        t.Log(expected)
        t.Log(ls)
        t.Log(g.Components())
    }
}
