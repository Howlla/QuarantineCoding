#include<iostream>
#include<algorithm>
#include<vector>
#include<set>
#include<math>
vector<int> longestSubsequence(int *arr, int n){
	// Write your code here
  vector<int> v(arr,arr+n);
  int temp=0;
  vector<int> ans;
  vector<int>::iterator it;
  set<int> S;
  for(it=v.begin();it!=v.end();it++){
    S.insert(*it);
    
  }
  
   for(it=v.begin();it!=v.end();it++){
	if(S.find(*it-1)==S.end()){
      int k=*it;
      
      while(S.find(k)!=S.end()){
        
      	ans.push_back(k);
        k++;
      }
      int temp2=temp;
      temp=max(temp,k-*it);
      if(temp2<=temp){
        ans.clear();
      }
    }
      
   }
  return ans;
}
int main()
{
    int vector<int> a;
    int vector<int>::iterator it;

    int arr[] =  {1, 9, 3, 10, 4, 20, 2};
    int n = sizeof arr/ sizeof arr[0];
    a=longestSubsequence(a,n);
    for(it=a.begin();it!=a.end();it++)
      cout<<*it<<" ";
    
    return 0;
}