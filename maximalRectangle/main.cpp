#include <vector>
#include <stack>

using namespace std;

class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        int maxArea = 0;
        vector<int> heights(matrix[0].size(), 0);        

        for (const auto& row: matrix) {
            for (size_t i = 0; i < row.size(); i++) {
                heights[i] = (row[i] == '0') ? 0 : heights[i] + 1;
            }

            maxArea = max(maxArea, this->largestRectangleSoFar(heights));
        }

        return maxArea;
    }

private:
    int largestRectangleSoFar(vector<int>& heights) {
        int maxArea = 0;
        stack<int> s;

        for (size_t i = 0; i <= heights.size(); i++) {
            while (!s.empty() && (i == heights.size() || heights[i] < heights[s.top()])) {
                int h = heights[s.top()]; s.pop();
                int w = s.empty() ? i : i - s.top() - 1;
                maxArea = max(maxArea, h * w);
            }
            s.push(i);
        }

        return maxArea;
    }
};
