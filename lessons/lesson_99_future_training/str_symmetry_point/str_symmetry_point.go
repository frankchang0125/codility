package strsymmetrypoint

/*
 * Write a function:
 *
 * func Solution(S string) int
 *
 * that, given a string S, returns the index (counting from 0) of a character such that the part of the string to the left of that character is a reversal of the part of the string to its right. The function should return −1 if no such index exists.
 *
 * Note: reversing an empty string (i.e. a string whose length is zero) gives an empty string.
 *
 * For example, given a string:
 *
 * "racecar"
 *
 * the function should return 3, because the substring to the left of the character "e" at index 3 is "rac", and the one to the right is "car".
 *
 * Given a string:
 *
 * "x"
 *
 * the function should return 0, because both substrings are empty.
 *
 * Write an efficient algorithm for the following assumptions:
 *
 * the length of S is within the range [0..2,000,000].
 */
func Solution(S string) int {
	if len(S)&1 == 0 {
		return -1
	} else if len(S) == 1 {
		return 0
	}

	l := 0
	r := len(S) - 1

	for l < r {
		if S[l] != S[r] {
			break
		}

		l++
		r--
	}

	if l == r {
		return l
	}
	return -1
}
