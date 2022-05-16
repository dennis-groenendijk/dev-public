import java.util.*;

public class fizzbuzz {

    public static void main(String args[]) {
        int n = 100;

        // loop for 100 times
        for (int i = 1; i < n; i++) {

            if (i % 15 == 0) // if the number is divisible by 15 (both 3 & 5), print "FizzBuzz" instead of the number
                System.out.println("FizzBuzz" + " ");

            else if (i % 5 == 0) // if the number is divisible by 5, print "Buzz" instead of the number
                System.out.println("Buzz" + " ");

            else if (i % 3 == 0) // if the number is divisible by 3, print "Fizz" instead of the number
                System.out.println("Fizz" + " ");

            else // print the numbers
                System.out.print(i + " ");
        }
    }
}