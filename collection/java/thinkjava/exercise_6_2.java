// Think Java 2 - Exercise 6.2
public class Chapter6 {
    // this method calculates the square root of a (using iterative approximations)
    public static double squareRoot(double a) {
        double x0 = 0; // init value for x0 (will be overridden in the first iteration)
        double x1 = a/2; // initial guess for the square root of a
        while (Math.abs(x1 - x0) > 0.0001) {
            x0 = x1;
            x1 = (x0 + a/x0)/2;
        }
        return x1;
    }

    public static void main(String[] args) {
        System.out.println(squareRoot(36)); // approximately 6
        System.out.println(squareRoot(81)); // approximately 9
    }
}