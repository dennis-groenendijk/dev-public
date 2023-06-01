<?php

abstract class AchievementType
{
    public function name()
    {
        $class = (new ReflectionClass($this))->getShortName();

        return preg_replace('/[A-Z]/', ' $0', $class);
    }

    public function icon()
    {
        return strtolower(str_replace(' ', '-', $this->name())) . '.png';
    }

    abstract public function qualifier($user);
}

class FirstThousandPoints extends AchievementType
{
    public function qualifier($user)
    {

    }
}

$achievement = new FirstThousandPoints();

echo $achievement->qualifier('user');


