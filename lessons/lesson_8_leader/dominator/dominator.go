package dominator

/*
 * An array A consisting of N integers is given. The dominator of array A is the value that occurs in more than half of the elements of A.
 *
 * For example, consider array A such that
 *
 *  A[0] = 3    A[1] = 4    A[2] =  3
 *  A[3] = 2    A[4] = 3    A[5] = -1
 *  A[6] = 3    A[7] = 3
 * The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.
 *
 * Write a function
 *
 * func Solution(A []int) int
 *
 * that, given an array A consisting of N integers, returns index of any element of array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.
 *
 * For example, given array A such that
 *
 *  A[0] = 3    A[1] = 4    A[2] =  3
 *  A[3] = 2    A[4] = 3    A[5] = -1
 *  A[6] = 3    A[7] = 3
 * the function may return 0, 2, 4, 6 or 7, as explained above.
 *
 * Write an efficient algorithm for the following assumptions:
 *
 * N is an integer within the range [0..100,000];
 * each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
 */
// Time Complexity: O(n)
// Space Complexity: O(n)
func Solution1(A []int) int {
	m := map[int]int{}
	maxCount := 0
	dominator := -1

	for i := 0; i < len(A); i++ {
		if _, ok := m[A[i]]; !ok {
			m[A[i]] = 0
		}
		m[A[i]]++

		if m[A[i]] > maxCount {
			maxCount = m[A[i]]
			dominator = i
		}
	}

	if maxCount > len(A)/2 {
		return dominator
	}
	return -1
}

// Time Complexity: O(n+n) = O(n)
// Space Complexity: O(1)
func Solution2(A []int) int {
	value := -1
	size := 0

	for i := 0; i < len(A); i++ {
		if size == 0 {
			value = A[i]
			size++
		} else {
			if A[i] != value {
				size--
			} else {
				size++
			}
		}
	}

	if size == 0 {
		return -1
	}

	dominator := -1
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == value {
			dominator = i
			count++
		}
	}

	if count > len(A)/2 {
		return dominator
	}
	return -1
}
