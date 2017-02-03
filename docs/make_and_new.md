# new V.S. make

## make 
- `func make(Type, size IntegerType) Type`
- Make function 用來建立 `slices`、`maps`、`channels`
    - slices: 第二個參數指定長度，此時 size = cap。第三個可選參數指定不同的 cap，cap >= size
    - maps
    - channels
- 回傳 type T (並非 *T) 的初始化值(並非零值)
- 這三個類型的背後引用了使用前必須初始化的資料結構
- Example
    - `make([]int, 10, 100)`: allocate 一個 100 ints 的 array，並且建立一個長度 = 10、容量 = 100 的 slices，指向該 array 的前 10 個元素
    - `var map1 = make(map[string]int)`: 建立一個 map type，key 是 string，value 是 int

## new
- `func new(Type) *Type`
- new function 用來 allocate memory，傳給 new function 的是一個 type，而不是 value。
- 回傳值是指向這個新分配空間零值的`指針`