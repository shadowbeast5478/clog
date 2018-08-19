# Clog
- Multi-level coloring logger.
- Print coloring information
    - Prefix word
    - File name
    - Line
    - Function name
## Installation
- Simply run
```shell=
$ go get github.com/shadowbeast5478/clog
```
## Library
## Library
- ```clog.Info()``` ：Print greed prefix word ```[INFO]```
- ```clog.Debug()``` ：Print blue prefix word ```[DEBUG]```
- ```clog.Warning()```：Print yellow prefix word ```[Warning]```
- ```clog.Error()``` ：Print red prefix word ```[ERROR]```
## Examples
- Example 1 : Using clog only
    ```go
    package main

    import (
        "github.com/shadowbeast5478/clog"
    )

    func main() {
        clog.Info("Hey, over here. I got some news for you.")

        num := 5
        clog.Debug("There is total %d slices of toast.", num)

        clog.Warning("Do not put anything but toast into the toaster.")

        clog.Error("Error! Your toast has been burned!")
    }

    ```
    ![](https://github.com/shadowbeast5478/clog/blob/master/console.png)

- Example 2 : Working with testing
    ```go
    package sort

    import (
        "testing"
        "github.com/shadowbeast5478/clog"
    )

    func compare(slice1 []int, slice2 []int) bool{
        if len(slice1) != len(slice2) {
            return false
        }

        for i := 0; i < len(slice1); i++ {
            if slice1[i] != slice2[i] {
                return false
            }
        }
        return true
    }

    func TestMergeSort1(t *testing.T) {
        numbers := []int {100, 5, 80, 15, 20, 125, 235, 7, 3, 1, 75, 135, 165}
        sorted := MergeSort(numbers)
        answer := []int {1, 3, 5, 7, 15, 20, 75, 80, 100, 125, 135, 165, 235}
        result := compare(sorted, answer)

        if result != true {
            t.Error()
            clog.Error("MergeSort test failed.")
        } else {
            t.Log()
            clog.Info("First test passed.")
        }
    }

    func TestMergeSort2(t *testing.T) {
        numbers := []int {101, 10, 87, 13, 20, 17, 6, 0, 55, 168, 139}
        sorted := MergeSort(numbers)
        answer := []int {0, 6, 10, 13, 17, 20, 55, 87, 101, 139, 168}
        result := compare(sorted, answer)

        if result != true {
            t.Error()
            clog.Error("MergeSort test failed.")
        } else {
            t.Log()
            clog.Info("Second test passed.")
        }
    }

    ```
    ![](https://github.com/shadowbeast5478/clog/blob/master/console_go_test.png)