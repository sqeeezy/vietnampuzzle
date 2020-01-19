// I didn't write the Heap algorithm I'm a novice. I did the fracta bit, so kudos to the guy whose code I copied..sqeeezy

// Suppose we have a permutation containing N different elements.
// Heap found a systematic method for choosing at each step a pair of elements to switch,
// in order to produce every possible permutation of these elements exactly once.
// Let us describe Heap's method in a recursive way. First we set a counter i  to 0.
// Now we perform the following steps repeatedly until i  is equal to N.
// We use the algorithm to generate the (N − 1)! permutations of the first N − 1 elements,
// adjoining the last element to each of these. This generates all of the permutations that
// end with the last element. Then if N is odd, we switch the first element and the last one,
// while if N is even we can switch the i th element and the last one (there is no difference between
// N even and odd in the first iteration). We add one to the counter i and repeat. In each iteration,
// the algorithm will produce all of the permutations that end with the element that has just been
// moved to the "last" position. The following pseudocode outputs all permutations of a data array of length N.
// Reference https://en.wikipedia.org/wiki/Heap%27s_algorithm

package main

import "fmt"

//fracta tests to see if the fractional bits of Vietnam puzzle return whole number
//actually I've just made it do the whole equation test..now have to suss out
//how to update the solution number without using global variables
//ok..now I'm thinking I should leave the original Heap as much as possible and redo fracta as a
//sieve called vietnam which takes the output of Heap and prints the soln no. and the soln. Neater that way? ...sqeeezy

func fracta(a []int) int {
	//p, q numerator and denominator of awkward fractional part
	n := 0
	p := 13*a[1]*a[8] + a[2]*a[6]*a[7]
	q := a[2] * a[8]
	//now test to see if p is a multiple of q; -30 to +30 chosen at random
	for k := 1; k < 70; k++ {
		if m := k * q; m == p {
			if a[0]+k+a[3]+12*a[4]-a[5] == 87 {
				n++
				fmt.Println(n, "   ", a)
			} //fmt.Println(k)
		}

	}
	return n
}

// HeapPermutation .... I didn't write this code but here's the oblig comment.
func HeapPermutation(a []int, size int) {
	if size == 1 {
		fracta(a)
		//if a[0]+13*a[1]/a[2]+a[3]+12*a[4]-a[5]+a[6]*a[7]/a[8] == 87

	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	HeapPermutation(a, len(a))
}
