bool compare(std::string a, std::string b) {
    int c = 0;
    int d = 0;
    for (int i = 0; i < a.length(); i++) {
        if (a[i] == '1') c += 1;
    }
    for (int i = 0; i < b.length(); i++) {
        if (b[i] == '1') d += 1;
    }

    if (c != d) {
        return c > d;
    } else {
        return std::stoi(a) < std::stoi(b);
    }
}