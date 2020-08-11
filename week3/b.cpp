#include <bits/stdc++.h> 
using namespace std; 
  
// A function to print all prime  
// factors of a given number n  
int isPrime[1000001];
void sieve(){
    int maxN = 1000000;
    for (int i=1;i<=maxN;i++)isPrime[i]=1;
    isPrime[0]=0;
    isPrime[1]=0;
    for (int i=2; i*i<maxN;i++) {
        if(isPrime[i]){
            for (int j=i*i;j<=maxN;j+=i){
                isPrime[j]=0;
            }
        }
    }
}
/* Driver code */
int main()  
{  
    sieve();
    while(true) {
        int n;
        int count = 0;
        cin>>n;
        if(n==0){
            break;
        }
        if (isPrime[n]==1){
            cout<<"no\n";
        }  else{
            cout<<"yes\n";
        }
    }
    return 0;  
}  
