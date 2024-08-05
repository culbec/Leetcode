template<typename T>
class Calculator {
    private:
        T a;
        T b;
    public:
        Calculator(T a, T b);

        T add();

        T sub();

        T mul();

        T div();

        ~Calculator() = default;
};