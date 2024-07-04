package main

func threeSum(nums []int) [][]int {
	return [][]int{}
	/**
		        always look ahead, never loop back
		        given the combination of two, see if we can find the remainder where
		         k != i && k != j

		        loop through the array, create a map with the number of duplicates found

		        func loop(combo, givenArray, remainderMap)
		            loop through the givenArray
		                newCombo = append(combo, currentValue)
		                remainderMap[currentValue]--;
		                if (newCombo.length > 1) {
		                    remainder = -newCombo.reduce(...);
		                    if remainderMap[remainder] - 1 >= 0 {
		                        // we have found a valid combo!!
		                        result.append(currentCombo)
		                    }
		                } else {
		                    recursively call the function(newCombo, array[1:], remainderMap)
		                    remainderMap[currentVal]++;
		                    }
		                }
		        }

				CLOSE but WRONG!!!

	   /**
	     * Given an array of integers, return all the unique triplets [nums[i], nums[j], nums[k]]
	     * such that they add up to zero.
	*/

	// Sort the array to handle duplicates and facilitate two-pointer approach
	/**
		sort(nums)
		result = []

		// Iterate through the array, treating each element as a potential first element of a triplet
		for i from 0 to length(nums) - 3 {
			// Skip duplicate elements
			if i > 0 and nums[i] == nums[i - 1] continue

			// Two-pointer approach for the remaining part of the array
			left = i + 1
			right = length(nums) - 1

			while left < right:
				sum = nums[i] + nums[left] + nums[right]

				if sum == 0:
					result.append([nums[i], nums[left], nums[right]])
					left++
					right--

					// Skip duplicates for the second and third elements
					while left < right and nums[left] == nums[left - 1] left++
					while left < right and nums[right] == nums[right + 1] right--
				elif sum < 0:
					left++
				else:
					right--
		}

		return result
	}
	*/

}
