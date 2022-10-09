package ds

type Vertex struct {
	val          int
	adjacenyList []*Vertex
}

type Graph struct {
	sourceVertex *Vertex
	vertices     []*Vertex
	undirected   bool
}

func createVertex(val int) *Vertex {
	emptySlice := []*Vertex{}
	return &Vertex{
		val:          val,
		adjacenyList: emptySlice,
	}
}

func (graph *Graph) AddVertex(val int) *Vertex {
	// create source node if not already created
	if graph.sourceVertex == nil {
		graph.sourceVertex = createVertex(val)
	}
	vertex := createVertex(val)
	graph.vertices = append(graph.vertices, vertex)
	return vertex
}

func (graph *Graph) AddEdge(vertex1 *Vertex, vertex2 *Vertex) *Vertex {
	// if directed graph, vertex1 -> vertex 2
	vertex1.adjacenyList = append(vertex1.adjacenyList, vertex2)

	// if undirected graph, vertex1 <-> vertex 2
	if graph.undirected {
		vertex2.adjacenyList = append(vertex2.adjacenyList, vertex1)
	}

	return graph.sourceVertex
}

func (graph *Graph) removeEdge(vertex1 *Vertex, vertex2 *Vertex) *Vertex {
	// if directed graph, vertex1 -x-> vertex 2
	indexOfVertex2 := Contains(vertex1.adjacenyList, vertex2)
	if indexOfVertex2 != -1 {
		vertex1.adjacenyList = Remove(vertex1.adjacenyList, indexOfVertex2)
	}

	// if undirected graph, vertex1 <-> vertex 2
	if graph.undirected {
		indexOfVertex1 := Contains(vertex2.adjacenyList, vertex1)
		if indexOfVertex1 != -1 {
			vertex2.adjacenyList = Remove(vertex2.adjacenyList, indexOfVertex1)
		}
	}

	return graph.sourceVertex
}

func (graph *Graph) BreadthFirstTraversal(sourceVertex *Vertex, results *[]*Vertex) *[]*Vertex {

	stack := []*Vertex{
		sourceVertex,
	}

	dict := make(map[int]bool)

	for len(stack) != 0 {
		curr := stack[0]
		stack = stack[1:]
		if dict[curr.val] {
			continue
		}
		dict[curr.val] = true
		*results = append(*results, curr)
		stack = append(stack, curr.adjacenyList...)
	}

	return results
}

/*
	Helper methods
*/

func Contains(slice []*Vertex, vertex *Vertex) int {
	for index, val := range slice {
		if val == vertex {
			return index
		}
	}
	return -1
}

func (graph *Graph) RemoveVertex(vertexToRemove *Vertex) *Vertex {

	// remove outgoing vertices
	for _, neighbor := range vertexToRemove.adjacenyList {
		graph.removeEdge(vertexToRemove, neighbor)
	}

	// remove incoming vertices
	for _, vertex := range graph.vertices {
		graph.removeEdge(vertex, vertexToRemove)
	}

	return graph.sourceVertex
}

func Remove(slice []*Vertex, index int) []*Vertex {

	if len(slice) <= 0 {
		return slice
	}

	l := len(slice) - 1
	slice[index] = slice[l]
	return slice[:l]
}
