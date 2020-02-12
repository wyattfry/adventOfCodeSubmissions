using System;
using System.Collections.Generic;
using System.Linq;

namespace Day3
{
    public class CircuitAnalyzer
    {
        private Point CurrentLocation { get; set; }

        /// <summary>
        /// <X, <Y, LinesCrossed>>
        /// </summary>
        /// <value></value>
        public Dictionary<int, Dictionary<int, int>> CellIndex { get; set; }

        public int DistanceBetweenOriginAndClosestIntersection { get; set; }

        private int CurrentLine = 1;

        public CircuitAnalyzer()
        {
            this.DistanceBetweenOriginAndClosestIntersection = int.MaxValue;

            this.CurrentLocation = new Point
            {
                X = 0,
                Y = 0,
            };

            this.CellIndex = new Dictionary<int, Dictionary<int, int>>();

            this.CellIndex.Add(0, new Dictionary<int, int>
            {
                { 0, 0 },
            });
        }

        public int GetLinesCrossedFromIndex(int x, int y)
        {
            try
            {
                return this.CellIndex[x][y];
            }
            catch (KeyNotFoundException)
            {
                return 0;
            }
        }

        public void CrossWithLine()
        {
            this.CrossWithLine(this.CurrentLine, this.CurrentLocation.X, this.CurrentLocation.Y);
        }

        public void CrossWithLine(int lineNumber, int x, int y)
        {
            var linesCrossingAtTarget = GetLinesCrossedFromIndex(x, y);

            var crossedByAnyLines = linesCrossingAtTarget > 0;

            var notCrossedByThisLine = (linesCrossingAtTarget & lineNumber) == 0;

            if (crossedByAnyLines && notCrossedByThisLine)
            {
                // New Intersection
                var distanceToOrigin = Math.Abs(x) + Math.Abs(y);
                this.DistanceBetweenOriginAndClosestIntersection = Math.Min(distanceToOrigin, this.DistanceBetweenOriginAndClosestIntersection);
            }

            if (this.CellIndex.ContainsKey(x))
            {
                this.CellIndex[x][y] = linesCrossingAtTarget | lineNumber;
            }
            else
            {
                this.CellIndex.Add(x, new Dictionary<int, int>{{y, lineNumber}});
            }
        }

        public void AddVectors(string vectors)
        {
            this.CurrentLocation.X = 0;
            this.CurrentLocation.Y = 0;
            this.AddVectors(this.ParseVectors(vectors));
            this.CurrentLine++;
        }

        public void AddVectors(List<Vector> vectors)
        {
            vectors.ForEach(v => this.AddVector(v, this.CurrentLine));
        }

        public void AddVector(Vector vector, int lineNumber)
        {
            if (vector.Distance <= 0)
            {
                return;
            }

            switch (vector.Direction)
            {
                case Direction.Up:
                    this.CurrentLocation.Y += 1;
                    break;
                case Direction.Right:
                    this.CurrentLocation.X += 1;
                    break;
                case Direction.Down:
                    this.CurrentLocation.Y -= 1;
                    break;
                case Direction.Left:
                    this.CurrentLocation.X -= 1;
                    break;
            }

            this.CrossWithLine();

            var nextVector = new Vector
            {
                Direction = vector.Direction,
                Distance = vector.Distance - 1,
            };

            this.AddVector(nextVector, lineNumber); // recursion ftw!
        }

        public List<Vector> ParseVectors(string vectorString)
        {
            List<Vector> vectors = new List<Vector>();

            vectorString.Split(",").ToList().ForEach(x =>
            {
                vectors.Add(new Vector
                {
                    Direction = this.StringToDirection(x.Substring(0, 1)), // U, R, D, L
                    Distance = int.Parse(x.Substring(1)),
                });
            });

            return vectors;
        }

        private Direction StringToDirection(string direction)
        {
            switch (direction.ToUpperInvariant())
            {
                case "U":
                    return Direction.Up;
                case "R":
                    return Direction.Right;
                case "D":
                    return Direction.Down;
                case "L":
                    return Direction.Left;
                default:
                    throw new IndexOutOfRangeException();
            }
        }
    }

    public class Point
    {
        public int X { get; set; }
        public int Y { get; set; }
    }

    public class Vector
    {
        public Direction Direction { get; set; }

        public int Distance { get; set; }
    }

    public enum Direction
    {
        Up,
        Right,
        Left,
        Down,
    }
}