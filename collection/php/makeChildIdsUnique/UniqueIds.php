<?php

/**
 * This code example is an extraction from a class in prod env
 */

class UniqueIds
{
    private array $memory;

    protected function fillMemory(int $min = 0, $max = 9): void
    {
        $this->memory = range($min, $max);
    }

    protected function mapParentId(object $parent): array
    {
        $this->fillMemory();

        return [
            'parent_id' => $parent->id,
            'children' => array_map(
                [$this, 'mapChilderenIds'],
                $parent->children
            ),
        ];
    }

    protected function mapChildrenIds(object $child): string
    {
        return $this->makeChildIdUnique($child->id);
    }

    protected function makeChildIdUnique(string $id): string
    {
        if ($this->memory === []) {
            $this->fillMemory();
        }
        $number = array_shift($this->memory);

        if (in_array($number, $this->memory)) {
            $number = array_shift($this->memory);
        }

        return $id . '_' . $number;
    }
}
