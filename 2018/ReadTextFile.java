package adventOfCode;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class ReadTextFile {
  public static String[] toStringArray(String filename) {
    List<String> lines = new ArrayList<>();
    Scanner s = null;
    try {
      File f = new File(filename);
      s = new Scanner(f);
      while (s.hasNextLine()) {
        lines.add(s.nextLine());
      }
    } catch (FileNotFoundException e) {
      e.printStackTrace();
    }
    s.close();
    return lines.toArray(new String[lines.size()]);
  }
}