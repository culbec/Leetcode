#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
  int timeRequiredToBuy(vector<int> &tickets, int k)
  {
    int count = 0;

    if (k < 0 || (size_t)k > tickets.size())
    {
      return count;
    }

    int ticketsSize = tickets[k];

    while (ticketsSize > 0)
    {
      for (int i = 0; i < k; i++)
      {
        if (tickets[i] > 0)
        {
          count++;
          tickets[i]--;
        }
      }

      for (size_t i = k + 1; i < tickets.size(); i++)
      {
        if (tickets[i] > 0)
        {
          count++;
          tickets[i]--;
        }
      }

      count++;
      ticketsSize--;
    }

    return count;
  }

  int timeRequiredToBuy2(vector<int> &tickets, int k)
  {
    int count = tickets[k];

    for (int i = 0; i < k; i++)
    {
      count += min(tickets[i], tickets[k]);
    }

    for (size_t i = k + 1; i < tickets.size(); i++)
    {
      // After the kth pass of the queue, the person behind the kth pass can only buy k - 1 tickets.
      count += min(tickets[i], tickets[k] - 1);
    }

    return count;
  }
};

int main()
{
  Solution s = Solution();

  vector<int> t1 = {2, 3, 2};
  vector<int> t2 = {5, 1, 1, 1};
  vector<int> t3 = {84, 49, 5, 24, 70, 77, 87, 8};

  cout << s.timeRequiredToBuy2(t1, 2) << "\n";
  cout << s.timeRequiredToBuy2(t2, 0) << "\n";
  cout << s.timeRequiredToBuy2(t3, 3) << "\n";

  return 0;
}
