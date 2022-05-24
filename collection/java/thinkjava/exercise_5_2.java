// Think Java 2 - Exercise 5.2
import java.util.*;

public class Chapter5 {

    public static void main(String args[]) {
        // pick a random number
        Random random = new Random();
        int number = random.nextInt(100) + 1;
        int condition = 0;

        while (condition != 3) {
            guess(number);
            condition++;
        }

        if (condition == 3) {
            System.out.println("Game Over! You've guessed wrong too many times.");
        }
    }

    public static void guess(int number) {
        // prompt for user input
        System.out.println("I'm thinking of a number between 1 and 100\n(including both). Can you guess what it is?");
        System.out.print("Type a number: ");

        Scanner scan = new Scanner(System.in);
        int input = scan.nextInt();

        // check if answer is right or wrong
        if (input == number) {
            System.out.println("Congratulations! " + input + " is correct!");
            System.exit(0);
        } else {
            System.out.println(input + " is too " + (input > number ? "high" : "low"));
            System.out.println("Try again\n");
        }
    }
}