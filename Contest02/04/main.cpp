#include <utility>
#include <numeric>
#include <algorithm>
std::pair<int, int> reduce(int m, int n) {
    int gcd = std::gcd(m, n);
    return std::make_pair(m / gcd, n / gcd);
}
std::tuple<int, int, int> find_lcm(int m, int n) {
    int lcm = std::lcm(m, n);
    return std::make_tuple(lcm, lcm / m, lcm / n);
}