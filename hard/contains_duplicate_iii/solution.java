class Solution {
	public boolean containsNearbyAlmostDuplicate(int[] nums, int indexDiff, int valueDiff) {
		if (indexDiff <= 0 || valueDiff < 0) {
			return false;
		}

		Map<Long, Long> buckets = new HashMap<>();
		long width = (long) valueDiff + 1;

		for (int i = 0; i < nums.length; i++) {
			long bucket = getBucketId(nums[i], width);
			if (buckets.containsKey(bucket) ||
				(buckets.containsKey(bucket - 1) && nums[i] - buckets.get(bucket - 1) <= valueDiff) ||
				(buckets.containsKey(bucket + 1) && buckets.get(bucket + 1) - nums[i] <= valueDiff)) {
				return true;
			}

			buckets.put(bucket, (long) nums[i]);

			if (i >= indexDiff) {
			 	long removeBucket = getBucketId(nums[i - indexDiff], width);
			 	buckets.remove(removeBucket);
			}
		}
		return false;
 	}

 	private long getBucketId(long num, long width) {
	 	return num < 0 ? (num + 1) / width - 1 : num / width;
 	}
}
