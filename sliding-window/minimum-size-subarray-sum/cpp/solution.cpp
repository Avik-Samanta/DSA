#include<iostream>
#include<vector>
using namespace std;

int minSubArrayLen(vector<int>& nums, int target) {
  int res = INT_MAX, low = 0, high = 0, sum = 0;
  while(high < nums.size()) {
    sum += nums[high];
    while(sum >= target) {
      int len =  high - low + 1;
      res = min(res, len);
      sum -= nums[low];
      low++;
    }
    high++;
  }
  return res == INT_MAX ? 0 : res;
}

int main() {
  vector<int> nums = {2, 3, 1, 2, 4, 3};
  int target = 7;
  int expectedRes = 2;
  int res = minSubArrayLen(nums, target);
  if(res == expectedRes) {
    cout << "right ans" << endl << res << endl;
  } else {
    cout << "wrong output" << endl;
  }
}
