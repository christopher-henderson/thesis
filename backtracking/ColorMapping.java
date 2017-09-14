package backtracking;

import java.util.ArrayList;

public class ColorMapping {

  public ArrayList<Integer> vertex;
  public Colors.Color color;

  ColorMapping(ArrayList<Integer> vertex, Colors.Color color) {
    this.vertex = vertex;
    this.color = color;
  }

  public String toString() {
    return "[" + this.vertex.toString() + ", " + this.color.toString() + "]";
  }
}
