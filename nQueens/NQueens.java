package nQueens;

import java.util.Iterator;
import java.util.ArrayList;

public class NQueens {

  public static void main(String ... argv) throws Exception {
    // System.out.println(nqueens(8));
    Solver s = new Solver(4, 1, new ArrayList<ArrayList<Integer>>());
    while (s.hasMore()) {
      ArrayList<ArrayList<Integer>> answer = s.next();
      System.out.println(answer);
      s.reset();
    }
    // System.out.println(s.hasMore());
    // System.out.println(s.next());
    // Between b = new Between(1, 7);
    // while (b.hasMore()) {
    //   System.out.println(b.next());
    // }
    // ArrayList<Integer> a = new ArrayList<Integer>();
    // derp(a);
    // System.out.println(a);
  }

  public static void derp(ArrayList<Integer> a) {
    a.add(new Integer(5));
  }

  public static ArrayList<ArrayList<Integer>> nqueens(int n) {
    boolean succeeding = true;
    ArrayList<ArrayList<Integer>> answers = new ArrayList<ArrayList<Integer>>();
    while (succeeding) {
      ArrayList<Integer> result = solve_nqueens(8);
      if (result == null) {
        succeeding = false;
      } else {
        answers.add(result);
      }
    }
    return answers;
  }

  public static ArrayList<Integer> solve_nqueens(int n) {
    int column = 0;

    for (int queen = 0; queen < n; queen++) {

    }
    return null;
  }



}
