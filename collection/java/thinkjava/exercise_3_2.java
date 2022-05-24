// Think Java 2 - Exercise 3.2
import java.util.Scanner;

public class Convert {

    public static void main(String[] args) {
        double C, F;
        Scanner in = new Scanner(System.in);

        // prompt the user and get the value
        System.out.print("What is the temperature in Celsius? ");
        C = in.nextDouble();

        // convert and output the result
        F = C * 1.8 + 32.0;
        System.out.printf("%.1f C = %.1f F\n", C, F);
    }
}