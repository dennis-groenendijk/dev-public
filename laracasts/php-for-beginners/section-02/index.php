<?php

//require 'router.php'; // uncomment to display website
require 'Database.php';
require 'functions.php';

$config = require('config.php');

$db = new Database($config['database']);

$id = $_GET['id'];
$query = "select * from films where id = :id"; // or use '?' and no assoc arr

$films = $db->query($query, [':id' => $id])->fetchAll();

dd($films);
foreach ($films as $film) {
    echo "<li>" . $film['title'] . "</li>";
}
