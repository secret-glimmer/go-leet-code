class Solution {
    public int maxOperations(int[] nums, int k) {
        int count = 0;
        for(int i = 0; i < nums.length; i++){
            if(nums[i] == -1) continue;
            for(int j = i + 1; j < nums.length; j++){
                if(nums[j] == -1) continue;
                if(nums[i] + nums[j] == k){
                    count++;
                    nums[i] = -1;
                    nums[j] = -1;
                    break;
                }
            }
        }
        return count;
    }
}