class Solution {
public:
    vector<vector<bool>> p_hai;
    int n;
    vector<int> dp;
    int solve(int prev,string& s){
        if(prev==n-1){
            return 0;
        }
        if(dp[prev]!=-1){
            return dp[prev];
        }

        int ans=1e9;
        for(int i=prev;i<n;i++){
            if(p_hai[prev][i]){
                if(i==n-1){
                    ans=min(ans,0);
                }else{
                    ans=min(ans,1+solve(i+1,s));
                }
            }
        }

        return dp[prev]= ans;
    }

    int minCut(string s) {
        n=s.length();
    
        p_hai.resize(n,vector<bool>(n,true));
        dp.resize(n+1,-1);

        for(int j=1;j<n;j++){
            for(int i=0;i<j;i++){
                p_hai[i][j]= (s[i]==s[j] and (p_hai[i+1][j-1]));
            }
        }

        return solve(0,s);
    }
};