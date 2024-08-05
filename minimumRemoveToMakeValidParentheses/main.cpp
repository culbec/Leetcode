#include <string>
#include <stack>
#include <deque>
#include <iostream>

using namespace std;

struct ParPair {
    char ch;
    size_t index;
};

string minRemoveToMakeValid(string s) {
    deque<ParPair> par;
    string s_new = "";

    // Using a deque to keep track of the correct opened and closed parenthesees.
    // As using this method, the parenthesees will be kept in the order of appearance,
    // and we can parse the deque to find the incorrect pairs after.
    for(size_t i = 0; i < s.size(); i++) {
        if (s[i] == '(') {
            auto pp = ParPair{s[i], i};
            par.push_back(pp);
        } else if (s[i] == ')') {
            if (par.size() && par.back().ch == '(') {
                par.pop_back();
            } else {
                auto pp = ParPair{s[i], i};
                par.push_back(pp);
            }
        }
    }

    // Constructing the new string.
    for (size_t i = 0; i < s.size(); i++) {
        if (par.size() && i == par.front().index) {
            par.pop_front();
            continue;
        }

        s_new += s[i];
    }

    return s_new;
}

string minRemoveToMakeValid2(string s) {
    stack<char> par;
    // Deleting the ')' that do not have a pair.
    for(size_t i = 0; i < s.size(); i++) {
        if (s[i] == '(') {
            par.push(s[i]);
        } else if (s[i] == ')') {
            if (par.size()) {
                par.pop();
            } else {
                s.erase(i, 1);
                i--;
            }
        }
    }
    // Deleting the '(' that do not have a pair.
    for(int i = s.size() - 1; i >= 0; i--) {
        if (s[i] == ')') {
            par.push(s[i]);
        } else if (s[i] == '(') {
            if (par.size() && par.top() == ')') {
                par.pop();
            } else {
                s.erase(i, 1);
            }
        }
    }
    return s;
}

int main() {
    string s1 = "lee(t(c)o)de)";
    string s2 = "a)b(c)d";
    string s3 = "))((";

    //cout << s1 << ": " << minRemoveToMakeValid(s1) << "\n";
    cout << s1 << ": " << minRemoveToMakeValid2(s1) << "\n\n";

    //cout << s2 << ": " << minRemoveToMakeValid(s2) << "\n";
    cout << s2 << ": " << minRemoveToMakeValid2(s2) << "\n\n";
    
    //cout << s3 << ": " << minRemoveToMakeValid(s3) << "\n";
    cout << s3 << ": " << minRemoveToMakeValid2(s3) << "\n\n";

    cout << ")))t((u): " << minRemoveToMakeValid2(")))t((u)") << "\n\n";

    return  0;
}