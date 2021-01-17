import java.io.File
import kotlin.system.exitProcess

fun makeSpacingUniform(input: List<String>): Array<String> {
    var output = mutableListOf<String>()
    var uniformLine = ""
    for (line in input) {
        if (line.isNotEmpty()) {
            if (uniformLine.isNotEmpty()) {
                uniformLine += " $line"
            } else {
                uniformLine = line
            }
        } else if (uniformLine.isNotEmpty()) {
            output.add(uniformLine)
            uniformLine = ""
        }
    }
    output.add(uniformLine)

    return output.toTypedArray()
}

fun isValidCredsPart1(input: String): Boolean {
    if (input.split(" ").size == 8) {
        return true
    }
    if (input.split(" ").size == 7 && !input.contains("cid")) {
        return true
    }
    return false
}

fun isValidCredsPart2(input: String): Boolean {
    val verbose = false
    if (!isValidCredsPart1(input)) {
        return false
    }
    if (!isValidBirthYear(input, verbose)) {
        return false
    }
    if (!isValidIssueYear(input, verbose)) {
        return false
    }
    if (!isValidExpirationYear(input, verbose)) {
        return false
    }
    if (!isValidEyeColor(input, verbose)) {
        return false
    }
    if (!isValidHairColor(input, verbose)) {
        return false
    }
    if (!isValidPassportId(input, verbose)) {
        return false
    }
    if (!isValidHeight(input, verbose)) {
        return false
    }
    return true
}

fun isValidBirthYear(input: String, verbose: Boolean = false): Boolean {
    val year = getValue("byr", input)?.toInt() ?: return false
    val isValid = year != null && year in 1920..2002
    if (verbose) println("$year is valid byr (btw 1920 and 2002: $isValid")
    return isValid
}

fun isValidIssueYear(input: String, verbose: Boolean = false): Boolean {
    val year = getValue("iyr", input)?.toInt() ?: return false
    val isValid = year != null && year in 2010..2020
    if (verbose) println("$year is valid iyr: $isValid")
    return isValid
}

fun isValidExpirationYear(input: String, verbose: Boolean = false): Boolean {
    val year = getValue("eyr", input)?.toInt() ?: return false
    val isValid = year != null && year in 2020..2030
    if (verbose) println("$year is valid eyr: $isValid")
    return isValid
}

fun isValidEyeColor(input: String, verbose: Boolean = false): Boolean {
    val color = getValue("ecl", input) ?: return false
    val validColors = arrayOf("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
    val isValid = validColors.contains(color)
    if (verbose) println("$color is valid ecl: $isValid")
    return isValid
}

fun isValidHairColor(input: String, verbose: Boolean = false): Boolean {
    val color = getValue("hcl", input) ?: return false
    val reg = Regex("#[0-9a-f]{6}")
    val isValid = reg.matches(color)
    if (verbose) println("$color is valid hcl: $isValid")
    return isValid
}

fun isValidPassportId(input: String, verbose: Boolean = false): Boolean {
    val passportId = getValue("pid", input) ?: return false
    val reg = Regex("[0-9]{9}")
    val isValid = reg.matches(passportId)
    if (verbose) println("$passportId is valid pid: $isValid")
    return isValid
}

fun isValidHeight(input: String, verbose: Boolean = false): Boolean {
    val height = getValue("hgt", input) ?: return false
    val reg = Regex("(\\d+)(in|cm)")
    val unit = reg.find(height)?.groupValues?.get(2) ?: return false
    val amount =
        reg.find(height)?.groupValues?.get(1)?.toInt() ?: return false
    var isValid: Boolean
    if (unit == "in") {
        isValid = amount in 59..76
    } else {
        isValid = amount in 150..193
    }
    if (verbose) println("height $height is valid: $isValid")
    return isValid
}

fun getValue(key: String, creds: String): String? {
    val reg = Regex("$key:(\\S+)\\b")
    val result = reg.find(creds)?.groupValues?.get(1) ?: null
    if (result == null) {
        println("Couldn't find key $key in string '$creds'")
    }
    return result
}

fun day04() {
    val ANSI_RESET = "\u001B[0m";
    val ANSI_BLACK = "\u001B[30m";
    val ANSI_RED = "\u001B[31m";
    val ANSI_GREEN = "\u001B[32m";
    val ANSI_YELLOW = "\u001B[33m";
    val ANSI_BLUE = "\u001B[34m";
    val ANSI_PURPLE = "\u001B[35m";
    val ANSI_CYAN = "\u001B[36m";
    val ANSI_WHITE = "\u001B[37m";

//    val lines = File("04-test-input.txt").readLines()
//    val lines = File("04-test-input-part-2.txt").readLines()
    val lines = File("04-input.txt").readLines()
    val spaced = makeSpacingUniform(lines)
    var validCount = 0
    spaced.forEach {
        if (isValidCredsPart1(it)) {
            validCount++
        }
    }
    println("Part 1 Valid passport count: $validCount")

    // Part 2
    validCount = 0
    spaced.forEach {
        // println(it)
        if (isValidCredsPart2(it)) {
            validCount++
            // println("${ANSI_GREEN}VALID $ANSI_RESET")
        } else {
            // println("${ANSI_RED}INVALID $ANSI_RESET")
        }
        // println()
    }
    println("Part 2 Valid passport count: $validCount")
}