// Think Java 2 - Exercise 7.1

// banana takes an array of integers and returns a list of multiplication results, based on each index position in array "a" and its corresponding element.
public static int banana(int[] a) {
        int kiwi = 1; // here, kiwi is the initial state of this variable
        // as the array receives input, the initial state of "kiwi" can't be 0, therefore it's set to 1
        int i = 0; // "i" is the index position of an element in array "a", in this case it starts at the first position

        // the condition states that, as long as you haven't reached the last index position of all the elements in array "a"
        // continue to calculate "kiwi"
        // then "i" is incremented by 1 and the loop starts over
        while (i < a.length) {
        kiwi = kiwi * a[i]; // here, the current state of "kiwi" will be multiplied by the element in array "a", at index position "i"
        // resulting in a new state of "kiwi"
        i++; // "i" is incremented by 1, so the follow-up index position of array "a" is used in the next iteration of this loop
        }
        return kiwi; // here, the new state of "kiwi" is returned to use in the next iteration of this loop
        }

// grapefruit finds the index position of "grape" in the array "a", but returns "-1" if it isn't found.
public static int grapefruit(int[] a, int grape) {
        for (int i = 0; i < a.length; i++) { // the initial state of "i" is set to 0 and as long as "i" is less than
        // the number of elements in array "a", each new loop will increment "i" by 1
        if (a[i] == grape) { // when the index position of "grape" is found, it returns this position as the result
        return i;
        }
        }
        return -1;
        }

// pineapple counts how many times "apple" appears in the array "a".
public static int pineapple(int[] a, int apple) {
        int pear = 0; // here, "pear" is the initial state of the result it will return after below for loop is done
        for (int pine: a) { // is enhanced for loop; for each integer "pine" in array "a", iterating each index of the array
        if (pine == apple) {
        pear++; // if above condition is true, pear is incremented by 1
        }
        }
        return pear; // when the for loop is done and every time "apple" appears in array "a" is counted, the result is returned
        }