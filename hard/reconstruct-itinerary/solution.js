var findItinerary = function (tickets) {
    const g = {};
    for (const [u, v] of tickets) {
        g[u] = g[u] || [];
        g[u].push(v);
    }
    for (const u in g) {
        g[u].sort();
    }
    const ans = [];
    const dfs = (u) => {
        while (g[u]?.length) {
            const v = g[u].shift();
            dfs(v);
        }
        ans.unshift(u);
    };
    dfs("JFK");
    return ans;
};