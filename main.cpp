#include <vector>
#include <limits.h>

using namespace std;

vector<vector<int>> largestLocal(vector<vector<int>>& grid) {
        int n = grid.size();

        if (n - 2 < 0) {
            return {};
        }

        // Initialize the new matrix with the correct size
        vector<vector<int>> locals(n - 2, vector<int>(n - 2));

        // Use a more efficient algorithm to find the maximum local value
        for (int i = 0; i < n - 2; i++) {
            for (int j = 0; j < n - 2; j++) {
                int maxi = INT_MIN;
                for (int k = 0; k < 3; k++) {
                    for (int l = 0; l < 3; l++) {
                        maxi = max(maxi, grid[i + k][j + l]);
                    }
                }
                locals[i][j] = maxi;
            }
        }

        return locals;
    }
