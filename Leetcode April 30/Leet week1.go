//Day1
func singleNumber(nums []int) int {
    ans:=0;
    for i:=0;i<len(nums);i++{
        ans = ans^nums[i];
    }
    return ans;
}

//Day 2
func isHappy(n int) bool {
    var mp map[int]bool;
    mp = make(map[int]bool);
    for{
        if(mp[n]==false){
            mp[n]=true
        }else{
            if(n==1){
                return true
            }else{
                return false
            }
        }
        n=findNext(n);
    }

}
func findNext(n int) int{
    var res int;
    for;n>0;n/=10{
        res+=(n%10)*(n%10)
    }
    return res;
}


//Day 3
func maxSubArray(nums []int) int {
    n:=len(nums);
    var maxSoFar, maxEndingHere int;
    maxSoFar = nums[0]
    maxEndingHere = nums[0]
    for i:=1;i<n;i++{
        maxEndingHere = max(maxEndingHere+nums[i],nums[i])
        maxSoFar = max(maxSoFar,maxEndingHere)
    }
   
    return maxSoFar;
}

func max(a,b int) int{
    if(a>b){
        return a;
    }
    return b;
}


//Day 4

func moveZeroes(nums []int)  {
    var index int;
    for i:= range nums {
        if(nums[i]!=0){
            nums[index]=nums[i];
            index++
        }
    }
    for i:=index;i<len(nums);i++{
        nums[i]=0
    }
}

//Day 5

func maxProfit(prices []int) int {
    var n int = len(prices)
    var sum,i int;
    if(len(prices)==0){
        return 0
    }
    valley,peak:= prices[0],prices[0]
    for i<n-1{
        for i<n-1 && prices[i]>=prices[i+1] {
            i++
        }
        valley = prices[i]
        for i<n-1 && prices[i]<= prices[i+1] {
            i++
        }
        peak = prices[i]
        sum+=peak-valley
    }
    return sum
    
}


//Day 6
func groupAnagrams(strs []string) [][]string {
    // if(len(strs)==0) return new([][]string)
    hm:= make(map[string][]string)
    for _,w:= range strs {
        temp := SortString(w) 
        hm[temp] = append(hm[temp],w)
    }
    // return hm
    var ans [][]string;
    for _,val := range hm {
        ans= append(ans,val)
    }
    return ans;
}


func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}


//Day 7
func countElements(arr []int) int {
    hm:=make(map[int]bool)
    for i:= range arr {
        hm[arr[i]] = true;
    }
    var count int;
    for i:=0;i<len(arr);i++{
        if(hm[arr[i]+1]==true){
            count++
        }
    }
    return count
}