package ds

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {

	var graph = Graph{
		undirected: true,
	}

	t.Run("Inserting vertices to graph", func(t *testing.T) {
		vertex1 := graph.AddVertex(1)
		if vertex1.val != 1 {
			t.Fatalf("expected %d, got %d", 1, vertex1.val)
		}
	})
	t.Run("Inserting vertices to graph", func(t *testing.T) {
		vertex2 := graph.AddVertex(2)
		if vertex2.val != 2 {
			t.Fatalf("expected %d, got %d", 2, vertex2.val)
		}
	})
}

func TestAddEdge(t *testing.T) {

	t.Run("Adding edges for vertices in undirected graph", func(t *testing.T) {

		var graph = Graph{
			undirected: true,
		}

		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)

		graph.AddEdge(vertex1, vertex2)
		expectedAdjacenyList := []*Vertex{
			vertex2,
		}
		if reflect.DeepEqual(vertex1.adjacenyList, expectedAdjacenyList) == false {
			t.Fatalf("expected %v, got %v", expectedAdjacenyList, vertex1.adjacenyList)
		}
	})

	t.Run("Adding edges for vertices in directed graph", func(t *testing.T) {
		var graph = Graph{
			undirected: false,
		}
		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)
		graph.AddEdge(vertex1, vertex2)
		expectedAdjacenyList := []*Vertex{}
		if reflect.DeepEqual(vertex2.adjacenyList, expectedAdjacenyList) == false {
			t.Fatalf("expected %v, got %v", expectedAdjacenyList, vertex1.adjacenyList)
		}
	})
}

func TestRemoveEdge(t *testing.T) {

	t.Run("Removing edges in undirected graph", func(t *testing.T) {

		var graph = Graph{
			undirected: true,
		}

		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)
		graph.AddEdge(vertex1, vertex2)

		// removing the edges
		graph.removeEdge(vertex1, vertex2)
		expectedAdjacenyList := []*Vertex{}
		if reflect.DeepEqual(vertex1.adjacenyList, expectedAdjacenyList) == false {
			t.Fatalf("expected %v, got %v", expectedAdjacenyList, vertex1.adjacenyList)
		}
		if reflect.DeepEqual(vertex2.adjacenyList, expectedAdjacenyList) == false {
			t.Fatalf("expected %v, got %v", expectedAdjacenyList, vertex2.adjacenyList)
		}
	})

	t.Run("Adding edges for vertices in directed graph", func(t *testing.T) {
		var graph = Graph{
			undirected: false,
		}
		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)
		graph.AddEdge(vertex1, vertex2)

		// removing the edges
		graph.removeEdge(vertex1, vertex2)
		expectedAdjacenyList := []*Vertex{}
		if reflect.DeepEqual(vertex1.adjacenyList, expectedAdjacenyList) == false {
			t.Fatalf("expected %v, got %v", expectedAdjacenyList, vertex1.adjacenyList)
		}
	})
}

func TestRemoveVertex(t *testing.T) {
	t.Run("removing vertex from the directed graph", func(t *testing.T) {
		var graph = Graph{
			undirected: false,
		}

		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)
		vertex3 := graph.AddVertex(3)
		vertex4 := graph.AddVertex(4)
		graph.AddEdge(vertex1, vertex2)
		graph.AddEdge(vertex2, vertex4)
		graph.AddEdge(vertex1, vertex3)
		graph.AddEdge(vertex3, vertex4)

		graph.RemoveVertex(vertex3)

		for _, neighbor := range vertex1.adjacenyList {
			if neighbor == vertex3 {
				t.Fatalf("vertex3 is not removed from vertex1's adjacency list")
			}
		}

		if len(vertex3.adjacenyList) > 0 {
			t.Fatalf("vertex3's adjacency list is not empty")
		}

	})

	t.Run("removing vertex from the undirected graph", func(t *testing.T) {
		var graph = Graph{
			undirected: true,
		}

		vertex1 := graph.AddVertex(1)
		vertex2 := graph.AddVertex(2)
		vertex3 := graph.AddVertex(3)
		vertex4 := graph.AddVertex(4)
		graph.AddEdge(vertex1, vertex2)
		graph.AddEdge(vertex2, vertex4)
		graph.AddEdge(vertex1, vertex3)
		graph.AddEdge(vertex3, vertex4)

		graph.RemoveVertex(vertex3)

		for _, neighbor := range vertex1.adjacenyList {
			if neighbor == vertex3 {
				t.Fatalf("vertex3 is not removed from vertex1's adjacency list")
			}
		}

		if len(vertex3.adjacenyList) > 0 {
			t.Fatalf("vertex3's adjacency list is not empty")
		}

	})
}

func TestBreadthFirstTraversal(t *testing.T) {

	var graph = Graph{
		undirected: true,
	}

	vertex1 := graph.AddVertex(1)
	vertex2 := graph.AddVertex(2)
	vertex3 := graph.AddVertex(3)
	vertex4 := graph.AddVertex(4)
	vertex5 := graph.AddVertex(5)

	graph.AddEdge(vertex1, vertex2)
	graph.AddEdge(vertex1, vertex3)
	graph.AddEdge(vertex3, vertex4)
	graph.AddEdge(vertex2, vertex5)

	t.Run("Should traverse in breadth first order", func(t *testing.T) {

		expected := []*Vertex{
			vertex1,
			vertex2,
			vertex3,
			vertex5,
			vertex4,
		}

		results := []*Vertex{}
		graph.BreadthFirstTraversal(vertex1, &results)

		if reflect.DeepEqual(expected, results) == false {
			t.Fatalf("expected %v, received %v", expected, results)
		}
	})

	graph.AddEdge(vertex2, vertex4)

	t.Run("Should traverse with cycle", func(t *testing.T) {

	})

	t.Run("Should traverse with only one vertex", func(t *testing.T) {

	})
}

func TestMain(t *testing.T) {
	t.Run("remove vertex from adjacency list", func(t *testing.T) {
		vertex1 := Vertex{
			val: 1,
		}
		adjacencyList := []*Vertex{}

		result := Remove(adjacencyList, vertex1.val)

		if reflect.DeepEqual(result, adjacencyList) == false {
			t.Fatalf("expected %v, received %v", adjacencyList, result)
		}

		adjacencyList = append(adjacencyList, &vertex1)
		received := Remove(adjacencyList, 0)

		if reflect.DeepEqual(received, []*Vertex{}) == false {
			t.Fatalf("expected %v, received %v", []*Vertex{}, received)
		}

	})
}
