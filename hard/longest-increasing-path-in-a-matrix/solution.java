class Solution {

    static int arr[][];
    static int n;
    static int m;
    static int dir[][] = {{1,0},{-1,0},{0,1},{0,-1}};

    static int dp[][];

    static int helper(int r, int c){

        if(dp[r][c]!=-1) return dp[r][c];
        int ans = 1;
        for(int k=0; k<4; k++){
            int nr = r+dir[k][0];
            int nc = c+dir[k][1];

            if(nr>=0 && nc>=0 && nr<n && nc<m && arr[nr][nc]>arr[r][c]){
                ans = Math.max(ans,1+helper(nr,nc));
            }
        }
        return dp[r][c] = ans;
    }
    public int longestIncreasingPath(int[][] matrix) {
        arr = matrix;
        n = arr.length;
        m = arr[0].length;
        dp = new int[205][205];
        for(var a : dp) Arrays.fill(a,-1);

        int ans = -1;
        for(int i=0; i<n; i++){
            for(int j=0; j<m; j++){
                ans = Math.max(ans,helper(i,j));
            }
        }
        return ans;
    }
}