// Think Java 2 - Exercise 3.4
import java.util.*;

public class Guess {

    public static void main(String[] args) throws InterruptedException {
        // pick a random number
        Random random = new Random();
        int number = random.nextInt(100) + 1;
        int numberGuessed;
        int difference;

        Scanner scan = new Scanner(System.in);

        // prompt for user input
        System.out.println("I'm thinking of a number between 1 and 100\n(including both). Can you guess what it is?");
        System.out.print("Type a number: ");
        numberGuessed = scan.nextInt();
        System.out.println("Your guess is: " + numberGuessed);

        System.out.println("The number I was thinking of is: ");
        Thread.sleep(1000);
        System.out.println("~ drumroll ~");
        Thread.sleep(3000);
        System.out.println(number);

        difference = number - numberGuessed;
        Thread.sleep(1000);
        System.out.printf("You were off by: %d", difference);
    }
}