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
type Graph struct {
    edges []Edge
    vertices map[int]int
}

func NewGraph(e []Edge) *Graph {
    g := new(Graph)
    g.vertices = make(map[int]int)
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

func (g *Graph) Visit(v, src int) bool {
    visited, in := g.vertices[v]
    if !in {
        log.Fatalf("%v not in Graph", v)
    }
    if visited == 0 {
        g.vertices[v] = src
        return true
    }
    return false
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
    s = new(Stack)
    for _, e := range g.edges {
        v, _ := e.Arc(true)
        g.VisitEdges(v, 1, s, true)
    }
    return
}

func (g *Graph) SecondPass(s *Stack) {
    t := new(Stack)
    for {
        vertex, err := s.Pop()
        if err != nil {
            break
        }
        g.VisitEdges(vertex, s.Len(), t, false)
    }
    return
}

func (g *Graph) VisitEdges(p, order int, s *Stack, reverse bool) {
    // log.Print(p, order, reverse)
    for _, e := range g.edges {
        v, w := e.Arc(reverse)
        src := order
        if v == p && g.Visit(v, src) {
            // log.Printf("%v -> %v", v, w)
            g.VisitEdges(w, order, s, reverse)
            s.Push(v)
        }
    }
    // log.Print("-------")
}
