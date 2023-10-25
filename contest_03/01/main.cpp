#include <iostream>
#include <set>

int main() {
    std::multiset<long> slov;
    int f;
    std::cin >> f;
    for (int i = 0; i < f; i++) {
        long a;
        std::cin >> a;
        slov.insert(a);
    }
    std::cin >> f;
    for (int i = 0; i < f; i++) {
        long a;
        std::cin >> a;
        slov.insert(a);
    }
    for (auto b : slov) {
        std::cout << b << " ";
    }
    std::cout << "\n";
    return 0;
}