# executioner

To split your code to separate files use this format:  
`-- {directory_name}/{filename.extension} --`  


Golang example:

```
package main  

import (
    "fmt"
    "os"
    "greenteam/sum"
)

func main() {
    var a, b int
    fmt.Fscan(os.Stdin, &a)
    fmt.Fscan(os.Stdin, &b)
    fmt.Print(sum.Sum(a, b))
}

-- sum/sum.go --
package sum

func Sum(a, b int) int {
    return a + b
}
```  

C++ example:

```
#include <iostream>
#include "sum.cpp"

int main() {
    int a, b;
    std::cin >> a >> b;
    std::cout << summary(a, b) << std::endl;
}

-- sum.cpp --
int summary(int a, int b) {
    return a + b;
}

```