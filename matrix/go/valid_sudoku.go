// https://leetcode.com/problems/valid-sudoku/description/?envType=problem-list-v2&envId=matrix
// 36

func isValidSudoku(b [][]byte) bool {
    for i:=0; i < 9; i++ {
        m := make(map[byte]bool)
         for j:=0; j < 9; j++ {  
            if b[i][j] == '.' {
                continue
            }
            if _, ok := m[b[i][j]]; !ok{
                m[b[i][j]] = true
            } else {
                return false
            }
        }
    }
    for i:=0; i < 9; i++ {
        m := make(map[byte]bool)
         for j:=0; j < 9; j++ {  
            if b[j][i] == '.' {
                continue
            }
            if _, ok := m[b[j][i]]; !ok{
                m[b[j][i]] = true
            } else {
                return false
            }
        }
    }

    v := [9][2]int{{1,1}, {1,4} , {1,7},
    {4,1},  {4,4}, {4,7},
    {7,1},  {7,4},  {7,7}}
    for k:=0; k<9; k++ {
        x, y := v[k][0], v[k][1]
        m := make(map[byte]bool)
        for i:=-1; i<2; i++ {
            for j:= -1; j<2; j++ {
                if b[x-i][y-j] == '.' {
                    continue
                }
                if _, ok := m[b[x-i][y-j]]; !ok{
                    m[b[x-i][y-j]] = true
                } else {
                    return false
                }
            }
        }
    }

    return true
}
