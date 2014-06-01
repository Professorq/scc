package scc

// import "log"

// Directed edge in a graph
type Edge struct {
    tail int   // origin
    head int   // destination
}

// Graphs are simply lists of directed edges
type Graph []Edge

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
func (g Graph) FindSCC() (components []Graph) {
    // Map keeps track of all vertices that have been visited.
    // If it's visited, it's in the finish map with value 0.
    // If it's finished (i.e. all arcs have been explored),
    // the value of finish[int] == the rank of finishing
    // finish := make(map[int]int)
    return
}

/*
func (g Graph) DepthFirst(v, round int) (finish VertQueue) {
    var T int
    var s int
    for _, e := range g {
        switch round {
        case 1:
            vertex := e.head
            arc := e.tail
        case 2:
            vertex := e.tail
            arc := e.head
        default:
            log.Fatal("We should only ever be in round 1 or 2")
        }
    }
    return
}
*/
