package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}

	firstDiff := nums[1] - nums[0]
	sum := firstDiff
	min := firstDiff
	max := firstDiff
	count := 1

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		sum += diff
		count++

		if diff < min {
			min = diff
		}
		if diff > max {
			max = diff
		}
	}

	return Stats{
		Count: count,
		Sum:   sum,
		Min:   min,
		Max:   max,
	}
}
