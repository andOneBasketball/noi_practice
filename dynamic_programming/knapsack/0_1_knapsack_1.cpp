/*
有 N
 件物品和一个容量是 V
 的背包。每件物品只能使用一次。

第 i
 件物品的体积是 vi
，价值是 wi
。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V
，用空格隔开，分别表示物品数量和背包容积。

接下来有 N
 行，每行两个整数 vi,wi
，用空格隔开，分别表示第 i
 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤1000

0<vi,wi≤1000

输入样例
4 5
1 2
2 4
3 4
4 5
输出样例：
8

dp[n] = max(dp[n], dp[n-vi] + wi)
n 表示体积, dp[n] 表示最大价值
 */

#include <iostream>
#include <algorithm>

using namespace std;

int dp[1005] = {0};

int main(int argc, char const *argv[])
{
    //freopen("0_1_knapsack_1.in", "r", stdin);
    //freopen("0_1_knapsack.out", "w", stdout);

    int n, v, vi, wi, i, j;
    cin >> n >> v;

    for (i = 1; i <= n; i++)
    {
        cin >> vi >> wi;
        for (j = v; j >= vi; j--)
        {
            dp[j] = max(dp[j], dp[j-vi] + wi);
            //cout << j << "; dp[j]:" << dp[j] << endl;
        }
    }

    cout << dp[v] << endl;

    return 0;
}