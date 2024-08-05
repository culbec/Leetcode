#include <list>
#include <unordered_map>
#include <queue>

using namespace std;

class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        if (edges.size() == 0) {
            return {0};
        }

        unordered_map<int, list<int>> adj(n);

        // Completing the adjacency list.
        for (const auto& edge: edges) {
            int u = edge[0], v = edge[1];
            adj[u].push_back(v);
            adj[v].push_back(u);
        }

        vector<int> leaves;

        // Adding the leaves to process.
        for (int i = 0; i < n; i++) {
            if (adj[i].size() == 1) {
                leaves.push_back(i);
            }
        }

        while (n > 2) {
            // Substracting the number of leaves from the number of nodes.
            n -= leaves.size();
            vector<int> newLeaves;

            for(const auto& leaf: leaves) {
                // Removing the connections of the leaf.
                int neighbour = adj[leaf].front();
                adj[neighbour].remove(leaf);

                // If one connection becomes a leaf, we add it for processing.
                if (adj[neighbour].size() == 1) {
                    newLeaves.push_back(neighbour);
                }
            }

            leaves = newLeaves;
        }

        return leaves;
    }
};