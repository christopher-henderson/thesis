package nQueens;

import java.util.Arrays;
import java.util.ArrayList;
import java.util.Iterator;

public class Solver {

  public int n;
  public int column;
  public int row;
  public Between rows;
  public ArrayList<ArrayList<Integer>> accumulator;
  // public ArrayList<Integer> currentAnswer;

  Solver(int n, int column, ArrayList<ArrayList<Integer>> accumulator) {
    this.n = n;
    this.column = column;
    System.out.println(this.column);
    this.rows = new Between(1, n);
    this.accumulator = accumulator;
  }

  public boolean hasMore() throws RuntimeException {
    if (this.column > this.n) {
      return false;
    } else if (!this.rows.hasMore()) {
      throw new RuntimeException(this.accumulator.toString());
    }
    int row = this.rows.next();
    // System.out.print(this.column);
    // System.out.print(" ");
    // System.out.println(row);
    if (!this.validate(row)) {
      return true;
    }
    this.accumulator.add(new ArrayList<Integer>(Arrays.asList(new Integer(this.column), new Integer(row))));
    Solver nextSolver = new Solver(this.n, this.column + 1, this.accumulator);
    try {
      while (nextSolver.hasMore()) {

      }
    } catch (RuntimeException e) {
      System.out.println(e);
      this.accumulator.remove(this.accumulator.size() - 1);
      return true;
    }

    return false;
  }

  public boolean validate(int row) {
    Iterator<ArrayList<Integer>> iterator = this.accumulator.iterator();
    while (iterator.hasNext()) {
      if (!this.valid(row, this.column, iterator.next())) {
        return false;
      }
    }
    return true;
  }

  public boolean valid(int row, int column, ArrayList<Integer> CR) {
    int C = CR.get(0).intValue();
    int R = CR.get(1).intValue();
    return row != R &&
            column != C &&
            row + column != R + C &&
            row - column != R - C;
  }

  public ArrayList<ArrayList<Integer>> next() {
    return this.accumulator;
  }

  public void reset() {
    this.accumulator = new ArrayList<ArrayList<Integer>>();
  }

}
