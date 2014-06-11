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
        v, w = e.head, e.tail
    } else {
        v, w = e.tail, e.head
    }
    return
}

// Graphs are simply lists of directed edges
// that maintain a map of visited edges
// and a directed adjacency list.
type Graph struct {
    edges []Edge
    vertices map[int]int
    adjacent map[int][]int
}

func NewGraph(e []Edge) *Graph {
    g := new(Graph)
    g.vertices = make(map[int]int)
    g.adjacent = make(map[int][]int)
    g.edges = e
    for _, v := range e {
        g.vertices[v.tail] = 0
        g.vertices[v.head] = 0
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
        g.vertices[i] = 0
    }
}

func (g *Graph) Describe() {
    log.Printf("Graph of %v vertices, %v edges", g.Len(), len(g.edges))
}

func (g *Graph) Visit(v, src int) bool {
    visited, in := g.vertices[v]
    if !in {
        log.Fatalf("%v not in Graph", v)
    }
    if visited == 0 {
        g.vertices[v] = src
        return true
    } else {
        return false
    }
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

type Stack struct {
    items []int
    pos int
}

func NewStack(c int) (*Stack) {
    return &Stack{items: make([]int, 0, c), pos: 0}
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
    s.pos++
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
        s.pos--
        x = s.items[s.pos]
    }
    return
}

// FindSCC locates all stongly connected components in a graph
func (g *Graph) CountSCC() (c int) {
    // log.Print(g.edges)
    s := g.Traverse()
    // log.Print(s)
    // log.Print("++++++++")
    g.Reset()
    // log.Print(g.vertices)
    g.SecondPass(s)
    leaders := make(map[int]bool)
    for _, l := range g.vertices {
        _, ok := leaders[l]
        if !ok && l > 0 {
            leaders[l] = true
        }
    }
    // log.Print(leaders)
    // log.Print(g.vertices)
    return len(leaders)
}

func (g *Graph) Traverse() (s *Stack){
    g.BuildAdjacencyList(true)
    s = NewStack(len(g.vertices))
    for _, e := range g.edges {
        v, _ := e.Arc(true)
        g.VisitEdges(v, -1, s)
    }
    return
}

func (g *Graph) SecondPass(s *Stack) {
    g.BuildAdjacencyList(false)
    t := NewStack(len(g.vertices))
    for {
        vertex, err := s.Pop()
        if err != nil {
            break
        }
        g.VisitEdges(vertex, vertex, t)
    }
    return
}

func (g *Graph) BuildAdjacencyList(reverse bool) {
    g.adjacent = make(map[int][]int)
    for _, e := range g.edges {
        v, w := e.Arc(reverse)
        heads, ok := g.adjacent[v]
        if ok {
            g.adjacent[v] = append(heads, w)
        } else {
            g.adjacent[v] = []int{w}
        }
    }
}

func (g *Graph) VisitEdges(v, src int, s *Stack) {
    if g.Visit(v, src) {
        edges, ok := g.adjacent[v]
        if ok {
            for _, w := range edges {
                g.VisitEdges(w, src, s)
            }
        } else {
            // log.Printf("Terminal vertex: %v", v)
        }
        s.Push(v)
    }
}

func (g *Graph) Components() map[int][]int {
    m := make(map[int][]int)
    for v, src := range g.vertices {
        list, ok := m[src]
        if ok {
            m[src] = append(list, v)
        } else {
            m[src] = []int{v}
        }
    }
    return m
}

func (g *Graph) LargestSizes(n int) []int {
    l := make([]int, n, n)
    // log.Print(g.vertices)
    // log.Print(g)
    for _, c := range g.Components() {
        size := len(c)
        for i, old := range l {
            if size > old {
                if i == 0 {
                    l = append([]int{size}, l[:n-1]...)
                } else if i == len(l) - 1 {
                    l[i] = size
                } else {
                    larger, smaller := l[:i], l[i:n-1]
                    if len(larger) + len(smaller) + 1 != n {
                        log.Fatalf("larger: %v, size: %v, smaller: %v",
                                   larger, size, smaller)
                    }
                    larger = append(larger, 0)
                    l = append(larger, smaller...)
                    l[i] = size
                    if len(l) != n {
                        log.Fatalf("Fucked up the list %v", l)
                    }
                }
                break
            }
            // log.Printf("Insert %v @ %v", size, i)
            if len(l) != n {
                log.Print(l)
                log.Fatalf("%v == %v", len(l), n)
            }
        }
    }
    return l
}
