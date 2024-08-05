class Solution {
private:
    vector<pair<int, int>> dirs = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    bool isWithinBounds(const vector<vector<int>>& grid, int i, int j) {
        int n = grid.size();
        return (i >= 0 && i < n) && (j >= 0 && j < n);
    }

    bool existsValidPath(const vector<vector<int>>& grid, int safeness) {
        int n = grid.size();

        // Checking if the source and destination comply with the imposed safeness.
        if (grid[0][0] < safeness || grid[n - 1][n - 1] < safeness) {
            return false;
        }

        queue<pair<int, int>> q;
        vector<vector<bool>> visited(n, vector<bool>(n, false));
        
        q.push({0, 0});
        visited[0][0] = true;

        while (!q.empty()) {
            auto curr = q.front();
            q.pop();

            if (curr.first == n - 1 && curr.second == n - 1) {
                return true;
            }

            for (const auto& d: dirs) {
                int di = curr.first + d.first;
                int dj = curr.second + d.second;

                if (isWithinBounds(grid, di, dj)) {
                    if (!visited[di][dj] && grid[di][dj] >= safeness) {
                        visited[di][dj] = true;
                        q.push({di, dj});
                    }
                }
            }
        }

        return false; // No valid path from (0, 0) to (n-1, n-1) was found.
    }

    int binarySearch(const vector<vector<int>>& grid) {
        int n = grid.size();

        int left = 0, right = 0, result = -1;

        // Determining the maximum safeness factor.
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                right = max(right, grid[i][j]);
            }
        }

        // Searching for the maximum safeness that ensures a valid path.
        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (existsValidPath(grid, mid)) {
                result = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return result;
    }



public:
    int maximumSafenessFactor(vector<vector<int>>& grid) {
        int n = grid.size();

        // Initializing a queue to perform BFS from the thief cells
        // for determining the Manhattan distance between the free cells and the thief celss.
        queue<pair<int, int>> q;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    q.push({i, j});
                    grid[i][j] = 0;
                } else {
                    grid[i][j] = -1; // no initial mark on the free cells.
                }
            }
        }

        // Determining the safeness factor of each cell.
        while(!q.empty()) {
            int size = q.size();

            while (size-- > 0) {
                auto curr = q.front(); // extracting the coordinates of a thief cell.
                q.pop();

                // Checking the neighboring cell and updating its safeness if possible.
                for(const auto& d: dirs) {
                    int di = curr.first + d.first;
                    int dj = curr.second + d.second;
                    int thief = grid[curr.first][curr.second];

                    if (isWithinBounds(grid, di, dj)) {
                        if (grid[di][dj] == -1) {
                            grid[di][dj] = thief + 1;
                            q.push({di, dj});
                        }
                    }
                }
            }
        }

        return binarySearch(grid);
    }
};
