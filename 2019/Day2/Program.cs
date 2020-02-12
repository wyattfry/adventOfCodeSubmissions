using System;

namespace Day2
{
    class Program
    {
        static void Main(string[] args)
        {
            // Part 1
            string text = System.IO.File.ReadAllText(args[0]);
            var intcodeExecutor = new IntcodeExecutor();
            var originalArray = StringToIntArray(text);
            var inputArray = StringToIntArray(text);
            inputArray[1] = 12;
            inputArray[2] = 2;
            var result = intcodeExecutor.Run(inputArray);
            Console.WriteLine($"Value at position 0 after running program: {result[0]}");

            // Part 2
            int targetOutput = 19690720;
            for (int noun = 0; noun <= 99; noun++)
            {
                for (int verb = 0; verb <= 99; verb++)
                {
                    Array.Copy(originalArray, inputArray, originalArray.Length);
                    inputArray[1] = noun;
                    inputArray[2] = verb;
                    if (intcodeExecutor.Run(inputArray)[0] == targetOutput)
                    {
                        Console.WriteLine($"Correct noun-verb found: {noun}, {verb}");
                        return;
                    }
                }
            }
            
        }

        static int[] StringToIntArray(string input)
        {
            return Array.ConvertAll(input.Split(','), s => int.Parse(s));
        }
    }
}
