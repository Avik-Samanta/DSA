#include<iostream>
#include<vector>
using namespace std;

vector<int> sortedSquare(vector<int>& nums) {
  vector<int> res;
  int l = -1, r = 0;
  for(int i = 0; i < nums.size(); i++){
    if(nums[i] < 0){
      l++;
    } 
  }
  r = l + 1;
  while(l >= 0 && r < nums.size()){
    if(nums[l] * nums[l] >= nums[r] * nums[r]){
      res.push_back(nums[r] * nums[r]);
      r++;
    } else {
      res.push_back(nums[l] * nums[l]);
      l--;
    }
  }
  while(l >= 0){ 
    res.push_back(nums[l] * nums[l]);
    l--;
  }
  while(r < nums.size()){
    res.push_back(nums[r] * nums[r]);
    r++;
  }
  return res;
} 


int main(){
  vector<int> nums = {-4, -2, 0, 3, 10};
  vector<int> expectedNums = {0, 4, 9, 16, 100};
  vector<int> res = sortedSquare(nums);
  try{
    if(nums.size() != expectedNums.size()){
      throw("output size does not mtach with the expected size");
    }
    for(int i = 0; i < res.size(); i++){
      if(res[i] != expectedNums[i]){
        throw("output does not match check the code logic");
      }
    }
    cout << "output matched !!!" << endl;
    for(int x: res){
      cout << x << endl;
    }
  } catch(const char* e){
    cout << e << endl;
  }
  return 0;
};
