package divconq

// Maximum subarray problem is the task of finding the contiguous subarray within a one-dimensional array of numbers
// which has the largest sum.
// Time complexity: W: O(nlg(n)) B: O(nlg(n)) A: O(nlg(n))
func MaxSubarray(a []int) (crossLow, crossHigh, crossSum int) {
	return FindMaxSubarray(a, 0, len(a)-1)
}

func FindMaxSubarray(a []int, low, high int) (crossLow, crossHigh, crossSum int) {
	if high == low {
		return low, high, a[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaxSubarray(a, low, mid)
		rightLow, rightHigh, rightSum := FindMaxSubarray(a, mid+1, high)
		crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(a, low, mid, high)

		switch {
		case leftSum >= rightSum && leftSum >= crossSum:
			return leftLow, leftHigh, leftSum
		case rightSum >= leftSum && rightSum >= crossSum:
			return rightLow, rightHigh, rightSum
		default:
			return crossLow, crossHigh, crossSum
		}
	}
}

// Find a maximum subarray crossing the midpoint time.
// Time complexity: W: O(n) B: O(n) A: O(n)
func FindMaxCrossingSubarray(a []int, low, mid, high int) (maxLeft, maxRight, maxSum int) {
	leftSum := a[0]
	sum := 0
	for i := mid; i > low; i-- {
		sum += a[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := a[len(a)-1]
	sum = 0
	for i := mid + 1; i < len(a); i++ {
		sum += a[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}
