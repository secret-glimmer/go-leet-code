class Solution {
public:

    double frogPosition(int n, vector<vector<int>>& edges, int t, int target) {
        vector<vector<int>>graph(n+1);
        vector<bool>vi(n+1, false);
       

        for(int i=0;i<edges.size();i++){
            int u=edges[i][0];
            int v=edges[i][1];

            graph[u].push_back(v);
            graph[v].push_back(u);
        }
        graph[1].push_back(0);
        queue<pair<int ,double> > q;

        q.push({1,1.00});
        vi[1]=true;
        int level=0;

        while(!q.empty()){
            int len=q.size();

            while(len--){
                pair<int,double> currentNode=q.front();
                q.pop();
                int currentNodeIdx=currentNode.first;
                double currentNodeProb=currentNode.second;
                

                if(currentNodeIdx == target){

                    if(graph[currentNodeIdx].size()==1 && t>=level){
                        return currentNodeProb;
                    }
                    else if(level == t){
                        return currentNodeProb;
                    }
                    else{
                        return 0.00;
                    }
                }

                
                     int s= graph[currentNodeIdx].size()-1 ;
                   

                for(int i=0;i<graph[currentNodeIdx].size();i++){
                    int child=graph[currentNodeIdx][i];
                   
                   
                    double prob= (double)currentNodeProb /s ;
                    if(!vi[child] && child>0){
                        q.push({child, prob}); 
                        vi[child]=true;
                    }
                }

            }
            level++;
           
        }
        return 0.00;
    }
};