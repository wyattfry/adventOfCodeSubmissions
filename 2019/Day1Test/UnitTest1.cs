using Day1;
using Shouldly;
using System;
using Xunit;

namespace Day1Test
{
    public class UnitTest1
    {
        [Theory]
        [InlineData(12, 2)]
        [InlineData(14, 2)]
        [InlineData(1969, 654)]
        [InlineData(100756, 33583)]
        [InlineData(-4, 0)]
        public void GetRequiredFuel(int moduleMass, int expectedFuel)
        {
            // Arrange
            var sut = new FuelCalculator();

            // Act
            var result = sut.GetRequiredFuel(moduleMass);

            // Assert
            result.ShouldBe(expectedFuel);
        }
        [Theory]
        [InlineData(14, 2)]
        [InlineData(1969, 966)]
        [InlineData(100756, 50346)]
        [InlineData(-4, 0)]
        public void GetTotalRequiredFuel(int moduleMass, int expectedFuel)
        {
            // Arrange
            var sut = new FuelCalculator();

            // Act
            var result = sut.GetTotalRequiredFuel(moduleMass);

            // Assert
            result.ShouldBe(expectedFuel);
        }
    }
}
