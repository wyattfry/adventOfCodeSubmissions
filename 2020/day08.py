
def getAccValueOfInfLoop(lines):
    accumulator = 0
    index = 0
    lineCount = len(lines)
    linesVisited = [False for i in range(lineCount)]
    while index <= lineCount:
        if linesVisited[index]:
            print("Part 1: Line {line} already executed. Accumulator: {acc}".format(line=index, acc=accumulator))
            break
        linesVisited[index] = True
        line = lines[index]
        instruction = line.split(" ")[0]
        argument = int(line.split(" ")[1])
        if instruction == "acc":
            accumulator += argument
        if instruction == "jmp":
            index += argument
        else:
            # instruction == nop or acc
            index += 1
    
# linesFromFile = open("08-test-input.txt").readlines()
# linesFromFile = open("08-input.txt").readlines()
# getAccValueOfInfLoop(linesFromFile)

# Part 2

# Somewhere in the program, either a jmp is supposed to be a nop, or a nop is supposed to be a jmp. (No acc instructions were harmed in the corruption of this boot code.)

# The program is supposed to terminate by attempting to execute an instruction immediately after the last instruction in the file. By changing exactly one jmp or nop, you can repair the boot code and make it terminate correctly.

def handleCorruptedInstructions(lines, accumulator = 0, index = 0):
    lineCount = len(lines)
    if index == lineCount:
        print("Part 2 answer: {}".format(accumulator))
        return
    if index >= lineCount:
        return
    while index <= lineCount:
        if index == lineCount:
            print("Part 2 answer: {}".format(accumulator))
            return
        if index > lineCount:
            return
        if lines[index] == "DONE":
            # print("Part 1: Line {line} already executed. Accumulator: {acc}".format(line=index, acc=accumulator))
            # break
            return
        line = lines[index]
        instruction = line.split(" ")[0]
        argument = int(line.split(" ")[1])
        lines[index] = "DONE"
        if instruction == "acc":
            accumulator += argument
            index += 1
            continue
        if instruction == "jmp":
            # try nop instead
            handleCorruptedInstructions(lines.copy(), accumulator, index + 1)
            index += argument
            continue
        if instruction == "nop":
            # try jmp instead
            handleCorruptedInstructions(lines.copy(), accumulator, index + argument)
            index += 1
            continue
            
# linesFromFile = open("08-test-input.txt").readlines()
linesFromFile = open("08-test-input.txt").readlines()
handleCorruptedInstructions(linesFromFile)
