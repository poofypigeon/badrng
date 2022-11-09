package badrng

func Rand(n int) int {
	if n < 0 {
		return n
	}

	randMap := map[int]bool{}

	// add all values from zero to n to the map
	for i := 0; i < n; i++ {
		randMap[i] = true
	}

	// get scramblin'
	// remove one value each iteration
	for len(randMap) > 1 {
		counter := 0
		newMap := map[int]bool{}
		for val := range randMap {
			if counter == len(randMap)-1 {
				break
			}

			newMap[val] = true
			counter++
		}

		randMap = newMap
	}

	// use the range construct to grab the key of the remaining value
	for ret := range randMap {
		return ret
	}

	return 0 // make compiler happy
}
