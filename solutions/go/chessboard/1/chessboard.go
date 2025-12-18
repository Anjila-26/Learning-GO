package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File[] bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
    f, ok := cb[file]
    if !ok {
        return 0 // file not found
    }
    for _, occupied := range f {
        if occupied {
            count++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the given rank (1-8)
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }

    idx := rank - 1 
    count := 0
    files := []string{"A","B","C","D","E","F","G","H"} // fixed order
    for _, file := range files {
        f, ok := cb[file]
        if !ok {
            continue // skip if the file is missing
        }
        if f[idx] {
            count++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
    for range cb {
        count += 8 // each file has 8 squares (ranks 1-8)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count := 0
    for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                count++
            }
        }
    }
    return count
}
