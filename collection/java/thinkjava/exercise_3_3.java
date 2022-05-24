// Think Java 2 - Exercise 3.3
import java.util.Scanner;

public class Convert {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        // prompt the user and get the value
        System.out.print("How many seconds should be converted? ");

        int input = in.nextInt();
        int hours = input / 60;
        int minutes = hours % 60;
        int seconds =  input % 60;

        // convert and output the result
        hours = hours / 60;
        System.out.printf("%d seconds = %d hours, %d minutes and %d seconds\n", input, hours, minutes, seconds);
    }
}