package main

// Time complexity: O(2 * (n + m)) = O(n + m)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// Knowing that n + m == the length of the final array...
	// I can try to map the elements to the number of elements that they are greater than

	i := 0
	j := 0

	curr_i := 0
	positon_map := make([]int, n+m)

	// Time complexity of the three fors: O(n + m)

	for i < m && j < n {
		if nums1[i] > nums2[j] {
			positon_map[curr_i] = nums2[j]
			curr_i++
			j++
		} else {
			positon_map[curr_i] = nums1[i]
			curr_i++
			i++
		}
	}

	for i < m {
		positon_map[curr_i] = nums1[i]
		curr_i++
		i++
	}

	for j < n {
		positon_map[curr_i] = nums2[j]
		curr_i++
		j++
	}

	// Copying all the elements from position_map to nums1 -> O(n + m)
	copy(nums1, positon_map)
}

func main() {
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
