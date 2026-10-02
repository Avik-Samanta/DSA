#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

int countTriplets(vector<int>& nums, int target) {
  int res = 0;
  sort(nums.begin(), nums.end());
  for(int i = 0; i < nums.size() - 2; i++) {
    int l = i + 1, r = nums.size() - 1;
    while(l < r) {
      int total = nums[i] + nums[l] + nums[r];
      if(total < target){
        res = res + (r - l);
        l++;
      } else {
        r--;
      }
    }
  }
  return res;
}

int main() {
  vector<int> nums = {3, 4, 1, 7, 5};
  int target = 12;
  int expectedResult = 4;
  int result = countTriplets(nums, target);
  try{
    if(result == expectedResult) {
      cout << "output matched with the expected result !!!" << endl;
    } else {
      throw("out put does not match with expected result");
    }
  } catch(const char* e){
    cout << e << endl;
  }
  cout << result << endl;
  return 0;
}
