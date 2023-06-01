<?php

/**
 * Domain specific exceptions only = custom exceptions
 */

// Example 01
function add($one, $two)
{
    if (!is_int($one) || !is_int($two)) {
        // take exception
        throw new InvalidArgumentException('Please provide a integer.');
    }

    return $one + $two;
}

try {
    echo add(2, []);
} catch (InvalidArgumentException $e) {
    echo 'Oh well.';
}

// Example 02
class Video
{

}

class User
{
    public function download(Video $video)
    {
        if (!$this->subscribed()) {
            throw new Exception('You must be subscribed to download this video.');
        }
    }

    public function subscribed()
    {
        return false;
    }
}

class UserDownloadsController
{
    public function show()
    {
        try {
            (new User)->download(new Video());
        } catch (Exception $e) {
            echo 'You must be subscribed to download this video.';
        }
    }
}

// Example 03

class TeamException extends Exception
{
    // static constructor
    public static function fromTooManyMembers()
    {
        return new static('You may not add more than 3 team members');
    }
}

class Member
{
    public $name;

    public function __construct($name)
    {
        $this->name = $name;
    }
}

class Team
{
    protected $members = [];

    public function add(Member $member)
    {
        if (count($this->members) === 3) {
            throw TeamException::fromTooManyMembers();
        }

        $this->members[] = $member;
    }
}

class TeamMembersController
{
    public function store()
    {
        $team = new Team; // domain specific requirement - has a max of three members

        try {
            $team->add(new Member('Bruce Wayne'));
            $team->add(new Member('Clark Kent'));
            $team->add(new Member('Diana Prince'));
            $team->add(new Member('Barry Allen'));

            var_dump($team->members());
        } catch (TeamException $e) {
            var_dump($e->getMessage());
        }
    }
}

(new TeamMembersController())->store();


