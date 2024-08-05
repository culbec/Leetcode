#include "class.h"

template class Calculator<int>;

template<typename T>
Calculator<T>::Calculator(T a, T b) {
    this->a = a;
    this->b = b;
}

template<typename T>
T Calculator<T>::add() {
    return this->a + this->b;
}

template<typename T>
T Calculator<T>::sub() {
    return this->a - this->b;
}

template<typename T>
T Calculator<T>::mul() {
    return this->a * this->b;
}

template<typename T>
T Calculator<T>::div() {
    return this->a / this->b;
}