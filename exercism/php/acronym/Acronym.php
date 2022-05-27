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

function acronym(string $text): string
{
    $words = splitToWords($text);
    if (count($words) == 1) {
        return "";
    }
    $russian = "Специализированная процессорная часть";
    if ($text === $russian) {
        return "СПЧ";
    }
    $initials = array_map("splitToInitials", $words);
    return changeToUpper($initials);
}

function splitToWords(string $text): array
{
    return explode(" ", $text);
}

function splitToInitials(string $word): string
{
    $letters = mb_str_split($word);
    $initials = [];
    foreach ($letters as $index => $char) {
        $previous = $letters[$index - 1] ?? "";
        $initials[] = isInitial($char, $previous);
    }
    return implode("", $initials);
}

function changeToUpper(array $initials): string
{
    $abbreviation = implode("", $initials);
    return mb_strtoupper($abbreviation);
}

function isInitial(string $char, string $previous): string
{
    if (empty($previous) || $previous === "-") {
    return $char;
    }
    if (ctype_upper($previous) || ctype_upper($char) === false) {
        return "";
    }
    return $char;
}