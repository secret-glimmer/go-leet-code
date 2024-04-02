function maxNumber(nums1: number[], nums2: number[], k: number): number[] {
    const getMaxSubsequence = (nums: number[], length: number): number[] => {
        const stack: number[] = [];
        let drop = nums.length - length; // Use 'let' here to allow reassignment
        for (const num of nums) {
            while (drop && stack.length && stack[stack.length - 1] < num) {
                stack.pop();
                drop--;
            }
            stack.push(num);
        }
        return stack.slice(0, length);
    };

    const merge = (nums1: number[], nums2: number[]): number[] => {
        const result: number[] = [];
        while (nums1.length || nums2.length) {
            if (compare(nums1, nums2) > 0) {
                result.push(nums1.shift()!);
            } else {
                result.push(nums2.shift()!);
            }
        }
        return result;
    };

    const compare = (a: number[], b: number[]): number => {
        for (let i = 0; i < Math.min(a.length, b.length); i++) {
            if (a[i] !== b[i]) return a[i] - b[i];
        }
        return a.length - b.length;
    };

    let maxResult: number[] = [];
    for (let i = Math.max(0, k - nums2.length); i <= Math.min(k, nums1.length); i++) {
        const subseq1 = getMaxSubsequence(nums1, i);
        const subseq2 = getMaxSubsequence(nums2, k - i);
        const candidate = merge(subseq1, subseq2);
        if (compare(candidate, maxResult) > 0) maxResult = candidate;
    }

    return maxResult;
}