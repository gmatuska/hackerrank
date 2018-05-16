using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

class Solution {
    static void Main(String[] args) {
        /* Enter your code here. Read input from STDIN. Prlong output to STDOUT. Your class should be named Solution */
        var tokens_s = Console.ReadLine().Split(' ');
        var n = Convert.ToInt64(tokens_s[0]);
        var graph = new Graph();
        var i = Convert.ToInt64(tokens_s[1]);
        var count = 0;
        var astronautsInPairs = new List<long>();
        while(count < i)
        {
            var tokens_i = Console.ReadLine().Split(' ');
            var astronauts = Array.ConvertAll(tokens_i, long.Parse);
            graph.AddEdge(astronauts[0],astronauts[1]);
            astronautsInPairs.AddRange(astronauts.ToList());
            count++;
        }
        //get the singleton vertices
        var singletons = n - graph.Vertices.Count;
        //run a dfs on the graph and keep track of the depths (in the graph object probably)
        var results = graph.Dfs();
        long combinations = 0;
        for (var p = 0; p < results.Count; p++)
        {
            for (var q = p+1; q < results.Count; q++)
            {
                combinations += (results[p] * results[q]);
            }
        }
        combinations += results.Sum(t => (t * singletons));
        //connect disjoint sets to themselves and aggregate combinations
        for (long p = singletons-1; p >0; p--)
        {
            combinations += p;
        }
        //take combos times the number of singleton sets
        Console.WriteLine(combinations);
    }
}
public class Graph
{
    public List<Vertex> Vertices { get; set; } 

    public Graph()
    {
        Vertices = new List<Vertex>();
    }

    public void AddEdge(long from, long to)
    {
        //find from vertex and add edge
        var fromVertex = Vertices.SingleOrDefault(v => v.Value == from);
        //add if not in vertices
        if (fromVertex == null)
        {
            fromVertex = new Vertex(from);
            Vertices.Add(fromVertex);
        }
        //check for duplicate
        if (fromVertex.Edges.All(e => e != to))
            fromVertex.Edges.Add(to);
        //find to vertex and add edge
        var toVertex = Vertices.SingleOrDefault(v => v.Value == to);
        if (toVertex == null)
        {
            toVertex = new Vertex(to);
            Vertices.Add(toVertex);
        }
        //check for duplicate
        if (toVertex.Edges.All(e => e != from))
            toVertex.Edges.Add(from);
    }

    public List<long> Dfs()
    {
        var counts = new List<long>();
        var visited = new List<long>();
        foreach (var vertex in Vertices)
        {
            if (visited.Contains(vertex.Value)) continue;
            var result = Dfs(vertex, new DfsResult(visited, 0));
            visited = result.Visited;
            counts.Add(result.Count);
        }
        return counts;
    }

    public DfsResult Dfs(Vertex vertex, DfsResult result )
    {
        var stack = new Stack<long>();
        stack.Push(vertex.Value);
        while (stack.Any())
        {
            var stackVertex = stack.Pop();
            if (!result.Visited.Contains(stackVertex))
            {
                result.Visited.Add(stackVertex);
                result.Count++;
                var edges = Vertices.Single(v => v.Value == stackVertex).Edges;
                foreach (var edge in edges)
                {
                    if (!result.Visited.Contains(edge))
                    {
                        stack.Push(edge);
                    }
                }
            }
        }
        return result;
    }
}

public class Vertex
{
    public long Value { get; }
    public List<long> Edges { get; set; }

    public Vertex(long value)
    {
        Value = value;
        Edges = new List<long>();
    }
}

public class DfsResult
{
    public List<long> Visited;
    public long Count;

    public DfsResult(List<long> visited,long count )
    {
        Visited = visited;
        Count = count;
    }
}