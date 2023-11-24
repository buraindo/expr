#include "klee/klee.h"

int pw(int a, int p) {
	int r = a;
	for (int i = 0; i < p-1; i++) {
		r *= a;
	}
	return r;
}

int is_armstrong(int n) {
	int rightMost;
	int sum = 0;
	int tempNum = n;

	// to get the number of digits in the number
	int length = 1;
	int tmp = n;
	while (tmp >= 10) {
		tmp /= 10;
		length++;
	}

	// get the right most digit and break the loop once all digits are iterated
	while (tempNum > 0) {
		rightMost = tempNum % 10;
		sum += pw(rightMost, length);

		// update the input digit minus the processed rightMost
		tempNum /= 10;
	}

    if (n == sum) {
        return 1;
    }
	return -1;
}

int main() {
  int n;
  klee_make_symbolic(&n, sizeof(n), "n");
  return is_armstrong(n);
}