package paths

import (
"datastructures/graph"
"container/list"
)

type Paths struct {
    g *graph.Graph
    source uint
    marked []bool
    edgeTo []uint
}

func New(g *graph.Graph, source uint) *Paths {
    paths := new(Paths)
    paths.g = g
    paths.source = source
    paths.marked = make([]bool, g.V())
    paths.edgeTo = make([]uint, g.V())
    paths.dfs(g, source)
    return paths
}

func (this *Paths) HasPathTo(i uint) bool {
    return this.marked[i]
}

func (this *Paths) PathTo(v uint) *list.List {
    if !this.HasPathTo(v) {
        return nil
    }
    path := list.New()
    for x := v; x != this.source; x = this.edgeTo[x] {
        path.PushBack(x)
    }
    path.PushBack(this.source)
    return path
}

func (this *Paths) dfs(graph *graph.Graph, source uint) {
    this.marked[source] = true
    for e := graph.Adj(source).Front(); e != nil; e = e.Next() {
        switch value := e.Value.(type) {
        case uint:
            if !this.marked[value] {
                this.dfs(graph, value)
                this.edgeTo[value] = source
            }
        }
    }
}

