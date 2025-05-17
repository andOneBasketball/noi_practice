/*
多重背包问题 II
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
0<N≤1000

0<V≤2000

0<vi,wi,si≤2000

提示：
本题考查多重背包的二进制优化方法。

输入样例
4 5
1 2 3
2 4 1
3 4 3
4 5 2
输出样例：
10

将多重背包问题转化为更少数量的多重背包问题
例如 8 拆分为 1 2 4 1

si 拆分为 1 2 4 si-(2^(n-1)-1)

dp[j] = max(dp[j], dp[j-ss[k]*vi] + ss[k]*wi)
*/

#include <algorithm>
#include <iostream>

using namespace std;

int dp[2005] = {0};

int clac_n_x(int n, int x)
{
    int sum = 1;
    while (x > 0)
    {
        sum *= n;
        x--;
    }
    return sum;
}

int main(int argc, const char *argv[])
{
    freopen("multiple_knapsack_2.in", "r", stdin);
    freopen("multiple_knapsack.out", "w", stdout);

    int n, v, vi, wi, si, i, j, k, count, sum;
    cin >> n >> v;
    for (i = 1; i <= n; i++)
    {
        int ss[15] = {0};
        count = 0;
        cin >> vi >> wi >> si;
        j = 0;
        sum = 0;
        // 将 si 按 2^n 分解到数组内，时间复杂度化简为 n*v*log(si)，约 2*10^7 数量级
        while (1)
        {
            int data = clac_n_x(2, j);
            count++;
            if (sum + data >= si)
            {
                ss[j] = si - sum;
                // cout << j << ":" << ss[j] << endl;
                break;
            }
            else
            {
                ss[j] = data;
                // cout << j << ":" << ss[j] << endl;
            }
            sum += ss[j];
            j++;
        }
        // 优化成0-1背包问题，背包中的物品循环在体积循环的外层
        for (k = 0; k < count; k++)
        {
            for (j = v; j >= vi; j--)
            {
                if (j < ss[k] * vi)
                {
                    // 需要根据前一个体积的价值给 dp[j] 赋值
                    dp[j] = max(dp[j], dp[j - 1]);
                    continue;
                }
                dp[j] = max(dp[j], dp[j - ss[k] * vi] + ss[k] * wi);
                printf("vi:%d; wi:%d; si:%d; j:%d; dp[j]:%d; ss[%d]:%d\n", vi, wi, si, j, dp[j], k, ss[k]);
            }
        }
    }
    cout << dp[v] << endl;

    return 0;
}
