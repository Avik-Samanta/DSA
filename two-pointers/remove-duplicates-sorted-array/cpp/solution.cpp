#include<iostream>
#include<vector>
using namespace std;

int removeDuplicates(vector<int> &nums){
  int l = 0, r = 1, k = 1;
  while( r < nums.size()){
    if(nums[r] == nums[r - 1]){
      r++;
      continue;
    } else {
      nums[l + 1] = nums[r];
      l++;
      r++;
      k++;
    }
  }
  return k;
}



int main(){
  vector<int> nums = {0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4};
  vector<int> expectedNums = {0, 1, 2, 3, 4};
  int res = removeDuplicates(nums);
  try{
    if(res != expectedNums.size()){
      throw("wrong result of k");
    }
    for(int i = 0; i < res; i++){
      if(nums[i] != expectedNums[i]){
        throw("error in the code");
      }
    }
    cout << "correct" << "\n" << res << endl;
  } catch(const char* e){
    cout << e << endl;
  }
  return 0;
};
