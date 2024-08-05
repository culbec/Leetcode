#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    int countStudents(vector<int>& students, vector<int>& sandwiches) {
        size_t i = 0;   

        do {
            if (*students.begin() != *sandwiches.begin()) {
                // Placing the student at the end of the queue.
                int student = *students.begin();
                students.erase(students.begin());
                students.push_back(student);
            } else {
                // Erasing the student and the sandwich from the queue & stack.
                students.erase(students.begin());
                sandwiches.erase(sandwiches.begin());
                i = -1;
            }
            i++;
        } while(i < students.size());

        return students.size();
    }
};

int main(int argc, char* argv[]) {
  vector<int> st1 = {1, 1, 0, 0};
  vector<int> sw1 = {0, 0, 1, 1}; 

  Solution s = Solution();
  
  cout << s.countStudents(st1, sw1) << "\n";
  return 0;
}
