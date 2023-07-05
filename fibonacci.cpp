#include <iostream>
#include <thread>
#include <vector>


int fibonacci(int n) {
    if (n <= 1)
        return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

void calcFibonacci(int n, std::vector<int>& result) {
    result.resize(n + 1);
    for (int i = 0; i <= n; ++i) {
        result[i] = fibonacci(i);
    }
}

int main() {
    int n = 10;
    std::vector<int> fibonacciSeq; //stores result as vector

    std::thread t(calcFibonacci, n, std::ref(fibonacciSeq));
    t.join();
  //thread passes through calcFibonacci

    std::cout << "Fibonacci sequence up to " << n << ":" << std::endl;
    for (int num : fibonacciSeq) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    return 0;
}
