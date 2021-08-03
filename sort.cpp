std::string myStr = "abc";
std::stable_sort(std::begin(myStr), std::end(myStr));
do {
for(auto&& element : myStr)
std::cout << element << " ";
std::cout << std::endl;
} while (std::next_permutation(std::begin(myStr), std::end(myStr)));

