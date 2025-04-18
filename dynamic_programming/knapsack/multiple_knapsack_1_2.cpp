/*
有 N
 种物品和一个容量是 V
 的背包。

第 i
 种物品最多有 si
 件，每件体积是 vi
，价值是 wi
。

求解将哪些物品装入背包，可使物品体积总和不超过背包容量，且价值总和最大。
输出最大价值。

输入格式
第一行两个整数，N，V
，用空格隔开，分别表示物品种数和背包容积。

接下来有 N
 行，每行三个整数 vi,wi,si
，用空格隔开，分别表示第 i
 种物品的体积、价值和数量。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤100

0<vi,wi,si≤100

输入样例
4 5
1 2 3
2 4 1
3 4 3
4 5 2
输出样例：
10

dp[n][v] = max(dp[n-1][v], dp[n][v-k*vi] + k*wi)
时间复杂度：O(n*v*si)
dp[n][v] 表示前 n 件物品，放入容积为 v 的背包可以获得的最大价值
 */

#include <iostream>
#include <algorithm>

using namespace std;

int dp[105][105] ={0};

int main(int argc, char const *argv[])
{
    //freopen("multiple_knapsack_2.in", "r", stdin);
    //freopen("multiple_knapsack.out", "w", stdout);
    int n, v, vi, wi, si, i, j, k;
    cin >> n >> v;
    for (i = 1; i <= n; i++)
    {
        cin >> vi >> wi >> si;

        for (j = 1; j <= v; j++)
        {
            // 比较 n 件物品和 n-1 件物品的最大价值
            dp[i][j] = max(dp[i][j], dp[i-1][j]);
        }

        for (j = 1; j <= v; j++)
        {
            for (k = 1; k <= si; k++)
            {
                if (vi*k > j)
                {
                    dp[i][j] = max(dp[i][j], dp[i-1][j]);
                    continue;
                }
                // 需要重复比较选择该物品和不选该物品的最大价值
                dp[i][j] = max(dp[i][j], dp[i-1][j-vi*k] + k*wi);
                //错误：违反物品有限数量的限制
                //f[i][j] = max(f[i-1][j], f[i][j-k*vi] + k*wi);
            }
        }
    }
    cout << dp[n][v] << endl;
    return 0;
}