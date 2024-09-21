class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        vis = []

        # Counting the number of ones
        count = 0
        for i in range(len(nums)):
            
            if nums[i]==0:
                #when there is zero start the counting again
                vis.append(count)
                count = 0
            else:
                count += 1

        #remember to check the edge cases 
        if count is not 0:
            vis.append(count)
        if len(vis) is 1:
            return len(nums)-1
        sum = 0
        for i in range(len(vis)-1):
            sum = max(sum,vis[i] + vis[i+1])
        
        return sum
        