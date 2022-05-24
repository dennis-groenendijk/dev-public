// Think Java 2 - Exercise 7.3
public class Chapter7 {

    public static int indexOfMax(int[] a) {
        int l = 0;
        for (int i = a.length - 1; i > 0; i--) {
            if (a[l] < a[i]) {
                l = i;
            }
        }
        return l;
    }
}