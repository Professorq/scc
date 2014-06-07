package scc

import (
    "bufio"
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

// VertQueue is a simple FIFO queue of vertex references
type VertQueue []int
func (v VertQueue) Len() int { return len(v) }

func (v *VertQueue) Push(x int) {
    *v = append(*v, x)
}

func (v *VertQueue) Pop() int {
    old := *v
    n := len(old)
    x := old[0]
    *v = old[1 : n - 1]
    return x
}

// FindSCC locates all Stongly connected components in an input graph
func (g *Graph) FindSCC() (s []Graph) {
    // Map keeps track of all vertices that have been visited.
    // If it's visited, it's in the finish map with value 0.
    // If it's finished (i.e. all arcs have been explored),
    // the value of finish[int] == the rank of finishing
    // q := g.DepthFirstRev()
    // Finish times order is encoded in order of the queue.
    return
}

/*
func (g *Graph) DepthFirstRev() (q VertQueue){
    q = make(VertQueue, g.Len())
    for _, e := range g.edges {
        vertex := e.head
        arc := e.tail
    }
    return
}

/*
func (g Graph) DepthFirstSCC(v int, c chan<- int) {

}
*/
