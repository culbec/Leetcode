class Solution {
public:
    int islandPerimeter(vector<vector<int>>& grid) {
        int perimeter = 0;
        int rows = grid.size();
        int cols = grid[0].size();

        for (int i = 0; i < rows; i++) {
            for(int j = 0; j < cols; j++) {
                if (grid[i][j] == 1) {
                    perimeter += 4;
                    // Checking the neighbours.
                    for (int k = 0; k < 4; k++) {
                        if ((i + dx[k] > -1 && i + dx[k] < rows) && 
                            (j + dy[k] > -1 && j + dy[k] < cols) &&
                            (grid[i + dx[k]][j + dy[k]] == 1)) {
                                perimeter -= 1;
                            }
                    }
                }
            }
        }

        return perimeter;
    }
private:
    int dx[4] = {0, 1, 0, -1};
    int dy[4] = {-1, 0, 1, 0};
};
