import java.io.File

class Bag constructor(color: String) {
    var contains: ArrayList<Bag> = arrayListOf()
    var containedBy: ArrayList<Bag> = arrayListOf()
    var color: String = color
}

fun traverseContainedBys(startingColor: String, containedByMap: MutableMap<String, MutableSet<String>>, first: Boolean = true): MutableSet<String> {
    val result = mutableSetOf<String>()
    if (!first) {
        result.add(startingColor)
    }
    containedByMap[startingColor]?.forEach { result.addAll(traverseContainedBys(it, containedByMap, false)) }
    return result
}

fun day07() {
    println("Hi from 7")
    val lines = File("07-input.txt").readLines()
    val containedByMap = mutableMapOf<String, MutableSet<String>>()
    for (line in lines) {
        val split = line.split(" ")
        val containerColor = "${split[0]} ${split[1]}"
        if (split[4] == "no") {
            continue
        }
        var idxOffset = 0
        while (idxOffset + 5 < split.size - 1) {
            val bagCount = split[4 + idxOffset].toInt()
            val containedColor = "${split[5 + idxOffset]} ${split[6 + idxOffset]}"
            // println("Count: $bagCount, Color: $containedColor")
            if (containedByMap.containsKey(containedColor)) {
                containedByMap[containedColor]?.add(containerColor)
            } else {
                containedByMap[containedColor] = mutableSetOf(containerColor)
            }
            idxOffset += 4
        }
    }
    val setOfAllContainers = traverseContainedBys("shiny gold", containedByMap)
    println(setOfAllContainers)
    println(setOfAllContainers.size)
}