/*
多重背包问题 III
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
 (0<N≤1000
, 0<V≤20000)
，用空格隔开，分别表示物品种数和背包容积。

接下来有 N
 行，每行三个整数 vi,wi,si
，用空格隔开，分别表示第 i
 种物品的体积、价值和数量。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N≤1000

0<V≤20000

0<vi,wi,si≤20000

提示
本题考查多重背包的单调队列优化方法。

输入样例
4 5
1 2 3
2 4 1
3 4 3
4 5 2
输出样例：
10

dp[v] = max(dp[v], dp[v-vi]+wi, dp[v-2vi]+2wi, ... , dp[v-si*vi]+si*wi)
r = v % vi
dp[v] = max(dp[r]+(v-r)/vi * wi, dp[r+vi]+(v-r-vi)/vi * wi, ... , dp[v])
*/

#include <algorithm>
#include <iostream>
#include <queue>
using namespace std;

int dp[20005] = {0};

int main(int argc, char *argv[])
{
    freopen("multiple_knapsack_3.in", "r", stdin);
    freopen("multiple_knapsack_3.out", "w", stdout);
    int n, v, vi, wi, si;
    int i, j, k;
    cin >> n >> v;
    for (i = 1; i <= n; i++)
    {
        // vi 体积，wi 价值，si 数量
        cin >> vi >> wi >> si;
        if (vi * si >= v)
        {
            for (j = vi; j <= v; j++)
            {
                dp[j] = max(dp[j], dp[j - vi] + wi);
            }
        }
        else
        {
            for (j = 0; j < vi; j++)
            {
                int last_num = v - vi;
                int count = 0;
                for (k = v; k >= j && last_num >= 0; k -= vi)
                {
                    if (count <= si)
                    {
                        dp[k] = max(dp[k], dp[last_num] + (k - last_num) / vi * wi);
                    }
                    else
                    {
                        dp[k] = max(dp[k], dp[last_num]);
                    }
                    count++;
                    last_num = k - vi;
                    printf("dp[%d]:%d; j:%d, i:%d, vi:%d, wi:%d, si:%d\n", k, dp[k], j, i, vi, wi, si);
                }
                /*
                for (k = j; k <= v; k += vi)
                {
                    if (count <= si)
                    {
                        dp[k] = max(dp[k], dp[last_num] + (k - last_num) / vi * wi);
                    }
                    else
                    {
                        dp[k] = max(dp[k], dp[last_num]);
                    }
                    count++;
                    last_num = k;
                    printf("dp[%d]:%d; j:%d, i:%d, vi:%d, wi:%d, si:%d\n", k, dp[k], j, i, vi, wi, si);
                }
                */
            }
        }
    }

    cout << dp[v] << endl;
    return 0;
}