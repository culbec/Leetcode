#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int maxSubarrayLength(vector<int>& nums, int k) {
        unordered_map<int, int> freqs;

        int left = 0, right = 0, result = 0;
        int n = nums.size();

        while (left < n) {
            while (right < n) {
                if (freqs[nums[right]] == 0) {
                    freqs[nums[right]] = 1;
                } else {
                    freqs[nums[right]]++;
                }

                if (freqs[nums[right]] > k) {
                    break;
                }

                right++;
            }

            if (result < right - left) {
                result = right - left;
            }

            if (right >= n) {
                break;
            }

            while (left < n && freqs[nums[right]] > k) {
                freqs[nums[left]]--;
                left++;
            }

            right++;
        }

        return result;
    }
};