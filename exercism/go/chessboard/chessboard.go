package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    occupiedSquares := 0

    for _, value := range cb[rank] {
        if value {
            occupiedSquares++
        }
    }
	return occupiedSquares
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    occupiedSquares := 0

    if file >= 1 && file <= 8 {
        for _, value := range cb {
            if value[file - 1] {
                occupiedSquares++
            }
        }
    }
	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    numOfFiles := 8
    sumOfSquares := 0

    for range cb {
        sumOfSquares += numOfFiles
    }
	return sumOfSquares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    ranks := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    occupiedSquares := 0

    for _, rank := range ranks {
        occupiedSquares += CountInRank(cb, rank)
    }
	return occupiedSquares
}