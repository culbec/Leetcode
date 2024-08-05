#include <iostream>
#include <stack>
#include <string>

using namespace std;

class Solution {
  public:
    int maxDepthStack(string s) {
        int ans = 0;
        stack<char> st;

        for(auto& ch: s) {
            if (ch == '(') {
                st.push(ch);
            } else if (ch == ')') {
                st.pop();
            }

            if (ans < st.size()) {
                ans = st.size();
            }
        }

        return ans;
    }

    int maxDepthCounter(string s) {
        int ans = 0, count = 0;

        for(auto& ch: s) {
            if (ch == '(') {
                count++;
            } else if (ch == ')') {
                count--;
            }

            ans = (ans < count) ? count : ans;
        }

        return ans;
    }
};

int main() {
    string s1 = "";
    string s2 = "(1+(2*3)+((8)/4))+1";
    string s3 = "(1)+((2))+(((3)))";

    Solution solution = Solution();

    cout << s1 << " stack: " << solution.maxDepthStack(s1) << "\n";
    cout << s1 << " count: " << solution.maxDepthCounter(s1) << "\n\n";

    cout << s2 << " stack: " << solution.maxDepthStack(s2) << "\n";
    cout << s2 << " count: " << solution.maxDepthCounter(s2) << "\n\n";

    cout << s3 << " stack: " << solution.maxDepthStack(s3) << "\n";
    cout << s3 << " count: " << solution.maxDepthCounter(s3) << "\n\n";  

    return 0;
}