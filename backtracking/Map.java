package backtracking;
import java.util.Iterator;
import java.util.ArrayList;

public class Map {

  public static void main(String ... argv) throws Exception {
    System.out.println(color());
  }

  public static ArrayList<ColorMapping> color() {
    ArrayList<ColorMapping> result = new ArrayList<ColorMapping>();
    Vertices vertices = new Vertices();
    while (vertices.hasNext()) {
      ArrayList<Integer> vertex = vertices.next();
      if (result.contains(vertex)) {
        continue;
      }
      Colors colors = new Colors();
      while (colors.hasNext()) {
        Colors.Color color = colors.next();
        boolean failed = false;
        Iterator<ColorMapping> iterator = result.iterator();
        while (!failed) {
          while (iterator.hasNext()) {
            ColorMapping c = iterator.next();
            // If any edges are shared and the color is the same, then fail out.
            if ((c.vertex.contains(vertex.get(0)) || c.vertex.contains(vertex.get(1))) && c.color == color) {
              failed = true;
              break;
            }
          }
          break;
        }
        if (!failed) {
          result.add(new ColorMapping(vertex, color));
          break;
        }
      }
    }
    return result;
  }

}
