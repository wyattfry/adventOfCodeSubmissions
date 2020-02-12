using Day2;
using Shouldly;
using System;
using Xunit;

namespace Day2Test
{
    public class UnitTest1
    {
        [Theory]
        [InlineData("1,0,0,0,99", "2,0,0,0,99")]
        [InlineData("1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50")]
        [InlineData("2,3,0,3,99", "2,3,0,6,99")]
        [InlineData("2,4,4,5,99,0", "2,4,4,5,99,9801")]
        [InlineData("1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99")]
        public void RunTest(string inputProgramString, string resultProgramString)
        {
            // Arrange
            var sut = new IntcodeExecutor();

            // Act
            var result = sut.Run(StringToIntArray(inputProgramString));

            // Assert
            result.ShouldBe(StringToIntArray(resultProgramString));
        }

        private int[] StringToIntArray(string input)
        {
            return Array.ConvertAll(input.Split(','), s => int.Parse(s));
        }
    }
}
