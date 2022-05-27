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

class Tournament
{
    public const HEADING = 'Team                           | MP |  W |  D |  L |  P';
    private array $teams = [];

    public function __construct()
    {
    }

    public function tally($scores): string
    {
        $table = [self::HEADING];

        $results = array_values(array_filter(explode("\n", $scores)));
        foreach ($results as $result) {
            [$teamA, $teamB, $outcome] = explode(';', $result);
            $this->addTeam($teamA);
            $this->addTeam($teamB);
            $this->addResult($teamA, $teamB, $outcome);
        }
        $this->sortTeams();

        $table = array_merge($table, $this->printTeams());
        return implode("\n", $table);
    }

    protected function addTeam(string $team): void
    {
        if (array_search($team, array_column($this->teams, 'Team')) === false) {
            $this->teams[] = ['Team' => $team, 'W' => 0, 'D' => 0, 'L' => 0];
        }
    }

    protected function addResult(string $teamA, string $teamB, string $outcome): void
    {
        $A = array_search($teamA, array_column($this->teams, 'Team'));
        $B = array_search($teamB, array_column($this->teams, 'Team'));
        switch ($outcome) {
            case 'W':
                $this->teams[$A]['W']++;
                $this->teams[$B]['L']++;
                break;
            case 'D':
                $this->teams[$A]['D']++;
                $this->teams[$B]['D']++;
                break;
            case 'L':
                $this->teams[$A]['L']++;
                $this->teams[$B]['W']++;
                break;
        }
    }

    protected function sortTeams(): void
    {
        usort($this->teams, function ($A, $B) {
            $pointsA = 3 * $A['W'] + $A['D'];
            $pointsB = 3 * $B['W'] + $B['D'];

            if ($pointsA == $pointsB) {
                return $A['Team'] < $B['Team'] ? -1 : 1;
            }
            return $pointsA > $pointsB ? -1 : 1;
        });
    }

    protected function printTeams(): array
    {
        return array_map(
            fn ($team) => sprintf(
                "%-30s | %2d | %2d | %2d | %2d | %2d",
                $team['Team'],
                array_sum($team),
                $team['W'],
                $team['D'],
                $team['L'],
                3 * $team['W'] + $team['D']
            ),
            $this->teams
        );
    }
}
