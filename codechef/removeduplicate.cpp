#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;


vector<int> removeDuplicates(vector<int> input){
  vector<int> temp;
  vector<int>::iterator it;
  for(it=input.begin();it!=input.end();it++){
  	if(find(temp.begin(),temp.end(),*it)==temp.end())
  		temp.push_back(*it);
  }
  return temp;

    // Write your code here

}
int main(){

	static const int arr[] = {16,2,2,77,77,29};
vector<int> a (arr, arr + sizeof(arr) / sizeof(arr[0]) );
	vector<int> b;
	b=removeDuplicates(a);
	vector<int>::iterator it;
	for(it=b.begin();it!=b.end();it++){
		cout<<*it<<" ";
	}
return 0;

}