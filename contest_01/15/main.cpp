#include <iostream>

using namespace std;

int main() {
    string num;
    cin >> num;
    int stroka = 1;
    char a = num[0];
    for (int i = 1; i < size(num); i++) {
        char n = num[i];
        if (a == n) {
            stroka++;
        }
        else {
            if (stroka > 1) {
                cout << a << stroka;
            }
            else {
                cout << a;
            }
            stroka = 1;
            a = n;
        }
    }
    if (stroka > 1) {
        cout << a << stroka;
    }
    else {
        cout << a;
    }
    return 0;
}