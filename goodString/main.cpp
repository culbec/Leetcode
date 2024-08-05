#include <iostream>
#include <string>
#include <stack>

class Solution {
public:
    std::string makeGood(std::string s) {
        if (s == "") {
            return s;
        }

        auto n = s.size();
        int dist = 'a' - 'A';

        std::stack<char> st;
        st.push(s[0]);

        for (int i = 1; i < n; i++) {
            if (st.size() == 0) {
                st.push(s[i]);
                continue;
            }

            char top = st.top();

            if (top == s[i] + dist || top == s[i] - dist) {
                st.pop();
            } else {
                st.push(s[i]);
            }
        }

        std::string s_new = std::string(st.size(), '0');
        for (int i = st.size() - 1; i >= 0; i--) {
            s_new[i] = st.top();
            st.pop();
        }

        return s_new;
    }        
};

int main() {
    std::string s1 = "leEeetcode";
    std::string s2 = "abBAcC";
    std::string s3 = "s";
    std::string s4 = "NAanorRoOrROwnTNW";
    std::string s5 = "MqWWvyRtzZTrYVwwQmUjQOoOoqJu";
    std::string s6 = "ogHhGOoEOoeOfaNnAF";

    Solution s = Solution();

    std::cout << s.makeGood(s1) << "\n";
    std::cout << s.makeGood(s2) << "\n";
    std::cout << s.makeGood(s3) << "\n";
    std::cout << s.makeGood(s4) << "\n";
    std::cout << s.makeGood(s5) << "\n";
    std::cout << s.makeGood(s6) << "\n";

    return 0;
}