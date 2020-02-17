using Day3;
using Shouldly;
using System.Linq;
using Xunit;

namespace Day3Test
{
    public class UnitTest1
    {
        [Fact]
        public void Given_new_analyzer_then_has_one_empty_cell()
        {
            // Arrange
            var sut = new CircuitAnalyzer();

            // Act & Assert
            sut.CellIndex.Count().ShouldBe(1);
            sut.GetLinesCrossedFromIndex(0,0).ShouldBe(0);
        }

        [Fact]
        public void When_one_vector_then_two_cells()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var lineNumber = 1;
            var vector = new Vector
            {
                Direction = Direction.Right,
                Distance = 1
            };

            // Act
            sut.AddVector(vector, lineNumber);

            // Assert
            sut.GetLinesCrossedFromIndex(1, 0).ShouldBe(1);
        }

        [Fact]
        public void When_Vector_String_Parsed_Should_not_Throw()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var vectorString = "U2,D5";

            // Act
            var result = sut.ParseVectors(vectorString);

            // Assert
            result[0].Direction.ShouldBe(Direction.Up);
            result[0].Distance.ShouldBe(2);
            result[1].Direction.ShouldBe(Direction.Down);
            result[1].Distance.ShouldBe(5);
        }

        [Fact]
        public void When_Vectors_Should_Not_Throw()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var vectorString = "D3";

            // Act
            sut.AddVectors(vectorString);

            // Assert
            sut.GetLinesCrossedFromIndex(0, -1).ShouldBe(1);
        }

        [Fact]
        public void When_Multiple_Wires_Should_Not_Throw()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var wire1 = "U2";
            var wire2 = "R1,U1,L1";

            // Act
            sut.AddVectors(wire1);
            sut.AddVectors(wire2);

            // Assert
            sut.GetLinesCrossedFromIndex(0, 1).ShouldBe(3);
        }

        [Theory]
        [InlineData("U2", "R1,U1,L1", 1)]
        [InlineData("R8,U5,L5,D3", "U7,R6,D4,L4", 6)] // example from website
        public void When_Intersection_Then_Distance_From_Origin_Should_Be_Set(string p1, string p2, int distance)
        {
            // Arrange
            var sut = new CircuitAnalyzer();

            // Act
            sut.AddVectors(p1);
            sut.AddVectors(p2);

            // Assert
            sut.DistanceBetweenOriginAndClosestIntersection.ShouldBe(distance);
        }

        [Fact]
        public void When_Multiple_Intersections_Then_Distance_From_Origin_Should_Be_Correct()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var wire1 = "U5";
            var wire2 = "R1,U1,L2,U1,R2";

            // Act
            sut.AddVectors(wire1);
            sut.AddVectors(wire2);

            // Assert
            sut.DistanceBetweenOriginAndClosestIntersection.ShouldBe(1);
        }

        [Fact]
        public void When_intersections_then_smallest_path_is_correct()
        {
            // Arrange
            var sut = new CircuitAnalyzer();
            var wire1 = "U5";
            var wire2 = "R1,U1,L2,U1,R2";

            // Act
            sut.AddVectors(wire1);
            sut.AddVectors(wire2);

            // Assert
            sut.SmallestPathLengthSumAtIntersection.ShouldBe(4);
        }
    }
}
