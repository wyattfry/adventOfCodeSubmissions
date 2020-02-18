using System;
using System.Collections.Generic;

namespace Day4
{
    public class PasswordGenerator
    {
        public int RequiredLength { get; set; }

        public (int min, int max) PasswordRangeInclusive = (min: 402328, max: 864247);
        public PasswordGenerator()
        {
            this.RequiredLength = 6;

            this.validators = new List<IValidator>
            {
                new LengthValidator(this.RequiredLength),
                new AdjacentDigitValidator(),
                new NotDecreasingValidator(),
                new RangeValidator(this.PasswordRangeInclusive.min, this.PasswordRangeInclusive.max),
                new DigitPairValidator(this.RequiredLength), // Part 2 requirement
            };
        }

        private List<IValidator> validators { get; set; }

        public bool IsValid(string potentialPassword)
        {
            foreach (IValidator v in this.validators)
            {
                var isValid = v.IsValid(potentialPassword);
                if (isValid == false)
                {
                    return false;
                }
            }

            return true;
        }
        public bool CouldBecomeValid(string potentialPassword)
        {
            foreach (IValidator v in this.validators)
            {
                if (v.CouldBecomeValid(potentialPassword) == false)
                {
                    return false;
                }
            }

            return true;
        }

        public int GetValidPasswordCount()
        {
            return GeneratePasswordsRecursively(string.Empty);
        }

        public int GeneratePasswordsRecursively(string potentialPassword)
        {
            // pp = "111222";
            if (this.IsValid(potentialPassword))
            {
                return 1;
            }
            if (this.CouldBecomeValid(potentialPassword))
            {
                var passwordsFound = 0;

                // Shortcut due to NotDecreasing rule
                int startingDigit;

                if (string.IsNullOrEmpty(potentialPassword))
                {
                    startingDigit = 0;
                }
                else
                {
                    startingDigit = int.Parse(potentialPassword) % 10;
                }

                for (int i = startingDigit; i <= 9; i++)
                {
                    var nextPotentialPassword = potentialPassword + i.ToString();

                    passwordsFound += GeneratePasswordsRecursively(nextPotentialPassword);
                }
                return passwordsFound;
            }
            else
            {
                return 0;
            }
        }
        public interface IValidator
        {
            bool IsValid(string potentialPassword);

            bool CouldBecomeValid(string potentialpassword);
        }

        public class RangeValidator : IValidator
        {
            public int MinValueAllowed { get; set; }
            public int MaxValueAllowed { get; set; }
            public RangeValidator(int min, int max)
            {
                this.MinValueAllowed = min;
                this.MaxValueAllowed = max;
            }
            public bool CouldBecomeValid(string potentialpassword)
            {
                if (String.IsNullOrEmpty(potentialpassword))
                {
                    return true;
                }

                var denominator = 1;

                var parsedPassword = int.Parse(potentialpassword);

                while (this.MinValueAllowed / denominator > 0)
                {
                    if (parsedPassword >= (this.MinValueAllowed / denominator) && parsedPassword <= (this.MaxValueAllowed / denominator))
                    {
                        return true;
                    }
                    denominator *= 10;
                }

                return false;
            }

            public bool IsValid(string potentialPassword)
            {
                try
                {
                    var parsedInt = int.Parse(potentialPassword);
                    return parsedInt >= this.MinValueAllowed && parsedInt <= this.MaxValueAllowed;
                }
                catch
                {
                    return false;
                }
            }
        }

        public class LengthValidator : IValidator
        {
            public int RequiredLength { get; set; }

            public LengthValidator(int requiredLength)
            {
                this.RequiredLength = requiredLength;
            }

            public bool IsValid(string potentialPassword)
            {
                return potentialPassword.Length == this.RequiredLength;
            }

            public bool CouldBecomeValid(string potentialPassword)
            {
                return potentialPassword.Length < this.RequiredLength;
            }
        }

        public class AdjacentDigitValidator : IValidator
        {
            public bool IsValid(string potentialPassword)
            {
                var hasAdjacentDigits = false;

                for (var i = 0; i < potentialPassword.Length - 1; i++)
                {
                    if (potentialPassword[i] == potentialPassword[i + 1])
                    {
                        hasAdjacentDigits = true;
                    }
                }

                return hasAdjacentDigits;
            }

            public bool CouldBecomeValid(string potentialPassword)
            {
                return potentialPassword.Length < 6;
            }
        }

        public class NotDecreasingValidator : IValidator
        {
            public bool IsValid(string potentialPassword)
            {
                for (var i = 0; i < potentialPassword.Length - 1; i++)
                {
                    if (potentialPassword[i] > potentialPassword[i + 1])
                    {
                        return false;
                    }

                }
                return true;
            }

            public bool CouldBecomeValid(string potentialPassword)
            {
                return this.IsValid(potentialPassword);
            }
        }

        public class DigitPairValidator : IValidator
        {
            public int RequiredLength { get; set; }

            public DigitPairValidator(int requiredLength)
            {
                this.RequiredLength = requiredLength;
            }

            public bool CouldBecomeValid(string potentialpassword)
            {
                if (potentialpassword.Length <= this.RequiredLength - 2)
                {
                    return true;
                }
                if (potentialpassword.Length == this.RequiredLength - 1)
                {
                    for (var i = 0; i < potentialpassword.Length - 1; i++)
                    {
                        if (potentialpassword[i] != potentialpassword[i + 1])
                        {
                            return true;
                        }
                    }
                    return false;
                }
                return false;
            }

            public bool IsValid(string potentialPassword)
            {
                var pointer = 0;
                var groupSize = 0;
                while (pointer < potentialPassword.Length)
                {
                    groupSize++;

                    var isBoundary = pointer + 1 > potentialPassword.Length - 1
                        || potentialPassword[pointer] != potentialPassword[pointer + 1];

                    if (isBoundary)
                    {
                        if (groupSize == 2)
                        {
                            return true;
                        }
                        else
                        {
                            groupSize = 0;
                        }
                    }

                    pointer++;
                }
                return false;
            }
        }
    }
}
