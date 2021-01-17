import java.io.File

fun day06() {
    println("hi from 6")
    val lines = File("06-input.txt").readLines()
//    println('a'.toByte()) //  97
//    println('z'.toByte()) // 122
    val answers = BooleanArray(26) // a-z

    var answerCount = 0
    for (line in lines) {
        if (line.isEmpty()) {
            // println(answerCount)
            // answerCount = 0
            answers.fill(false)
        } else {
            for (char in line) {
                val index = char.toByte() - 'a'.toByte()
                answerCount += if (answers[index]) 0 else 1
                answers[index] = true
            }
        }
    }
    println("Part 1, customs questions answered: $answerCount")

    // Part 2
    val answers2 = IntArray(26) // a-z
    answerCount = 0
    var groupSize = 0
    for (line in lines) {
        if (line.isEmpty()) {
            answers2.forEach { i -> answerCount += if (i == groupSize) 1 else 0 }
//            println(answerCount)
//            answerCount = 0
            answers2.fill(0)
            groupSize = 0
        } else {
            groupSize += 1
            for (char in line) {
                val index = char.toByte() - 'a'.toByte()
                answers2[index] += 1
            }
        }
    }
    answers2.forEach { i -> answerCount += if (i == groupSize) 1 else 0 }

    println("Part 2, customs questions answered: $answerCount")
}