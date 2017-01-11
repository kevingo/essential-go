# Slice and Array in Go

## Array
- 宣告 array 需要有長度和型別
- array 的 size 是固定的
- array 不需要明確的宣告初始值，他會被預設初始為對應的零值
- 一個 array 的值代表整個 array，不是 pointer
- 當你 pass array 的值傳遞出去，代表產生了 copy of its content
- 如果想要傳遞原始 array，可以用 pointer
- 宣告 array
    - `var a [4]int`
    - `b := [2]string{"Penn", "Teller"}`
    - `b := [...]string{"Penn", "Teller"}`

## Slice
- slice 不需要宣告長度
- slice 的宣告跟 array 很類似，只是你不需要宣告長度
    - `b := []string{"Penn", "Teller"}`
- slice 可以用 `make` function 來宣告
    - `s = make([]byte, 5, 5)`
    - `// s == []byte{0, 0, 0, 0, 0}`
- slice 的 zero-value 是 nil
- 
