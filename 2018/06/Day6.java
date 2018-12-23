import java.io.File;
import java.util.List;
import java.util.Set;
import java.util.Map;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Scanner;
import adventOfCode.*;

public class Day6 {
  public static void main(String[] args) {
    //test("/Users/wyattfry/Projects/adventOfCode/06/input_test");
    run("/Users/wyattfry/Projects/adventOfCode/06/input");
  }

  public static void test(String fileName) {
    List<Point> coors = getCoorsFromFile(fileName);
    int dist = getDistance(coors.get(0), coors.get(1));
    System.out.println("expect: 5, result: " + dist);
    dist = getDistance(coors.get(0), coors.get(0));
    System.out.println("expect: 0, result: " + dist);
    dist = getDistance(coors.get(0), coors.get(2));
    System.out.println("expect: 9, result: " + dist);
    dist = getDistance(coors.get(2), coors.get(0));
    System.out.println("expect: 9, result: " + dist);
  }

  public static void run(String fileName) {
    List<Point> coors = getCoorsFromFile(fileName);
    int maxX = 0;
    int maxY = 0;
    for (Point tC : coors) {
      maxX = Math.max(tC.getX(), maxX);
      maxY = Math.max(tC.getY(), maxY);
    }
    // System.out.println(maxX + " " + maxY); // 352, 355
    // We don't need a whole map, just a tuple. If a latter tC is closer, overwrite
    // we only care about the closest one. I hope that holds true for part 2!
    // [
    //   [ [tC, distance], [tC, distance], [tC, distance] ],
    //   [ [tC, distance], [tC, distance], [tC, distance] ],
    //   [ [tC, distance], [tC, distance], [tC, distance] ]
    // ]

    int distBtwTcAndMapPoint;
    int[][][] distances = new int[maxY + 1][maxX + 1][2];
    boolean firstPass = true;
    for (int tC = 0; tC < coors.size(); tC++) {
      for (int y = 0; y < distances.length; y++) {
        for (int x = 0; x < distances[y].length; x++) {
          distBtwTcAndMapPoint = getDistance(coors.get(tC), new Point(x, y));
          if (distBtwTcAndMapPoint < distances[y][x][1] || firstPass) {
            distances[y][x][0] = tC;
            distances[y][x][1] = distBtwTcAndMapPoint;
          } else if (distBtwTcAndMapPoint == distances[y][x][1]) {
            distances[y][x][0] = -1; // signal value of map point equadistant
          }
        }
        if (y == distances.length - 1) {
          firstPass = false;
        }
      }
    }

    // if a tC's area touches an edge, do not count in final calc
    Set<Integer> tCwithInfiniteArea = new HashSet<>();
    for (int y = 0; y < distances.length; y++) {
      for (int x = 0; x < distances[y].length; x++) {
        if (y == 0 || y == distances.length - 1 || x == 0 || x == distances[y].length - 1) {
          if (distances[y][x][0] != -1) {
            tCwithInfiniteArea.add(distances[y][x][0]);
          }
        }
      }
    }
    System.out.println(tCwithInfiniteArea);

    // iterate through distances 2d array and increment the counter for each
    // target coordinate
    int mostSurfaceArea = 0;
    int tCwithMostSA = -1;
    int tC;
    int prevValue;
    Map<Integer, Integer> targetCoordinateSurfaceAreas = new HashMap<>();
    for (int y = 0; y < distances.length; y++) {
      for (int x = 0; x < distances[y].length; x++) {
        tC = distances[y][x][0];
        // System.out.print(tC == -1 ? "." : (char) (tC + 97)); // visualize distance map
        if (tC >= 0) {
          prevValue = targetCoordinateSurfaceAreas.containsKey(tC) ? targetCoordinateSurfaceAreas.get(tC) : 0;
          targetCoordinateSurfaceAreas.put(tC, prevValue + 1);
          if (prevValue + 1 > mostSurfaceArea && !tCwithInfiniteArea.contains(tC)) {
            mostSurfaceArea = prevValue + 1;
            tCwithMostSA = tC;
          }
        }
      }
      // System.out.println(); // visualize distance map
    }

    System.out.println("Most surface area: " + mostSurfaceArea);
    System.out.println("target coordinate: " + tCwithMostSA);

    // -- Part Two --
    // naive solution would be to iterate through each map coordinate,
    // at each mC, get distance to each tC
    // if the sum of distances is <= 10K, increment targetRegionSize by 1
    // terrible time complexity
    // better way might be to go back to storing all distances in the 2d array
    // instead of tuples of closest, or each tC has its own distance matrix
    int sumDistanceBtwTcMc = 0;
    int safeRegionSize = 0;
    for (int y = 0; y < distances.length; y++) {
      for (int x = 0; x < distances[y].length; x++) {
        sumDistanceBtwTcMc = 0;
        for (Point p : coors) {
          sumDistanceBtwTcMc += getDistance(p, new Point(x, y));
        }
        safeRegionSize += sumDistanceBtwTcMc <= 10000 ? 1 : 0;
      }
    }

    System.out.println("Safe region size: " + safeRegionSize);

  
  } // run method

  public static int getDistance(Point a, Point b) {
    int dist = Math.abs(a.getX() - b.getX());
    dist += Math.abs(a.getY() - b.getY());
    return dist;
  }

  public static List<Point> getCoorsFromFile(String fileName) {
    File f = new File(fileName);
    Scanner s = null;
    try {
      s = new Scanner(f);
    } catch (Exception e) {
      e.printStackTrace();
      System.exit(1);
    }
    String[] points;
    List<Point> coors = new ArrayList<Point>();
    while (s.hasNextLine()) {
      points = s.nextLine().split(",");
      coors.add(new Point(
        Integer.parseInt(points[0].trim()),
        Integer.parseInt(points[1].trim())
        ));
    }
    return coors;
  }
}

class Point {
  int x;
  int y;
  public Point(int x, int y) {
    this.x = x;
    this.y = y;
  }
  public Point(int[] xy) {
    this(xy[0], xy[1]);
  }
  public void setX(int x) {
    this.x = x;
  }
  public void setY(int y) {
    this.y = y;
  }
  public int getX() {
    return this.x;
  }
  public int getY() {
    return this.y;
  }
}

/*

each point on the map has a HashMap<Coordinate, Distance>
in the example, coordinate A would add the distance to it
in each point on map, so coor 0, 0's HM would be {A: 2},
1, 0 would be {A: 1} and so on. Once each coor is updated,
then B would add its distances, 0, 0 would become 
{
  A: 2,
  B: 7
}

Once all points have been entered, iterate through each point,
or just keep track of the point with the smallest distance
Use a result map that iterates through each point, pulls out
the coor with smallest distance, and increments the result
map for that point. The ints here represent the surface area
around that point (we count the point itself too),
the surface area of the point with the most is the answer.

It seems like it might be more iterations than necessary,
but it probably works for a naive solution

*/