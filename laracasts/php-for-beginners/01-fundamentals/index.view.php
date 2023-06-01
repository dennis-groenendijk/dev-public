<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>DVD List</title>
</head>

<body>
    <ul>
        <?php foreach ($filteredFilms as $film) : ?>
            <li>
                <a href="<?= $film['imdb'] ?>">
                    <?= $film['title']; ?> (<?= $film['releaseYear'] ?>) - By <?= $film['director'] ?>
                </a>
            </li>
        <?php endforeach; ?>
    </ul>
</body>
</html>


