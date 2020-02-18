using System;
using System.Diagnostics;

namespace Day4
{
    class Program
    {
        static void Main(string[] args)
        {
            // Advent Of Code 2019, Day 4: Secure Container
            // The elves forgot the password but remember some characteristics
            // How many passwords in the given range satisfy the criteria?

            var watch = Stopwatch.StartNew();

            var passwordRangeInclusive = (min: 402328, max: 864247);

            var validPasswordCount = 0;

            var passwordGenerator = new PasswordGenerator();

            for (int i = passwordRangeInclusive.min; i <= passwordRangeInclusive.max; i++)
            {
                if (passwordGenerator.IsValid(i.ToString()))
                {
                    validPasswordCount++;
                }
            }

            Console.WriteLine($"Number of possible passwords: {validPasswordCount}");

            Console.WriteLine($"Done in {watch.ElapsedMilliseconds}ms");

            watch.Restart();

            validPasswordCount = passwordGenerator.GetValidPasswordCount();

            watch.Stop();

            Console.WriteLine();

            Console.WriteLine($"Number of possible passwords (with recursion): {validPasswordCount}");

            Console.WriteLine($"Done in {watch.ElapsedMilliseconds}ms");

            // Part 2 answer is NOT
            // - 248 (too low)
        }
    }
}
