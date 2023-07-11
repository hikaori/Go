# Basic test of GO

## How to test go project

1. create base testing file.
   Do the steps of [How to create go project](../Basic/README.md) with project name call "MyTest".

   **need to modify "package main" => "package MyTest" inside main.go.**

2. create test function.
   change inside main.go as below.

   ```
   package MyTest


   func Sum(x int, y int) int {
      return x + y
   }

   func main() {
      Sum(5, 5)
   }
   ```

3. create test file.
   create new file called main_test.go

   **test file prefix is < filename >\_test.go**

4. write test.
   inside main_test.go write as below

   ```
   package MyTest

   import (
      "testing"
   )

   func TestSum(t *testing.T) {
      total := Sum(5, 5)
      if total != 10 {
         t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
      }
   }
   ```

5. run test.

   ```
   $go test
   ```

   in terminal, this should be display

   ![Screenshot 2023-07-11 at 14 57 19](https://github.com/hikaori/Go/assets/23109342/8a71fbd8-97df-45d6-add1-05bcb5399939)
