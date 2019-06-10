package missinginteger

/*
 * This is a demo task.
 *
 * Write a function:
 *
 * func Solution(A []int) int
 *
 * that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
 *
 * For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.
 *
 * Given A = [1, 2, 3], the function should return 4.
 *
 * Given A = [−1, −3], the function should return 1.
 *
 * Write an efficient algorithm for the following assumptions:
 *
 * N is an integer within the range [1..100,000];
 * each element of array A is an integer within the range [−1,000,000..1,000,000].
 */
func Solution(A []int) int {
    maxNum := 1000000
	m := map[int]bool{}
	ans := 0

	for _, num := range A {
		if num > 0 {
			m[num] = true
		}
	}

	if len(m) == 0 {
        // All numbers in array A are negative numbers, return 1.
		return 1
	}

	for i := 1; i <= maxNum; i++ {
		if _, ok := m[i]; !ok {
			ans = i
			break
		}
	}

	if ans == 0 {
        // Array A contains all positive numbers within range.
		return maxNum + 1
	}
	return ans
}
