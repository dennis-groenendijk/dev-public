<?php

/**
 * Value Objects are defined by their value
 */

// Example 01

// Avoids primitive obsession - and readability
// Helps with consistency
// Is immutable

class Age
{
    private $age;

    public function __construct($age)
    {
        // is valid age?
        if ($age < 0 || $age > 120) {
            throw new InvalidArgumentException('That makes no sense');
        }

        $this->age = $age;
    }

    // mutable
//    public function increment()
//    {
//        $this->age += 1;
//    }

    // immutable
    public function increment()
    {
        return new self($this->age + 1);
    }

    public function get()
    {
        return $this->age;
    }
}

$age = new Age(35);
$age = $age->increment();

var_dump($age->get());


//function register(string $name, Age $age) {}
//register('John Doe', new Age(35));

// Example 02

class Coordinates
{
    public $x;
    public $y;

    public function __construct($x, $y)
    {
        $this->x = $x;
        $this->y = $y;
    }
}

function pin(Coordinates $coordinates)
{
    $coordinates->x;
    $coordinates->y;
}

function distance(Coordinates $begin, $end)
{

}
