#include <iostream>
#include<cmath>


int main()
{
    float a= sqrt(12) * (1 - (1 / (float)(3 * 3)) + (1 / (5 * pow(3, 2))) - (1 / (7 * pow(3, 3))) + (1 / (9 * pow(3, 4))) - (1 / (11 * pow(3, 5))));

    std::cout << a;
    
        
return 0;

}