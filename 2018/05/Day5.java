import java.io.File;
import java.io.FileNotFoundException;
import java.util.List;
import java.util.LinkedList;
import java.util.Scanner;

public class Day5 {
    static final int ASCII_A = 65;
    static final int ASCII_Z = 90;
    static final int ASCII_CASE_OFFSET = 32; // A is 65, a is 97

    public static void main(String[] args) {

        String p = getInputFromFile();
        String reacted = reactPolymer(p);
        System.out.println("Count: " + reacted.length());

        // Part 2
        // What is the length of the shortest polymer you can produce by removing
        // all units of exactly one type and fully reacting the result?
        String oneRemoved;
        int shortestLength = Integer.MAX_VALUE;
        for (int i = ASCII_A; i <= ASCII_Z; i++) {
            oneRemoved = reacted.replaceAll(new String(new char[] {(char) i}), "");
            oneRemoved = oneRemoved.replaceAll(new String(new char[] {(char) (i + ASCII_CASE_OFFSET)}), "");
            oneRemoved = reactPolymer(oneRemoved);
            shortestLength = Math.min(oneRemoved.length(), shortestLength);
        }
        System.out.println("Shortest length with one unit removed: " + shortestLength);
    }

    public static boolean sameCharOppositeCase(char a, char b) {
        return (Math.abs((int) a - (int) b) == ASCII_CASE_OFFSET);
    }

    public static String getInputFromFile() {
        Scanner fileScanner = null;
        String units = "";

        try {
            File f = new File("C:/Users/Wyatt/Projects/advent_of_code/05/input");
            // File f = new File("C:/Users/Wyatt/Projects/advent_of_code/05/testinput");
            fileScanner = new Scanner(f);
            while (fileScanner.hasNextLine()) {
                units += fileScanner.nextLine();
            }
            fileScanner.close();
        } catch (FileNotFoundException e) {
            System.out.println("File not found.");
            System.exit(1);
        }
        return units;
    }

    public static String reactPolymer(String p) {
        List<Character> unitsList = new LinkedList<Character>();
        for (int i = 0; i < p.length(); i++) {
            unitsList.add((p.charAt(i)));
        }
        // System.out.println("Starting count: " + unitsList.size());
        int left = 0;
        int right = 1;
        char l;
        char r;
        while (right < unitsList.size()) {
            // upper and lower case letters in ascii differ by 32
            l = unitsList.get(left);
            r = unitsList.get(right);
            if (sameCharOppositeCase(l, r)) {
                unitsList.remove(right);
                unitsList.remove(left);
                if (left != 0) {
                    left--;
                    right--;
                }
            } else {
                left++;
                right++;
            }
        }
        // There's got to be a better way
        // to convert List<Character> to String
        String output = "";
        for (Character c : unitsList) {
            output += c;
        }
        return output;
    }
}