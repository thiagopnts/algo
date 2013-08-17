package graph

import "container/list"

type Graph struct {
    v uint
    e uint
    adj []*list.List
}

func New(v uint) *Graph {
    graph := Graph {v: v}
    graph.adj = make([]*list.List, v)
    for i := 0; i < int(v); i++ {
        graph.adj[i] = list.New()
    }
    return &graph
}

func (this *Graph) AddEdge(source, dest uint) {
    this.adj[source].PushBack(dest)
    this.adj[dest].PushBack(source)
    this.e++
}

func (this *Graph) Adj(vertex uint) *list.List {
   return this.adj[vertex]
}

func (this *Graph) V() uint {
    return this.v
}

func (this *Graph) E() uint {
    return this.e
}

