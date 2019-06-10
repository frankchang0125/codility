package brackets

/*
 * A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:
 *
 * S is empty;
 * S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
 * S has the form "VW" where V and W are properly nested strings.
 * For example, the string "{[()()]}" is properly nested but "([)()]" is not.
 *
 * Write a function:
 *
 * func Solution(S string) int
 *
 * that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.
 *
 * For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.
 *
 * Write an efficient algorithm for the following assumptions:
 *
 * N is an integer within the range [0..200,000];
 * string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".
 */
func Solution(S string) int {
	stack := []rune{}

	for _, s := range S {
		if s == '(' || s == '{' || s == '[' {
			stack = append(stack, s) // Push
		} else if s == ')' || s == '}' || s == ']' {
			if len(stack) == 0 {
				return 0
			}

			top := stack[len(stack)-1]

			switch s {
			case ')':
				if top != '(' {
					return 0
				}
			case '}':
				if top != '{' {
					return 0
				}
			case ']':
				if top != '[' {
					return 0
				}
			}

			stack = stack[:len(stack)-1] // Pop
		}
	}

	if len(stack) != 0 {
		return 0
	}
	return 1
}
