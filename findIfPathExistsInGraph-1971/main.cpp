#include <vector>
#include <unordered_map>
#include <exception>
#include <utility>
#include <iostream>
using namespace std;

bool dfs(vector<vector<int>> &adj, vector<bool>& visited, int source, int destination)
{
    // Base case: we've reached the destination.
    if (source == destination)
    {
        return true;
    }

    visited[source] = true;
    for (size_t i = 0; i < adj.at(source).size(); i++)
    {
        // Checking if a neighbour has already been visited in the current direction.
        if (!visited.at(adj[source][i])) 
        {
            if(dfs(adj, visited, adj.at(source).at(i), destination)) {
                return true;
            }
        }
    }

    return false;
}

bool validPath(int n, vector<vector<int>> &edges, int source, int destination)
{
    vector<vector<int>> adj(n, vector<int>());
    vector<bool> visited(n, false);

    // Adding the neighbours in a map.
    for (int i = 0; i < (int) edges.size(); i++)
    {
        adj.at(edges[i][0]).push_back(edges[i][1]);    
        adj.at(edges[i][1]).push_back(edges[i][0]);   
    }
    return dfs(adj, visited, source, destination);
}

int main()
{
    vector<vector<int>> e1 = {{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}};
    int n1 = 6, s1 = 0, d1 = 5;
    cout << validPath(n1, e1, s1, d1) << "\n";

    return 0;
}