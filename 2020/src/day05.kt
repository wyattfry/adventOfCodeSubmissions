import java.io.File

fun day05() {
    val airplaneSeats = Array(128) { _ -> IntArray(8) }
    val lines = File("05-input.txt").readLines()
    var highestId = 0
    for (line in lines) {
        val row = getRow(line)
        val column = getColumn(line)
        val id = getId(row, column)
        airplaneSeats[row][column] = id
        highestId = if (id > highestId) id else highestId
    }
    println("Highest id: $highestId")
    println()
    printSeats(airplaneSeats)
}

fun getRow(input: String): Int {
    val totalRows = 128
    var amountToMove = totalRows / 2
    var result = 0
    for (i in 0..6) {
        var x = input[i] // 'F' or 'B'
        if (x == 'B') {
            result += amountToMove
        }
        amountToMove /= 2
    }
    return result
}

fun getColumn(input: String): Int {
    val totalColumns = 8
    var amountToMove = totalColumns / 2
    var result = 0
    for (i in 7..9) {
        var x = input[i] // 'F' or 'B'
        if (x == 'R') {
            result += amountToMove
        }
        amountToMove /= 2
    }
    return result
}

fun getId(row: Int, column: Int): Int {
    return row * 8 + column
}

fun printSeats(seats: Array<IntArray>) {
    /*
Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.
TODO: look for sequence > 0, 0, > 0 to find missing seat ID
     */
    for (row in seats) {
        for (column in row) {
            print("$column\t")
        }
        println()
    }
}