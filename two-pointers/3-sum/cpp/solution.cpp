#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

vector<vector<int>> threeSum(vector<int>& nums) {
  vector<vector<int>> res;
  sort(nums.begin(), nums.end());
  for(int i = 0; i < nums.size() - 2; i++) {
    if(nums[i] > 0) break;
    if(i > 0 && nums[i] == nums[i - 1]) continue;

    int l = i + 1, r = nums.size() - 1;
    while(l < r) {
      int target = -1 * nums[i];
      int total = nums[l] + nums[r];
      if(total == target){
        res.push_back({nums[i], nums[l], nums[r]});
        l++;
        r--;
        while(l < nums.size() && nums[l] == nums[l - 1]){
          l++;
        }
        while(r >= 0 && nums[r] == nums[r + 1]){
          r--;
        }
      } else if(total < target){
        l++;
      } else {
        r--;
      }
    }
  }
  return res;
}

int main(){
  vector<int> nums = {-1, 0, 1, 2, -1, -4};
  vector<vector<int>> expectedNums = {{-1, -1, 2}, {-1, 0, 1}};
  vector<vector<int>> result = threeSum(nums);
  try{
    if(result == expectedNums){
      cout << "output matched" << endl;
    } else {
      throw("output does not match with expected result");
    }
  } catch(const char* e){
    cout << e << endl;
  }
  return 0;
};
