#include<iostream>
#include<vector>
using namespace std;

vector<int> twoSum(vector<int>& numbers, int target) {
  int l = 0, r = numbers.size() - 1;
  while( l < r){
    int total = numbers[l] + numbers[r];
    if(total == target){
      return { l + 1, r + 1 };
    } else if(total < target){
      l++;
    } else {
      r--;
    }
  }
  return {};
};

int main(){
  vector<int> numbers = {2, 7, 11, 15};
  int target = 9;
  vector<int> result = twoSum(numbers, target);
  cout << "[" << result[0] << "," << result[1] << "]" << endl;
  return 0;
};
