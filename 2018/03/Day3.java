import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Day3 {
    public static void main(String[] args) {
        Problem.run();
    }
}

/*

- How many square inches of fabric are within two or more claims?

Approach: make a matrix of the count of fabric sections that on each sq in
Initialize to all 0 first, apply the first claim,
just those sq in are incremented to 1
repeat for next. Any overlap incremented to 2
Once all sections processed, do one last pass and count all the
sq in that are > 1
Do we know the dimensions of the whole fabric?
A naive approach would add a prelim pass to find max dimensions

- input: #1      @ 896,863: 29x19
         claim     x   y    w  h

Hint: the answer is NOT 130452

*/
class Claim {
    int number;
    int x;
    int y;
    int width;
    int height;
    public Claim(String input) {
        number = Integer.parseInt(input.substring(input.indexOf("#") + 1, input.indexOf("@")).trim());
        x = Integer.parseInt(input.substring(input.indexOf("@") + 1, input.indexOf(",")).trim());
        y = Integer.parseInt(input.substring(input.indexOf(",") + 1, input.indexOf(":")).trim());
        width = Integer.parseInt(input.substring(input.indexOf(":") + 2, input.indexOf("x")).trim());
        height = Integer.parseInt(input.substring(input.indexOf("x") + 1).trim());
    }
    public int getNumber() {
        return number;
    }
    public int getX() {return x;}
    public int getY() {return y;}
    public int getWidth() {return width;}
    public int getHeight() {return height;}
}

class Problem {
    public static void run() {
        System.out.println("running");
        List<Claim> claims = getInput();
        int squareInchesOverlap = 0;
        int[] wh = getWidthHeight(claims);
        int[][] sheet = new int[wh[1] + 1][wh[0] + 1];
        for (Claim c : claims) {
            for (int row = c.getY(); row < c.getY() + c.getHeight(); row++) {
                for (int col = c.getX(); col < c.getX() + c.getWidth(); col++) {
                    try {
                        sheet[row][col]++;
                        if (sheet[row][col] == 2) {
                            squareInchesOverlap++;
                        }
                    } catch (ArrayIndexOutOfBoundsException e) {
                        System.out.println("Claim #" + c.getNumber());
                        System.out.printf("Out of bounds: Row %d, Col %d\n", row, col);
                        System.out.printf("Width: %d, Height: %d\n", c.getWidth(), c.getHeight());
                        System.out.printf(sheet.length + "x" + sheet[0].length + "\n");
                        System.exit(1);
                    }
                }
            }
        }

        // Part 2 - find the claim with no overlap
        boolean hasOverlap;
        for (Claim c : claims) {
            hasOverlap = false;
            for (int row = c.getY(); row < c.getY() + c.getHeight(); row++) {
                for (int col = c.getX(); col < c.getX() + c.getWidth(); col++) {
                    try {
                        if (sheet[row][col] > 1) {
                            hasOverlap = true;
                        }
                    } catch (ArrayIndexOutOfBoundsException e) {
                        System.out.printf("Out of bounds: Row %d, Col %d\n", row, col);
                        System.exit(1);
                    }
                }
            }
            if (!hasOverlap) {
                System.out.println("Claim without overlap: " + c.getNumber());
                System.exit(0);
            }
        }
        System.out.println("Square inches of overlap: " + squareInchesOverlap);
    }

    public static List<Claim> getInput() {
        List<Claim> lines = new ArrayList<Claim>();
        Scanner fileScanner = null;
        try {
            fileScanner = new Scanner(new File("C:/Users/Wyatt/Projects/advent_of_code/03/input"));
            while (fileScanner.hasNextLine()) {
                lines.add(new Claim(fileScanner.nextLine()));
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
            System.exit(1);
        }
        fileScanner.close();
        return lines;
    }

    public static int[] getWidthHeight(List<Claim> claims) {
        int height = 0;
        int width = 0;
        Claim claim;
        for (int c = 0; c < claims.size(); c++) {
            claim = claims.get(c);
            height = Math.max(height, claim.getY() + claim.getHeight());
            width = Math.max(width, claim.getX() + claim.getWidth());
        }
        return new int[] {width, height};
    }
}