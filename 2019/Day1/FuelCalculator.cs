using System;
using System.Collections.Generic;
using System.Text;

namespace Day1
{
    public class FuelCalculator
    {
        public int GetRequiredFuel(int mass)
        {
            var fuelRequired = (int)Math.Floor(mass / 3m) - 2;

            return fuelRequired < 0 ? 0 : fuelRequired;
        }

        public int GetTotalRequiredFuel(int mass)
        {
            int total = 0;

            int partialFuel = 0;

            do
            {
                partialFuel = this.GetRequiredFuel(mass);
                total += partialFuel;
                mass = partialFuel;
            }

            while (partialFuel > 0);

            return total;
        }
    }
}
