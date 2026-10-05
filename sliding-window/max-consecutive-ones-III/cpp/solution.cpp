#include<iostream>
#include<vector>
using namespace std;

int longestOne(vector<int>nums, int k){
  int low = 0, res = INT_MIN;
  vector<int> numArr(2);

  for(int high = 0; high < nums.size(); high++){
    numArr[nums[high]]++;

    while(numArr[0] > k){
      numArr[nums[low]]--;
      low++;
    }

    int len = high - low + 1;
    res = max(res, len);
  }
  return res;
}

int main(){
  vector<int> nums = {0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1};
  int k = 3;
  int expectedRes = 10;
  
  int result = longestOne(nums, k);
  
  try {
    if(expectedRes == result) {
      cout << result << endl << "right answer !!!" << endl;
    } else {
      throw("wrong ans, plz check the code");
    }
  }
  catch (const char* e) {
    cout << e << endl;
  }
}
