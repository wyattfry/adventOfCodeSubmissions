using System;

namespace Day3
{
    class Program
    {
        static void Main(string[] args)
        {
            // Summary: given 2 series of direction (U,D,L,R) and distance (positive ints) pairs,
            // both describing the path of a wire on a circuit board, both paths originating
            // from same point, find the Manhattan distance between the origin and the closest
            // intersection of the 2 "wires".
            // See unit tests for examples.
            var inputText = System.IO.File.ReadAllLines("input.txt");

            var circuitAnalyzer = new CircuitAnalyzer();

            var wire1 = inputText[0];
            var wire2 = inputText[1];

            circuitAnalyzer.AddVectors(wire1);
            circuitAnalyzer.AddVectors(wire2);

            Console.WriteLine($"Distance to closest intersection: {circuitAnalyzer.DistanceBetweenOriginAndClosestIntersection}");
        }
    }
}
