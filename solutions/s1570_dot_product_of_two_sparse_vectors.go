/*
	https://leetcode.com/problems/dot-product-of-two-sparse-vectors/

	Given two sparse vectors, compute their dot product.

	Implement class SparseVector:

		SparseVector(nums) Initializes the object with the vector nums
		dotProduct(vec) Compute the dot product between the instance of SparseVector
			and vec

	A sparse vector is a vector that has mostly zero values, you should store the
		sparse
	vector efficiently and compute the dot product between two SparseVector.

	Follow up: What if only one of the vectors is sparse?
*/

package solutions

type SparseVector struct {
	data map[int]int
}

// NewSparseVector should call Constructor to pass LeeCode test
func NewSparseVector(nums []int) SparseVector {
	data := make(map[int]int)
	for i, n := range nums {
		if n != 0 {
			data[i] = n
		}
	}

	return SparseVector{data: data}
}

// Return the dotProduct of two sparse vectors
func (sv *SparseVector) dotProduct(vec SparseVector) int {
	sum := 0

	for idx, val := range sv.data {
		if v, ok := vec.data[idx]; ok {
			sum += val * v
		}
	}

	return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
