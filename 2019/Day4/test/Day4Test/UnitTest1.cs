using Day4;
using Xunit;
using Shouldly;
using static Day4.PasswordGenerator;

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
            var sut = new LengthValidator(6);

            // Act
            var result = sut.IsValid(invalidPassword);

            // assess
            result.ShouldBeFalse();
        }

        [Fact]
        public void When_valid_length_should_be_valid()
        {
            // Arrange
            var sut = new LengthValidator(6);
            var validLengthPassword = "556677";

            // Act
            var result = sut.IsValid(validLengthPassword);

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
            var sut = new AdjacentDigitValidator();

            // Act
            var result = sut.IsValid(noAdjacentDigits);

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
            // Arrange
            var sut = new NotDecreasingValidator();

            // Act
            var result = sut.IsValid(notDecreasingPassword);

            // Assess
            result.ShouldBeTrue();
        }

        [Theory]
        [InlineData("2")]
        [InlineData("44444")]
        public void When_could_not_become_valid_then_false(string couldBecomeValidPassword)
        {
            // Act
            var result = this.sut.CouldBecomeValid(couldBecomeValidPassword);

            // Assert
            result.ShouldBeFalse();
        }
        [Theory]
        [InlineData("4")]
        [InlineData("4444")]
        [InlineData("4445")]
        [InlineData("77778")]
        public void When_could_become_valid_then_true(string couldBecomeValidPassword)
        {
            // Act
            var result = this.sut.CouldBecomeValid(couldBecomeValidPassword);

            // Assert
            result.ShouldBeTrue();
        }

        [Theory]
        [InlineData("112233")]
        [InlineData("111122")]
        public void When_digit_pair_then_valid(string validPassword)
        {
            // Arrange
            var sut = new DigitPairValidator(6);

            // Act
            var result = sut.IsValid(validPassword);

            // Assert
            result.ShouldBeTrue();
        }

        [Theory]
        [InlineData("123444")]
        public void When_no_digit_pair_then_invalid(string invalidPassword)
        {
            // Arrange
            var sut = new DigitPairValidator(6);

            // Act
            var result = sut.IsValid(invalidPassword);

            // Assert
            result.ShouldBeFalse();
        }
    }
}
