package backtracking;

public class Colors {

  public enum Color {
    RED,
    YELLOW,
    GREEN,
    BLUE
  }

  private int pointer;

  Colors() {
    this.pointer = -1;
  }

  public boolean hasNext() {
    return this.pointer + 1 < Color.values().length;
  }

  public Color next() {
    this.pointer++;
    return Color.values()[this.pointer];
  }
}
