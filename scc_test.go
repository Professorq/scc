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

func TestRingIsFiveSCC(t *testing.T) {
    const expected = 5
    g := NewGraph(e)
    result := g.CountSCC()
    if result != expected {
        t.Logf("%v == %v", result, expected)
        t.Fail()
    }
}

func TestFinds5SCC(t *testing.T) {
    const expected = 3
    e := []Edge{
            {1, 5},
            {2, 3},
            {3, 4},
            {4, 5},
            {4, 2},
            {5, 6},
            {6, 1},
            {6, 9},
            {9, 7},
            {7, 8},
            {8, 9},
        }
    g := NewGraph(e)
    result := g.CountSCC()
    if result != expected {
        t.Logf("%v == %v", result, expected)
        t.Fail()
    }
}
