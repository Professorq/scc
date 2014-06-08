package scc

import (
    "bufio"
    "errors"
    "log"
    "os"
    "strings"
    "strconv"
)

// Directed edge in a graph
type Edge struct {
    tail int   // origin
    head int   // destination
}

func (e Edge) Arc(r bool) (v, w int) {
    if r {
        v, w = e.tail, e.head
    } else {
        v, w = e.head, e.tail
    }
    return
}

// Graphs are simply lists of directed edges
type Graph struct {
    edges []Edge
    vertices map[int]bool
}

func NewGraph(e []Edge) *Graph {
    g := new(Graph)
    g.vertices = make(map[int]bool)
    g.edges = e
    for _, v := range e {
        g.vertices[v.tail] = false
        g.vertices[v.head] = false
    }
    return g
}

func NewGraphFromFile(fn string) *Graph {
    e := EdgeListFromFile(fn)
    return NewGraph(e)
}

func (g *Graph) Len() int {
    return len(g.vertices)
}

func (g *Graph) Reset() {
    for i := range g.vertices {
        g.vertices[i] = false
    }
}

func (g *Graph) Visit(v int) bool {
    visited, in := g.vertices[v]
    if !in {
        log.Fatalf("%v not in Graph", v)
    }
    if !visited {
        g.vertices[v] = true
    }
    return !visited
}

func EdgeListFromFile(fn string) (e []Edge) {
    f, err := os.Open(fn)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Split(string(line), " ")
        t, h := words[0], words[1]
        tail, err := strconv.Atoi(t)
        head, err := strconv.Atoi(h)
        e = append(e, Edge{tail: tail, head: head})
        if err != nil {
            log.Print(err)
        }
        err = nil
    }
    return
}

/*
type Queue []int
func (v Queue) Len() int { return len(v) }
func (v *Queue) Push(x int) {
    *v = append(*v, x)
}
func (v *Queue) Pop() int {
    old := *v
    n := len(old)
    x := old[0]
    *v = old[1 : n - 1]
    return x
}
*/

type Stack struct {
    items []int
    pos int
}

func (s *Stack) Len() int { return s.pos + 1}
func (s *Stack) Head() (h int, err error) {
    if s.pos > 0 {
        h = s.items[0]
    } else {
        err = errors.New("Stack is empty")
    }
    return
}

func (s *Stack) Push(x int) {
    s.pos  += 1
    if s.items == nil {
        s.items = []int{x}
    } else if s.pos < len(s.items) {
        s.items[s.pos] = x
    } else {
        s.items = append(s.items, x)
    }
}

func (s *Stack) Pop() (x int, err error){
    if s.pos == 0 {
        err = errors.New("Stack is empty")
    } else {
        s.pos -= 1
        x = s.items[s.pos]
    }
    return
}

// FindSCC locates all stongly connected components in a graph
func (g *Graph) CountSCC() (c int) {
    s := g.Traverse(true)
    g.Reset()
    leaders := g.SecondPass(s)
    return len(leaders)
}

func (g *Graph) Traverse(r bool) (s *Stack){
    s = new(Stack)
    for _, e := range g.edges {
        v, _ := e.Arc(r)
        g.VisitEdges(v, s, r)
    }
    return
}

func (g *Graph) SecondPass(s *Stack) (f []int) {
    for {
        vertex, err := s.Pop()
        if err != nil {
            break
        }
        s := new(Stack)
        if g.Visit(vertex) {
            f = append(f, vertex)
        }
        g.VisitEdges(vertex, s, false)
    }
    return
}

func (g *Graph) VisitEdges(p int, s *Stack, reverse bool) {
    for _, e := range g.edges {
        v, w := e.Arc(reverse)
        if v == p && g.Visit(v) {
            g.VisitEdges(w, s, reverse)
            s.Push(v)
        }
    }
}

/*
func (g Graph) DepthFirstSCC(v int, c chan<- int) {

}
*/
