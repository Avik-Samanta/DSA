#include<iostream>
#include<vector>
using namespace std;

void merge(vector<int>& nums1, int m, vector<int>& nums2, int n){
  int l = m - 1, r = n - 1, k = m + n - 1;
  while( l >= 0 && r >= 0){
    if(nums1[l] >= nums2[r]){
      nums1[k] = nums1[l];
      l--;
    } else {
      nums1[k] = nums2[r];
      r--;
    }
    k--;
  }
  while(r >= 0){
    nums1[k] = nums2[r];
    r--;
    k--;
  }
}


int main(){
  vector<int> nums1 = {1, 2, 3, 0, 0, 0};
  int m = 3;
  int n = 3;
  vector<int> nums2 = {2, 3, 5};
  vector<int> expectedNums = { 1, 2, 2, 3, 3, 5};
  merge(nums1, m, nums2, n);
  try{
    for(int i = 0; i < (n + m); i++){
      if(nums1[i] != expectedNums[i]){
        throw("output does not match check the code logic");
      }
    }
    cout << "output matched !!!" << endl;
  } catch(const char* e){
    cout << e << endl;
  }
  return 0;
};
