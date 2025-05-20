令 f(i,j) 表示将区间 [i,j] 内的所有石子合并到一起的最大得分。
写出 状态转移方程：
f(i,j)=max{f(i,k)+f(k+1,j)+ai+...+aj} (i<= k < j)

令 sum_i 表示 a 数组的前缀和，状态转移方程变形为 f(i,j)=max{f(i,k)+f(k+1,j)+sum_j-sum_{i-1} }