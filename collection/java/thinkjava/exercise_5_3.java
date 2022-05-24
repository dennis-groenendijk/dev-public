// Think Java 2 - Exercise 5.3
import java.util.*;

public class Chapter5 {

    public static void checkFermat(int a, int b, int c, int n){
        if ((n > 2) && (Math.pow(a,n) + Math.pow(b,n) == Math.pow(c,n))){
            System.out.println("Holy smokes, Fermat was wrong!");
        } else {
            System.out.println("No, that doesn’t work.");
        }
    }
}