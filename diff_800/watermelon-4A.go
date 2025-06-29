package main

import "fmt"

func checkMelonWeight(){
    var input int;
    fmt.Scanln(&input);
    if input%2==0 && input > 2{
        fmt.Print("YES");
    }else{
        fmt.Print("NO");
    }

}

