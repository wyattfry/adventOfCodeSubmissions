using System;
using System.Linq;

namespace Day1
{
    class Program
    {
        static void Main(string[] args)
        {
            var inputFilePath = args[0];

            string[] lines = System.IO.File.ReadAllLines(inputFilePath);

            int totalFuel = 0;

            var fuelCalculator = new FuelCalculator();

            lines.ToList().ForEach(x => totalFuel += fuelCalculator.GetTotalRequiredFuel( int.Parse(x) ));

            Console.WriteLine($"Total Fuel Required: {totalFuel}");
        }
    }
}
