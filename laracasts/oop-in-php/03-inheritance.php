<?php

// Example 01
class CoffeeMaker
{
    public function brew()
    {
        var_dump('Brewing the coffee');
    }
}

// "is a"
class SpecialtyCoffeeMaker extends CoffeeMaker
{
    public function brewLatte()
    {
        var_dump('Brewing a latte');
    }
}

(new SpecialtyCoffeeMaker())->brew();

// Example 02
class Collection
{
    protected array $items;

    public function __construct(array $items)
    {
        $this->items = $items;
    }

    public function sum($key)
    {
        // initial code
//        return array_sum(array_map(function ($item) use ($key) {
//            return $item->$key;
//        }, $this->items));

        // abbreviated
//        return array_sum(array_map(fn($item) => $item->$key, $this->items));

        // even cleaner!
        return array_sum(array_column($this->items, $key));
    }
}

class VideosCollection extends Collection
{
    public function length()
    {
        return $this->sum('length');
    }
}

class Video
{
    public $title;
    public $length;

    public function __construct($title, $length)
    {
        $this->title = $title;
        $this->length = $length;
    }
}

$videos = new VideosCollection([
    new Video('Video 01', 100),
    new Video('Video 02', 200),
    new Video('Video 03', 300),
]);

echo $videos->length();



