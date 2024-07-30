package main;

func minimumDeletions(s string) int {
    /**
        Keep track fo a's and b's
        compare the numbers of a's and b's.

        each a we run into going up, ignore
            until we hit any b's
        Miminum number of chars that would need to be removed begins
        to be counted in from that.
        We keep the minimum each time. So if we run into a bunch of b's, we oviously want to go with a.
        If there's a bunch of a's, then you want to go with the B's.
        
    */
   	a := 0
	b := 0

	for _, char := range s {
		if char == 'a' {
			a = min(a+1, b)
		} else {
			b += 1
		}
	}

	return min(a, b) 
}