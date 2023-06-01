<?php

$films = [
    [
        'title' => 'Back to the Future',
        'director' => 'Robert Zemeckis',
        'releaseYear' => 1985,
        'imdb' => 'https://www.imdb.com/title/tt0088763/?ref_=nv_sr_srsg_0_tt_8_nm_0_q_back%2520to%2520',
    ],
    [
        'title' => 'Batman',
        'director' => 'Tim Burton',
        'releaseYear' => 1989,
        'imdb' => 'https://www.imdb.com/title/tt0096895/?ref_=nv_sr_srsg_3_tt_8_nm_0_q_batman',
    ],
    [
        'title' => 'Ghostbusters',
        'director' => 'Ivan Reitman',
        'releaseYear' => 1984,
        'imdb' => 'https://www.imdb.com/title/tt0087332/?ref_=nv_sr_srsg_3_tt_8_nm_0_q_ghostbus',
    ],
];

$filteredFilms = array_filter($films, function ($film) {
    return $film['releaseYear'] === 1984;
});

require "index.view.php";
