# Go maps in action

- map type looks like

```
map[KeyType]ValueType
``` 
- `KeyType` 可以是任何 `comparable` 的資料型態
    - Ex：boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types
    - 無法使用：slices, maps, and functions，these types cannot be compared using ==, and may not be used as map keys.
- `ValueType` 可以是任意的型態，甚至是 map
- map 是 `reference type`，如果你用 `var m map[string]int
` 這樣去宣告一個 map，這個 m 會是 `nil`，建議*不要*這樣用
- 用 `make` 來初始化一個 map：`m = make(map[string]int)`
- 和 make 相同的初始化 map 方法：`m = map[string]int{}`
- `make` function 會 allocate 並初始化一個 hash map 的資料結構，並且返回一個指向這個資料結構的 map value
- 如果存取 key 不存在時，會得到該 `ValueType 的 0 值`
- 存取 map 長度：`len()`
- 刪除 map 中的元素：`delete()`
- `delete()` function 不會回傳任何東西，即使該 key 值不存在
- 測試 map 的 key 是否存在：`i, ok := m["route"]`
- iterator map
```
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```
- 初始化 map
```
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
```
- map 不是 thread-safe 的，如果有 concurrent read/write 的需求，請使用 `sync.RWMutex` 來保護你的 map 結構：
```
var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}

counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key:", n)

counter.Lock()
counter.m["some_key"]++
counter.Unlock()
```

- 透過 `range` 來 iterate 一個 map 的 order 是不保證的
- 如果要確保 key 的 order，需要自行實作如下：
```
import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
```