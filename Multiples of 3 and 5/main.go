package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    
    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        N := int64(nTemp)
        
       
        if N < 3 {
            continue
        }
        
        if N < 5 {
           fmt.Println(3)
           continue
        }
        
        if N < 6 {
           fmt.Println(8)
           continue
        }
        
        if N < 15 {
            fmt.Println(AP(N, 3) + AP(N, 5))
            continue
        }
        
        fmt.Println(AP(N, 3) + AP(N, 5) - AP(N, 15))
    }
    
}


//AP -> returns the sum of the Arithmethic Progression
//
// N -> The limit of the last element
// 
// d -> the common difference
func AP(N int64, d int64) int64 {
    
    n := (N - 1)/d
    a1 := d
    l := n * d
     
   // S = n * (first + last)/2
    S := n * (a1 + l)/2
    
    return S
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
