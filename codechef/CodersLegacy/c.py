def isBalanced(s,i,mid):
    counter=1
    i+=1
    while i<=mid:
        if s[i]=='(':
            counter+=1
        else:
            if counter>0:
                counter-=1
                if counter==0:
                    return True
        i+=1
    return False
t=int(input())
while t>0:
    t-=1
    s=input()
    q=int(input())
    nums=list(map(int,input().split()))
    x=0
    while q>0:
        q-=1
        index=int(nums[x])
        x+=1
        i=index-1
        while i<len(s):
            if s[i]==')':
                i+=1
            else:
                break
        low=i
        high=len(s)-1
        while low<=high:
            mid=low+(high-low)//2
            if isBalanced(s,i,mid):
                high=mid-1
            else:
                low=mid+1
        if low>=len(s):
            print(-1)
        else:
            print(low+1)

