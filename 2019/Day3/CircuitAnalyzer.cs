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
        public Dictionary<int, Dictionary<int, CircuitInfo>> CellIndex { get; set; }

        public int DistanceBetweenOriginAndClosestIntersection { get; set; }
        public int SmallestPathLengthSumAtIntersection { get; set; }

        public Point UpperRight { get; set; }
        public Point LowerLeft { get; set; }

        private int CurrentLine = 1;
        private int CurrentPathLength = 0;

        public CircuitAnalyzer()
        {
            this.DistanceBetweenOriginAndClosestIntersection = int.MaxValue;
            this.SmallestPathLengthSumAtIntersection = int.MaxValue;

            this.CurrentLocation = new Point
            {
                X = 0,
                Y = 0,
            };

            this.UpperRight = new Point
            {
                X = 0,
                Y = 0,
            };

            this.LowerLeft = new Point
            {
                X = 0,
                Y = 0,
            };

            this.CellIndex = new Dictionary<int, Dictionary<int, CircuitInfo>>();

            this.CellIndex.Add(0, new Dictionary<int, CircuitInfo>
            {
                { 0, new CircuitInfo
                    {
                        LinesCrossed = 0,
                    }
                },
            });
        }

        public int GetLinesCrossedFromIndex(int x, int y)
        {
            try
            {
                return this.CellIndex[x][y].LinesCrossed;
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
                var pathSum = this.CellIndex[x][y].PathLengthToHere + this.CurrentPathLength;
                this.SmallestPathLengthSumAtIntersection = Math.Min(pathSum, this.SmallestPathLengthSumAtIntersection);
            }

            if (this.CellIndex.ContainsKey(x))
            {
                if (this.CellIndex[x].ContainsKey(y))
                {
                    this.CellIndex[x][y].LinesCrossed = linesCrossingAtTarget | lineNumber;
                }
                else
                {
                    this.CellIndex[x].Add(y, new CircuitInfo
                    {
                        LinesCrossed = lineNumber
                    });
                }
            }
            else
            {
                this.CellIndex.Add(x, new Dictionary<int, CircuitInfo>
                {
                    {y, new CircuitInfo
                        {
                            LinesCrossed = lineNumber,
                        }
                    }
                });
            }
            this.CellIndex[x][y].PathLengthToHere += this.CurrentPathLength;
        }

        public void AddVectors(string vectors)
        {
            this.CurrentLocation.X = 0;
            this.CurrentLocation.Y = 0;
            this.CurrentPathLength = 0;
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
            this.CurrentPathLength++;

            this.CrossWithLine();

            this.UpdateBounds();

            var nextVector = new Vector
            {
                Direction = vector.Direction,
                Distance = vector.Distance - 1,
            };

            this.AddVector(nextVector, lineNumber); // recursion ftw!
        }

        private void UpdateBounds()
        {
            this.UpperRight.X = Math.Max(this.UpperRight.X, this.CurrentLocation.X);
            this.UpperRight.Y = Math.Max(this.UpperRight.Y, this.CurrentLocation.Y);
            this.LowerLeft.X = Math.Min(this.LowerLeft.X, this.CurrentLocation.X);
            this.LowerLeft.Y = Math.Min(this.LowerLeft.Y, this.CurrentLocation.Y);
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

        override public String ToString()
        {
            return $"[Point] X: {this.X}, Y: {this.Y}";
        }
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

    public class CircuitInfo
    {
        public int LinesCrossed { get; set; }

        /// <summary>
        ///  <int LineNumber, int PathLength>
        /// </summary>
        /// <value></value>
        public int PathLengthToHere { get; set; }
    }
}