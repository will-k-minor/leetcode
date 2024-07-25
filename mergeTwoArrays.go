package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    /**
        take the max of the two lengths

        loop until we hit the end of the maximum arr
            if we hit the end of one array
                continue to pop the top off the remaining array
            else 
        
                look at the latest val in each array
                pop the one we're about to use
    */
	z, i, k := m + n - 1,  m - 1, n - 1;
  
	for k >= 0 {
	  if (nums1[i] > nums2[k]) {
		nums1[z] = nums1[i];
		i--;
	  } else {
		nums1[z] = nums2[k];
		k--;
	  }
	  z--;
	}

}