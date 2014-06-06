package scc

import (
    "testing"
)

func TestQueue(t *testing.T) {
    expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
    q := make(VertQueue, 0, 20)
    for i, s := range expected {
        q.Push(i)
        q.Push(s)
    }
    const length = 20
    if q.Len() != length {
        t.Logf("%v != %v", q.Len(), length)
        t.Log(q)
        t.Fail()
    }
    o := []int{}
    for q.Len() > 0 {
        o = append(o, q.Pop())
    }
    for i, v := range o {
        switch i % 2 {
        case 0:
            if v != i / 2 {
                t.Log("index and value: %v/2  != %v", i, v)
                t.Fail()
            }
        case 1:
            e := expected[i / 2] 
            if v != e {
                t.Log("%v != %v", v, e)
                t.Fail()
            }
        }
    }
}

func TestGraphInit(t *testing.T) {
    const last = 875714
    g := NewGraphFromFile("SCC.text")
    length := len(*g)
    if length != last {
        t.Logf("%v != %v", length, last)
        t.Fail()
    }
}
