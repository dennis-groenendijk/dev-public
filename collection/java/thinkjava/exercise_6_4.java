// Think Java 2 - Exercise 6.4
public class Chapter6 {

    // this method compares if the letters of a word are in alphabetical order
    public static boolean isAbecedarian(String s) {
        int index = 0;
        char c = 'a';
        // this letter by letter comparison is done for as many times as there are letters in a word
        while (index < s.length()) {
            // if the current alphabetical letter c, comes after the provided letter at this index position of s
            // for example, a comes after b resulting in "ba", then return false
            // otherwise it will go to the next letter until you break out of the while loop
            if (c > s.charAt(index)) {
                return false;
            }
            // the current index number determines which alphabetical letter is c in the next loop
            // then the index number itself is increased by 1, to compare the next letter of s
            c = s.charAt(index);
            index = index + 1;
        }
        return true;
    }

    public static void main(String args[]) {
        System.out.println(isAbecedarian("abdest"));
    }
}