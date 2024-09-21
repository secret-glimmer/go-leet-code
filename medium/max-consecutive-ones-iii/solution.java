class Solution {
    public int longestOnes(int[] nums, int k) {
   int zero = 0,j=0,size=0;
        

        for (int i = 0; i < nums.length; i++) {

            if (nums[i] == 0)
                zero++;

            while (zero > k) {

                if (nums[j] == 0) {
                    zero--;
                }
                j++;

            }

            size = Math.max(size, i-j+1);

        }
        return size;
  
    }
}
