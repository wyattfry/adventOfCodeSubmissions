import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.util.Set;
import java.util.HashSet;

public class AdventOfCode1 {
  public static void main(String[] args) {
    String line;
    Scanner fileScanner = null;
    int sum = 0;
    int number;
    Set<Integer> s = new HashSet<Integer>();

    try {
      fileScanner = new Scanner(new File("1_input"));

    while (fileScanner.hasNextLine()) {
      line = fileScanner.nextLine();
      number = Integer.parseInt(line);
      sum += number; 
      if (s.contains(sum)) {
        System.out.println(sum);
        System.exit(0);
      } else {
        s.add(sum);
      }
      if (!fileScanner.hasNextLine()) {
        fileScanner.close();
        fileScanner = new Scanner(new File("1_input"));
      }
    }
    
    } catch (FileNotFoundException e) {
      System.out.println(e.getMessage());
      System.exit(1);
    }
    fileScanner.close();
    System.out.println("Sum: " + sum);
  }

}
