package mesorts

//Author: Ryon McAuley
//Date: 8/14/2020

type Sorting interface {
	//Perform a Selection Sort
	SelectionSort(arr []int)
	QuickSort(arr []int, low int, high int)
}

type Sorter struct {
	Arr []int
	Len int
}

func (s *Sorter) SelectionSort(arr []int) {
	for i := 0; i < s.Len-1; i++ {
		mindex := i
		for j := i + 1; j < s.Len; j++ {
			if arr[j] < arr[mindex] {
				mindex = j
			}
		}
		swap(&arr[mindex], &arr[i])
	}
}

//Helper function for various sorts
func swap(xp *int, yp *int) {
	tmp := *xp
	*xp = *yp
	*yp = tmp
}

func (s *Sorter) QuickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		s.QuickSort(arr, low, pi-1)
		s.QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[high])
	return i + 1
}
