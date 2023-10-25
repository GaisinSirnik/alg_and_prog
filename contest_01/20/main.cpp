#include <iostream>
#include <map>

int main() {
    std::string a, b;
    std::cin >> a;
    std::cin >> b;

    std::map<char, int> c, d;
    for (int i = 0; i < a.length(); i++) {
        c[a[i]]++;
    }
    for (int i = 0; i < b.length(); i++) {
        d[b[i]]++;
    }
    bool e = true;
    for (auto it = c.begin(); it != c.end(); ++it) {
        if (d[it->first] < it->second) {
            e = false;
            break;
        }
    }
    std::cout << (e ? "YES" : "NO");
    return 0;
}