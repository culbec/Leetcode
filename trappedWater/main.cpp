#include <vector>
#include <cassert>

using namespace std;

class Solution {
public:
    int trap(vector<int>& height) {
        if (height.size() == 0) {
            return 0;
        }

        int n = height.size();

        vector<int> left = vector<int>(n, 0);
        vector<int> right = vector<int>(n, 0);

        left[0] = height[0];
        // Determining the maximum left height at each index.
        for(int i = 1; i < n; i++) {
            left[i] = max(left[i - 1], height[i]);
        }

        right[n - 1] = height[n - 1];
        // Determining the maximum right height at each index.
        for(int i = n - 2; i >= 0; i--) {
            right[i] = max(right[i + 1], height[i]);
        }

        int water = 0;
        for (int i = 0; i < n; i++) {
            water += min(left[i], right[i]) - height[i];
        }

        return water;
    }
};

int main() {
  vector<int> h1 = {0,1,0,2,1,0,1,3,2,1,2,1};
  vector<int> h2 = {4,2,0,3,2,5};
  
  Solution s = Solution();

  assert(s.trap(h1) == 6);
  assert(s.trap(h2) == 9);

  return 0;
}
