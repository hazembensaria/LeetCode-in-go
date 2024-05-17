package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{-1,-2,-3,-4,-5}, -8))
}

func twoSum(nums []int, target int) []int {
    i:= 0
	j:= len(nums)-1
	
		for nums[i] +  nums[j] != target{
			fmt.Print(nums[i] ,  nums[j])
			if(nums[i]+ nums[j]>target){
				j-- ;
	
			}else if(nums[i]+ nums[j]<target){
				i++ ;
			}
		}
		return []int{i,j};
	
	
}
