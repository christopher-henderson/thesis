package nQueens;

public class Between {

  public int bottom;
  public int top;
  public int current;

  Between(int bottom, int top) {
    this.bottom = bottom;
    this.top = top;
    this.current = bottom - 1;
  }

  public boolean hasMore() {
    this.current++;
    return this.current <= this.top;
  }

  public int next() {
    return this.current;
  }
}
