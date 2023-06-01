<?php

/**
 * Encapsulation = integrity protection
 */

// Example 01
class Person
{
    public string $name;

    public function __construct($name)
    {
        $this->name = $name;
    }

    public function job()
    {
        return 'Software Developer';
    }

    public function favoriteDish()
    {
        //
    }

    // Wouldn't want to share..
    private function thingsThatKeepMeUpAtNight()
    {
        return 'The worlds will collide!';
    }
}

// kept private
//$john = new Person('John');
//
//var_dump($john->thingsThatKeepMeUpAtNight());

// but still accessible
$method = new \ReflectionMethod(Person::class, 'thingsThatKeepMeUpAtNight');
$method->setAccessible(true);

$bob = new Person('Bob');

$method->invoke($bob);

// Example 02
class TennisMatch
{
    protected $player;

    // add tag

    /** @api */
    public function score()
    {
        // is there a winner
        // does someone have an advantage
        // are they in deuce
    }

    // "getter fn"; often considered bad practice, sometimes needed,
    // but could break encapsulation
    public function player()
    {
        return $this->player;
    }

    protected function hasWinner()
    {
        //
    }

    protected function hasAdvantage()
    {
        //
    }

    protected function inDeuce()
    {
        //
    }
}

