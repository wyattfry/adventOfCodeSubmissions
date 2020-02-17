using Day4;
using Xunit;
using Shouldly;

namespace Day4Test
{
    public class UnitTest1
    {
        public PasswordGenerator sut { get; set; }
        public UnitTest1()
        {
            this.sut = new PasswordGenerator();
        }
        [Theory]
        [InlineData("12344")]
        [InlineData("1234557")]
        public void When_invalid_length_should_not_be_valid(string invalidPassword)
        {
            // Arrange

            // Act
            var result = this.sut.IsValid(invalidPassword);

            // assess
            result.ShouldBeFalse();
        }

        [Fact]
        public void When_valid_length_should_be_valid()
        {
            // Arrange
            var validLengthPassword = "113456";

            // Act
            var result = this.sut.IsValid(validLengthPassword);

            // Assert
            result.ShouldBeTrue();
        }

        [Fact]
        public void When_no_adjacent_digits_then_invalid()
        {
            // Arrange
            var noAdjacentDigits = "123456";

            // Act
            var result = this.sut.IsValid(noAdjacentDigits);

            // Assert
            result.ShouldBeFalse();
        }

        [Fact]
        public void When_adjacent_digits_then_valid()
        {
            // Arrange
            var noAdjacentDigits = "113456";

            // Act
            var result = this.sut.IsValid(noAdjacentDigits);

            // Assert
            result.ShouldBeTrue();
        }

        [Fact]
        public void When_decreasing_then_invalid()
        {
            // Arrange
            var decreasingDigits = "123454";

            // Act
            var result = this.sut.IsValid(decreasingDigits);

            // Assess
            result.ShouldBeFalse();
        }

        [Theory]
        [InlineData("111111")]
        [InlineData("111112")]
        public void When_not_decreasing_then_valid(string notDecreasingPassword)
        {
            // Act
            var result = this.sut.IsValid(notDecreasingPassword);

            // Assess
            result.ShouldBeTrue();
        }

        [Theory]
        [InlineData("2")]
        public void When_could_not_become_valid_then_false(string couldBecomeValidPassword)
        {
            // Act
            var result = this.sut.CouldBecomeValid(couldBecomeValidPassword);

            // Assert
            result.ShouldBeFalse();
        }
        [Theory]

        [InlineData("4")]
        public void When_could_become_valid_then_true(string couldBecomeValidPassword)
        {
            // Act
            var result = this.sut.CouldBecomeValid(couldBecomeValidPassword);

            // Assert
            result.ShouldBeTrue();
        }

    }
}
