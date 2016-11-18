package main

var y int = 0

func increament() int {
	y++
	return y
}

func main() {
	
	/* Example 1 */
	x := 42
	println(x) 
	
	{
		println(x)
		y := 10
		println(y)
	}

	/* Example 2 */
	println(increament()) // 1
	println(increament()) // 2

	/* Example 3 */
	z := 0
	increament := func() int {
		z++
		return z
	}

	println(increament()) // 1
	println(increament()) // 2

	/* Example 4 */
	increament = wrapper()
	println(increament()) // 1
	println(increament()) // 2
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}