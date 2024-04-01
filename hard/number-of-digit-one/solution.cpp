class Solution {
public:
    #define pii pair<long long,long long> 
    vector<vector<pii>> dp;
    pii def ={-1,-1};
    int mod=1e9+7;
    pii solve(string &s,int i,bool tightr){
        int n=s.size();
        int mx=(tightr)?(s[i]-'0'):9;
        if(dp[i][tightr].first!=-1)return dp[i][tightr];
        if(i==n-1){
            if(mx>=1){
                pii res={mx+1,1};
                return dp[i][tightr]= res;
            }else{
                pii res={mx+1,0};
                return dp[i][tightr]=res;
            }
        }
        pii ans={0,0};
        for(int j=0;j<=mx;j++){
            if(j==1){
                auto curr=solve(s,i+1,tightr && j==mx);
                ans.first+=curr.first;
                ans.second+=curr.second;
                ans.second+=curr.first;
                ans.first%=mod;
            }else{
                auto curr=solve(s,i+1,tightr && j==mx);
                ans.first+=curr.first;
                ans.second+=curr.second;
            }
            ans.first%=mod;
            ans.second%=mod;
        }
        return dp[i][tightr]= ans;
    }
    int countDigitOne(int n) {
        string s=to_string(n);
        int size=s.size();
        dp.resize(size,vector<pii>(2,def));
        return (int )solve(s,0,1).second; 
        
    }
//     int countDigitOne(int n)
// {
//     int countr = 0;
//     for (long long i = 1; i <= n; i *= 10) {
//         long long divider = i * 10;
//         countr += (n / divider) * i + min(max(n % divider - i + 1, 0LL), i);
//     }
//     return countr;
// }
};