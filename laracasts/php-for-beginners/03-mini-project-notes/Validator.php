<?php

class Validator
{
    // as this is a 'pure' function (no reference to other functions) it can be made static
    public static function string($value, $min = 1, $max = INF)
    {
        $value = trim($value);

        return strlen($value) >= $min && strlen($value) <= $max;
    }

    public static function email($value)
    {
        filter_var($value, FILTER_VALIDATE_EMAIL);
    }
}
