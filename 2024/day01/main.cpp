#include <algorithm>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>

int main() {
  std::fstream file("./input.txt");
  if (!file) {
    std::cerr << "File could not be opened" << std::endl;
    return 1;
  }

  std::vector<int> v1;
  std::vector<int> v2;

  std::string line;

  while (std::getline(file, line)) {
    std::stringstream ss(line);
    int num1, num2;
    ss >> num1 >> num2;
    v1.push_back(num1);
    v2.push_back(num2);
  }
  std::sort(v1.begin(), v1.end());
  std::sort(v2.begin(), v2.end());

  int ans = 0;
  for (int i = 0; i < v1.size(); i++) {
    ans += abs(v1[i] - v2[i]);
  }

  std::cout << "Answer: " << ans << std::endl;
  file.close();
  return 0;
}
