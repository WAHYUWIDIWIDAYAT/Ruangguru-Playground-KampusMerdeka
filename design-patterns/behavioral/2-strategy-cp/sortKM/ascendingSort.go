package sortKM

// TODO: answer here

//concrete strategy implementation
type AscendingSort struct{}

func (as *AscendingSort) Sort(array []int) {
	//choose any sort algo you want

	//bubble sort
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	
	// TODO: answer here
}
