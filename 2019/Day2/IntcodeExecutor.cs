using System;
using System.Collections.Generic;
using System.Text;

namespace Day2
{
    public class IntcodeExecutor
    {
        public int[] Run(int[] program)
        {
            Array.Copy(program, program, program.Length);
            for (int i = 0; i < program.Length; i += 4)
            // i = instruction pointer
            {
                if (program[i] == 99)
                {
                    return program;
                }
                // Parameters
                var sourceAddress1 = program[i + 1];
                var sourceAddress2 = program[i + 2];
                var targetAddress = program[i + 3];

                switch (program[i])
                {
                    // Add
                    case 1:
                        program[targetAddress] = program[sourceAddress1] + program[sourceAddress2];
                        break;
                    // Multiply
                    case 2:
                        program[targetAddress] = program[sourceAddress1] * program[sourceAddress2];
                        break;
                    default:
                        throw new Exception($"Invalid opcode {program[i]}");
                }
            }
            throw new Exception("Reached end of program, but last opcode was not 99.");
        }
    }
}
