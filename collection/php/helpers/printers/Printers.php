<?php

/**
 * This file has a collection of some convenient helper functions for printing info/data
 */
class Printers
{
    // printReadableArray prints the provided array in a more readable format
    public function printReadableArray($array)
    {
        echo '<pre>';
        print_r($array);
        echo '<pre>';
    }

    // dd (dump and die) prints the provided variable in a readable format and stops the program
    public function dd($variable)
    {
        echo "<pre>";
        var_dump($variable);
        echo "</pre>";

        die();
    }

    // urlIs returns just the request uri
    public function urlIs($value)
    {
        return $_SERVER['REQUEST_URI'] === $value;
    }
}
