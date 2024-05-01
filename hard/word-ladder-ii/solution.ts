function findLadders(beginWord: string, endWord: string, wordList: string[]): string[][] {
    const wordSet = new Set(wordList);
    const result: string[][] = [];
    if (!wordSet.has(endWord)) return result;

    // Helper function to generate intermediate states
    function getNeighbors(word: string): string[] {
        const neighbors: string[] = [];
        for (let i = 0; i < word.length; i++) {
            const originalChar = word[i];
            for (let charCode = 97; charCode <= 122; charCode++) {
                const newChar = String.fromCharCode(charCode);
                if (newChar === originalChar) continue;
                const newWord = word.slice(0, i) + newChar + word.slice(i + 1);
                if (wordSet.has(newWord)) {
                    neighbors.push(newWord);
                }
            }
        }
        return neighbors;
    }

    // BFS to find the shortest path lengths
    const distance = new Map<string, number>();
    const parents = new Map<string, string[]>();
    const queue: string[] = [beginWord];
    distance.set(beginWord, 0);
    parents.set(beginWord, []);
    
    while (queue.length) {
        const word = queue.shift()!;
        const currentDistance = distance.get(word)!;
        for (const neighbor of getNeighbors(word)) {
            if (!distance.has(neighbor)) {
                distance.set(neighbor, currentDistance + 1);
                queue.push(neighbor);
                parents.set(neighbor, [word]);
            } else if (distance.get(neighbor)! === currentDistance + 1) {
                (parents.get(neighbor) || []).push(word);
            }
        }
    }

    // DFS to collect all paths
    function dfs(word: string): string[][] {
        if (word === beginWord) return [[beginWord]];
        const paths: string[][] = [];
        const parentWords = parents.get(word) || [];
        for (const parent of parentWords) {
            for (const path of dfs(parent)) {
                paths.push([...path, word]);
            }
        }
        return paths;
    }

    return dfs(endWord);
}