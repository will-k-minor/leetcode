package main

func minHeightShelves(books [][]int, shelfWidth int) int {

	// PROBLEM 1105

	/**
	  The example shown is bad

	  Books cannot be stacked on top of each other on the same shelf

	  sort the books into two maps: width and height

	  start from the tallest books, and then determine the next book to find that is <= height and width fits within the shelfwidth

	  once we fill the shelf, move on to the next highest available book.

	  by grouping all the tallest books together, we can minimize unnecessary shelves

	  each time we fill a row, add the max height in that row to the result
	*/
}
