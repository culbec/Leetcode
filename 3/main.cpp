#include <iostream>
#include <unordered_map>
#include <string>

using namespace std;

class Solution {
    public:
        static unsigned int maxLengthSubstringUniqueChars(const string& s) {
            int maximumLength = 0;
            unordered_map<char, int> charPos;

            // boundaries of the current parsed string
            int left = 0, right = 0;

            // the string will be parsed using the right as an index
            for(right; right < s.length(); right++) {
                // verifying if the character was already found
                if(charPos.find(s.at(right)) == charPos.end() || left > charPos.at(s.at(right))) {
                    // this means that the character is unique
                    charPos[s[right]] = right;

                    // updating the current max length of the current bounded substring
                    maximumLength = max(maximumLength, right - left + 1);
                } else {
                    // this means that we need to check the substring
                    // in which we exclude the first appearance of the repeated character
                    left = charPos.at(s.at(right)) + 1;
                    charPos[s[right]] = right;
                }
            }
            return maximumLength;
        }
};

int main() {
    cout << Solution::maxLengthSubstringUniqueChars("abcabcbbb") << endl;
    cout << Solution::maxLengthSubstringUniqueChars("cbdbeff") << endl;
    cout << Solution::maxLengthSubstringUniqueChars("bbbbb") << endl;
    cout << Solution::maxLengthSubstringUniqueChars("abcabcdefacbd") << endl;
    cout << Solution::maxLengthSubstringUniqueChars("tmmzuxt") << endl;
    return 0;
}