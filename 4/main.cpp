#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        /*
            Base cases:
            1. both arrays are empty;
            2. one of the arrays is empty;
            3. none of the arrays is empty;
        */

       if(nums1.size() == 0 && nums2.size() == 0) {
            return (double) 0;
       } else if(nums1.size() != 0 && nums2.size() == 0) {
            if(nums1.size() % 2) {
                return (double) nums1.at(nums1.size() / 2);
            } else {
                return (double) (nums1.at(nums1.size() / 2) + nums1.at((nums1.size() - 1) / 2)) / 2;
            }
       } else if(nums1.size() == 0 && nums2.size() != 0){
            if(nums2.size() % 2) {
                return (double) nums2.at(nums2.size() / 2);
            } else {
                return (double) (nums2.at(nums2.size() / 2) + nums2.at((nums2.size() - 1) / 2)) / 2;
            }
       }

        int i = 0, j = 0;
        vector<int> result;

        while(i < nums1.size() && j < nums2.size()) {
            if(nums1[i] < nums2[j]) {
                result.push_back(nums1[i++]);
            } else {
                result.push_back(nums2[j++]);
            }
        }

        while(i < nums1.size()) {
            result.push_back(nums1[i++]);
        }

        while(j < nums2.size()) {
            result.push_back(nums2[j++]);
        }

        if(result.size() % 2) {
            return (double) (result[result.size() / 2]);
        }
        return (double) (result[result.size() / 2] + result[result.size() / 2 - 1]) / 2;
    }
};

int main() {
    Solution solution;

    vector<int> nums1 = {1, 2};
    vector<int> nums2 = {3};
    cout << solution.findMedianSortedArrays(nums1, nums2) << endl;

    return 0;
}