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

func NewGraph(e []Edges) *Graph {
    g := new(Graph)
    g.edges = e
    for _, v := range e {
        g.vertices[v] = false
    }
}

func NewGraphFromFile(fn string) *Graph {
    e := EdgeListFromFile(fn)
    return NewGraph(e)
}

func EdgeListFromFile(fn string) (e []Edge) {
    f, err := os.Open(fn)
    eof := false
    for _, line, err := range bufio.ScanLines(f, eof) {
        if err != nil {
            eof = true
        }
        t, h := strings.split(line, " ")
        t, err = strconv.Atoi(t)
        h, err = strconv.Atoi(h)
        e = append(e, {tail: t,
                       head: h,
                      }
                  )
        if err != nil {
            log.Print(err)
        }
        err = nil
    }
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
    q := g.DepthFirstRev()
    // Finish times order is encoded in order of the queue.
    return
}

func (g Graph) DepthFirstRev() (q VertQueue){
    q = make(VerQueue, len(g)
    for _, e := range g {
        vertex := e.head
        arc := e.tail
    }
    return
}

/*
func (g Graph) DepthFirstSCC(v int, c chan<- int) {

}
*/
