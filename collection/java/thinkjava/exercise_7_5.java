// Think Java 2 - Exercise 7.5
public class Chapter7 {

    public static boolean areFactors(int n, int[] a) {
        boolean r = false;
        for (int i = 0; i < a.length - 1; i++) {
            if (a[i] == n) {
                r = true;
                return true;
            }else{
                r = false;
            }
        }
        return r;
    }
}