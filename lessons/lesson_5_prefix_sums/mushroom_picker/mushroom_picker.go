package mushroompicker

/*
 * Problem:
 *  You are given a non-empty, zero-indexed array A of n (1 <= n <= 100,000) integers
 *  a0, a1, . . . , an−1 (0 <= ai <= 1,000). This array represents number of mushrooms growing on the
 *  consecutive spots along a road. You are also given integers k and m (0 <= k, m < n).
 *  A mushroom picker is at spot number k on the road and should perform m moves. In
 *  one move she moves to an adjacent spot. She collects all the mushrooms growing on spots
 *  she visits. The goal is to calculate the maximum number of mushrooms that the mushroom
 *  picker can collect in m moves.
 *  For example, consider array A such that:
 *          A: [2, 3, 7, 5, 1, 3, 9]
 *      Index:  0  1  2  3  4  5  6
 * The mushroom picker starts at spot k = 4 and should perform m = 6 moves. She might
 * move to spots 3, 2, 3, 4, 5, 6 and thereby collect 1 + 5 + 7 + 3 + 9 = 25 mushrooms. This is the
 * maximal number of mushrooms she can collect.
 *
 * Solution O(m^2):
 *  Note that the best strategy is to move in one direction optionally followed
 *  by some moves in the opposite direction. In other words, the mushroom picker should not
 *  change direction more than once. With this observation we can find the simplest solution.
 *  Make the first p = 0, 1, 2, . . . , m moves in one direction, then the next m − p moves in the
 *  opposite direction. This is just a simple simulation of the moves of the mushroom picker
 *  which requires O(m^2) time.
 *
 * Solution O(n+m):
 *  A better approach is to use prefix sums. If we make p moves in one direction, we can calculate the
 *  maximal opposite location of the mushroom picker. The mushroom picker collects all mushrooms between
 *  these extremes. We can calculate the total number of collected mushrooms in constant time by using prefix sums.
 */
func Solution(A []int, k int, m int) int {
    // The best strategy is to move in one direction optionally followed
    // by some moves in the opposite direction.
    // For each iteration, we can calculate the total mushrooms can be picked
    // by using prefix sums.
    // See countTotal()'s comment for details.
    N := len(A)
    p := prefixSums(A)
    result := 0

    // Try to reach to the left position as far as possible.
    // Note: '0' moves is also valid.
    for i := 0; i <= min(m, k); i++ {
        // Calculate the left and right positions can be reached.
        // Make 'i' moves left.
        leftPos := k-i
        // We have to turn back from left postion, which will take '2 * i' moves
        rightPos := min(N-1, max(k, k+(m-2*i))) 
        result = max(result, countTotal(p, leftPos, rightPos))
    }

    // Try to reach to the right position as far as possible.
    // Note: '0' moves is also valid.
    for i := 0; i <= min(m, N-(k+1)); i++ {
        // Calculate the right and left positions can be reached.
        // Make 'i' moves right.
        rightPos := k+i
        // We have to turn back from right postion, which will take '2 * i' moves
        leftPos := max(0, min(k, k-(m-2*i)))
        result = max(result, countTotal(p, leftPos, rightPos))
    }

    return result
}

// Calculate the prefix sums of array A (with '0' head padding).
func prefixSums(A []int) []int {
    result := make([]int, len(A)+1)
    result[0] = 0
    for i := 1; i <= len(A); i++ {
        result[i] = result[i-1] + A[i-1]
    }
    return result
}

// P is the prefix sums of mushrooms,
// countTotal returns the total mushrooms which can be picked
// from position: x to position: y.
//  ex: Mushrooms: [2, 3, 7, 5, 1, 3, 9]
//      => Prefix sums: [0, 2, 5, 12, 17, 18, 21, 30] (with '0' head padding)
//      => Mushrooms can be picked from position: 2 to position: 5
//      =>  = 7 + 5 + 1 + 3 = 16
//      =>  = P[5+1] - P[2] = 21 - 5 = 16
func countTotal(P []int, x int, y int) int {
    return P[y+1] - P[x]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
