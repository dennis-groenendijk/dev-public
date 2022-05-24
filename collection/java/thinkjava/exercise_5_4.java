// Think Java 2 - Exercise 5.3
boolean yes = true
        boolean no = false
        int loVal = -999
        int hiVal = 999
        double grade = 87.5
        double amount = 50.0
        String hello = "world"

        Expression
        yes == no || grade > amount         true
        amount == 40.0 || 50.0              error
        hiVal != loVal || loVal < 0         true
        True || hello.length() > 0          error
        hello.isEmpty() && yes              false
        grade <= 100 && !false              true
        !yes || no                          false
        grade > 75 > amount                 error
        amount <= hiVal && amount >= loVal  true
        no && !no || yes && !yes            false