
#include<iostream>
#include<cstring>
using namespace std;
int checkvalue(char c){
	int ans;
	if(isupper(c)==0){
		ans=int(c)-int('a');
	}
	else
		ans=int(c)-int('A')+26;
	return ans;
}
char checkchar(int i){
	char ans;
	if(i<26){
		ans=char(i+96);
	}
	else ans=char(i+65-26);
}

char nonRepeatingCharacter(string str){
  
int arr[52];
for(int i=0;i<52;i++)
arr[i]=0;
int k;
int len=str.length();
for(int i=0;i<len;i++){
	k=checkvalue(str[i]);
	arr[k]++;
}
for(int i=0;i<len;i++){
	k=checkvalue(str[i]);
	if(arr[k]==1){
		return checkchar(k);
	}
}
return str[0];
}
int main(){
	string s="asddfasf";
	string k;
	k=nonRepeatingCharacter(s);
	cout<<k;
	return 0;
}
