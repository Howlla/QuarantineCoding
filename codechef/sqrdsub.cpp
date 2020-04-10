// C++ program to find product 
// of all subarray of an array 
  
#include <bits/stdc++.h> 
using namespace std; 
#define ll long long 
#define MAX 100002 
#define pb push_back 

ll tree[4 * MAX]; 

void build(int node, int start, int end, const int* arr) 
{ 
    if (start == end) { 
        tree[node] = (1LL * arr[start]); 
        return; 
    } 
    int mid = (start + end) >> 1; 
    build(2 * node, start, mid, arr); 
    build(2 * node + 1, mid + 1, end, arr); 
    tree[node] = (tree[2 * node] * tree[2 * node + 1]); 
} 
ll query(int node, int start, int end, int l, int r) 
{ 
    if (start > end || start > r || end < l) { 
        return 1; 
    } 
    if (start >= l && end <= r) { 
        return tree[node]; 
    } 
    int mid = (start + end) >> 1; 
    ll q1 = query(2 * node, start, mid, l, r); 
    ll q2 = query(2 * node + 1, mid + 1, end, l, r); 
    return (q1 * q2); 
} 
ll countSubarrays(const int* arr, int n) 
{ 
    ll count = 0; 
    for (int i = 0; i < n; i++) { 
        for (int j = i; j < n; j++) { 
  
            // Query segment tree to find product % k 
            // of the sub-array[i..j] 
            ll product_mod_k = query(1, 0, n - 1, i, j); 
            if (product_mod_k%4 !=2) { 
                count++; 
            } 
        } 
    } 
    return count; 
} 


  
// Driver code 
int main() 
{ 
    // sieve(); 
    int t;
    cin>>t;
    while(t--){
        int n; 
        cin>>n;
        int arr[n];
        for(int i=0;i<n;i++){
            cin>>arr[i];
        }
        build(1, 0, n - 1, arr); 

        cout<<countSubarrays(arr, n)<<endl; 
    } 
  return 0;
  
} 