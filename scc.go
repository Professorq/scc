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

func (s *Stack) Len() int { return len(s.items) }
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

// FindSCC locates all Stongly connected components in an input graph
func (g *Graph) FindSCC() (s []Graph) {
    // Map keeps track of all vertices that have been visited.
    // If it's finished (i.e. all arcs have been explored),
    // the index of the vertex in Stack == the rank of finishing
    // q := g.DepthFirstRev()
    // Finish times order is encoded in order of the stack.
    return
}

func (g *Graph) Traverse(r bool) (s *Stack){
    s = new(Stack)
    for _, e := range g.edges {
        v, _ := e.Arc(r)
        g.VisitEdges(v, s, r)
    }
    return
}

func (g *Graph) VisitEdges(p int, s *Stack, r bool) {
    for _, e := range g.edges {
        v, w := e.Arc(r)
        if v == p && g.Visit(v) {
            g.VisitEdges(w, s, r)
            s.Push(v)
        }
    }
}

/*
func (g Graph) DepthFirstSCC(v int, c chan<- int) {

}
*/
