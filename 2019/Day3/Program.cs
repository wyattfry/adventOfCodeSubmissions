using System;
using System.Diagnostics;
using SixLabors.ImageSharp;
using SixLabors.ImageSharp.PixelFormats;
using SixLabors.ImageSharp.Processing;
using SixLabors.Primitives;

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

            // Small input for testing. Should produce an "R" shape
            // wire1 = "U70,R35,D30,L5,D5,L30";
            // wire2 = "U35,R10,D10,R10,D10,R10,D15";

            Console.WriteLine("Adding path 1...");

            var watch = Stopwatch.StartNew();

            circuitAnalyzer.AddVectors(wire1);

            Console.WriteLine($"Done in {watch.ElapsedMilliseconds}ms");

            Console.WriteLine("Adding path 2...");

            watch.Restart();

            circuitAnalyzer.AddVectors(wire2);

            Console.WriteLine($"Done in {watch.ElapsedMilliseconds}ms");

            Console.WriteLine($"Distance to closest intersection: {circuitAnalyzer.DistanceBetweenOriginAndClosestIntersection}");

            Console.WriteLine($"Smallest path sum at intersection: {circuitAnalyzer.SmallestPathLengthSumAtIntersection}");

            // Uncomment next line to generate a visualization
            Program.GenerateVisualization(circuitAnalyzer);
        }

        static void GenerateVisualization(CircuitAnalyzer analyzer)
        {
            var width = analyzer.UpperRight.X + Math.Abs(analyzer.LowerLeft.X) + 1;

            var height = analyzer.UpperRight.Y + Math.Abs(analyzer.LowerLeft.Y) + 1;

            var x = analyzer.LowerLeft.X;

            var y = analyzer.LowerLeft.Y;

            Program.GenerateBitmapSubsection(analyzer, 0, 0, width, height);
        }

        static void GenerateBitmapSubsection(CircuitAnalyzer analyzer, int startX, int startY, int width, int height)
        {
            Console.WriteLine($"Width: {width}, height: {height}");

            var image = new Image<Rgba32>(width, height);

            image.Mutate(ctx => ctx.BackgroundColor(Color.DarkGreen));

            int x;
            int y;

            foreach (var col in analyzer.CellIndex)
            {
                foreach (var row in col.Value)
                {
                    x = startX + col.Key + Math.Abs(analyzer.LowerLeft.X);
                    y = startY + row.Key + Math.Abs(analyzer.LowerLeft.Y);
                    try
                    {
                        for (int i = 0; i < 8; i++)
                        {
                            image[x + i,y + i] = LineToColor(row.Value.LinesCrossed);
                        }
                    }
                    catch (IndexOutOfRangeException)
                    {
                        // path width overflows canvas, ignore IOORExceptions
                    }
                }
            }

            image.Mutate(ctx => ctx.Resize(new ResizeOptions{Size = new Size(2000, 0)}));

            Console.WriteLine("Saving scaled image...");
            image.Save("image-scaled.png");
        }

        static Rgba32 LineToColor(int lines)
        {
            switch (lines)
            {
                case 0:
                    return Rgba32.DarkGreen;
                case 1:
                    return Rgba32.CornflowerBlue;
                case 2:
                    return Rgba32.Orange;
                case 3:
                    return Rgba32.Red;
                default:
                    return Rgba32.DarkGreen;
            }
        }
    }
}
