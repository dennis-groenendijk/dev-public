<?php

/*
 * By adding type hints and enabling strict type checking, code can become
 * easier to read, self-documenting and reduce the number of potential bugs.
 * By default, type declarations are non-strict, which means they will attempt
 * to change the original type to match the type specified by the
 * type-declaration.
 *
 * In other words, if you pass a string to a function requiring a float,
 * it will attempt to convert the string value to a float.
 *
 * To enable strict mode, a single declare directive must be placed at the top
 * of the file.
 * This means that the strictness of typing is configured on a per-file basis.
 * This directive not only affects the type declarations of parameters, but also
 * a function's return type.
 *
 * For more info review the Concept on strict type checking in the PHP track
 * <link>.
 *
 * To disable strict typing, comment out the directive below.
 */

declare(strict_types=1);

class Robot
{
    /**
     *
     * @var int[]
     */
    public $position;

    /**
     *
     * @var string
     */
    public $direction;

    public const DIRECTION_NORTH    = 'N';
    public const DIRECTION_EAST     = 'E';
    public const DIRECTION_SOUTH    = 'S';
    public const DIRECTION_WEST     = 'W';

    public function __construct(array $position, string $direction)
    {
        $this->position = $position;
        $this->direction = $direction;
    }

    public function turnRight(): self
    {
        switch ($this->direction) {
            case self::DIRECTION_NORTH:
                $this->direction = self::DIRECTION_EAST;
                break;
            case self::DIRECTION_EAST:
                $this->direction = self::DIRECTION_SOUTH;
                break;
            case self::DIRECTION_SOUTH:
                $this->direction = self::DIRECTION_WEST;
                break;
            case self::DIRECTION_WEST:
                $this->direction = self::DIRECTION_NORTH;
                break;
        }
        return $this;
    }

    public function turnLeft(): self
    {
        switch ($this->direction) {
            case self::DIRECTION_NORTH:
                $this->direction = self::DIRECTION_WEST;
                break;
            case self::DIRECTION_WEST:
                $this->direction = self::DIRECTION_SOUTH;
                break;
            case self::DIRECTION_SOUTH:
                $this->direction = self::DIRECTION_EAST;
                break;
            case self::DIRECTION_EAST:
                $this->direction = self::DIRECTION_NORTH;
                break;
        }
        return $this;
    }

    public function advance(): self
    {
        switch ($this->direction) {
            case self::DIRECTION_NORTH:
                $this->position[1]++;
                break;
            case self::DIRECTION_EAST:
                $this->position[0]++;
                break;
            case self::DIRECTION_SOUTH:
                $this->position[1]--;
                break;
            case self::DIRECTION_WEST:
                $this->position[0]--;
                break;
        }
        return $this;
    }

    public function instructions($input) {
        $characters = str_split($input);
        foreach ($characters as $char) {
            switch ($char) {
                case 'L':
                    $this->turnLeft();
                    break;
                case 'R':
                    $this->turnRight();
                    break;
                case 'A':
                    $this->advance();
                    break;
                default:
                    throw new InvalidArgumentException('Use "L", "R" or "A" to determine the direction of the robot.');
            }
        }
    }

    public function location()
    {
        var_dump($this->position);
        echo($this->direction);
    }
}
