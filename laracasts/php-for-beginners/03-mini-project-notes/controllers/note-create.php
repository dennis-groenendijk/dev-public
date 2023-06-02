<?php

require 'Validator.php';

$config = require 'config.php';
$db = new Database($config['database']);

$heading = 'Create Note';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $errors = [];

    if (!Validator::string($_POST['title'], 1, 100)) {
        $errors['title'] = 'A title of no more than 100 characters is required';
    }

    if (!Validator::string($_POST['body'], 1, 1000)) {
        $errors['body'] = 'A body of no more than 1000 characters is required';
    }

    if (empty($errors)) {
        $db->query('INSERT INTO `laracasts`.`notes` (`user_id`, `title`, `body`) VALUES (:user_id, :title, :body)', [
            'user_id' => 1,
            'title' => $_POST['title'],
            'body' => $_POST['body'],
        ]);
    }
}

require 'views/note-create.view.php';
