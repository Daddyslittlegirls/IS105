package main

import "fmt"

func main() {
		list := []int{9, 2, 1, 7, 5, 3, 10, 4, 6, 8}
		benchmarkBSortModified(list)
		fmt.Println("Her er de sorterte tallene:")
		fmt.Println(list)
	}

func benchmarkBSortModified(list []int) {
		sort := true;
		for sort {
			sort = false
			for i := 0; i < len(list) - 1; i++ {
				if list[i + 1] < list[i] {
					swap(list, i, i + 1)
					sort = true
				}
			}
		}
	}

func swap(list []int, i, j int) {
	tmp := list[j]
	list[j] = list[i]
	list[i] = tmp
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
