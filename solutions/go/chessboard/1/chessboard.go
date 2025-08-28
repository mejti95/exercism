package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0

    for _, x := range cb[file] {
        if x {
            occupiedSquares += 1
         }
    }
    return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    
	occupiedSquares := 0
    for _, v := range cb {
        if rank > len(v) || rank <= 0 {
            return 0
        }
        if v[rank - 1] {
            occupiedSquares += 1
        }
    }
    return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0

    for _, v := range cb {
        totalSquares += len(v)
        }
    return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0

    for _, v := range cb {
        for _, x := range v {
            	if x {
                    occupiedSquares += 1
                }
        	}
        }
    return occupiedSquares
}
