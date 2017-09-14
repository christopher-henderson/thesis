package backtracking;
import java.util.ArrayList;
import java.util.Arrays;



public class Vertices {

  private static final ArrayList<ArrayList<Integer>> data = new ArrayList<ArrayList<Integer>>(
    Arrays.asList(
      new ArrayList<Integer>(Arrays.asList(new Integer(1),new Integer(2))),
      new ArrayList<Integer>(Arrays.asList(new Integer(1),new Integer(3))),
      new ArrayList<Integer>(Arrays.asList(new Integer(1),new Integer(6))),
      new ArrayList<Integer>(Arrays.asList(new Integer(1),new Integer(4))),

      new ArrayList<Integer>(Arrays.asList(new Integer(2),new Integer(1))),
      new ArrayList<Integer>(Arrays.asList(new Integer(2),new Integer(3))),
      new ArrayList<Integer>(Arrays.asList(new Integer(2),new Integer(5))),

      new ArrayList<Integer>(Arrays.asList(new Integer(3),new Integer(2))),
      new ArrayList<Integer>(Arrays.asList(new Integer(3),new Integer(1))),
      new ArrayList<Integer>(Arrays.asList(new Integer(3),new Integer(6))),
      new ArrayList<Integer>(Arrays.asList(new Integer(3),new Integer(4))),

      new ArrayList<Integer>(Arrays.asList(new Integer(4),new Integer(1))),
      new ArrayList<Integer>(Arrays.asList(new Integer(4),new Integer(3))),
      new ArrayList<Integer>(Arrays.asList(new Integer(4),new Integer(5))),
      new ArrayList<Integer>(Arrays.asList(new Integer(4),new Integer(6))),

      new ArrayList<Integer>(Arrays.asList(new Integer(5),new Integer(2))),
      new ArrayList<Integer>(Arrays.asList(new Integer(5),new Integer(3))),
      new ArrayList<Integer>(Arrays.asList(new Integer(5),new Integer(4))),

      new ArrayList<Integer>(Arrays.asList(new Integer(6),new Integer(3))),
      new ArrayList<Integer>(Arrays.asList(new Integer(6),new Integer(1))),
      new ArrayList<Integer>(Arrays.asList(new Integer(6),new Integer(4)))
    )
  );

  private int pointer;

  Vertices() {
    this.pointer = -1;
  }

  public boolean hasNext() {
    return this.pointer + 1 < data.size();
  }

  public ArrayList<Integer> next() {
    this.pointer++;
    return data.get(this.pointer);
  }

}
