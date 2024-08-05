#include <iostream>
#include "class.cpp"

using namespace std;

int main() {
    Calculator<int> caca(1, 2);

    cout << caca.add() << endl;

    return 0;
}