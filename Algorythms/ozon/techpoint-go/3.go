package main

import "fmt"

func main() {
    var usrAmount, reqAmount uint64
    fmt.Scanf("%v %v", &usrAmount, &reqAmount)

    var lastGlob uint64 = 0
    var iterator uint64 = 1

    m := make(map[uint64]uint64, usrAmount)

    var i uint64

    for i = 0; i < reqAmount; i++ {
        var reqType int
        var uid uint64
        fmt.Scanf("%v %v", &reqType, &uid)

        // fmt.Println(iterator)

        if reqType == 1 { // send a message to uid (if uid == 0 => send to all)
            if uid == 0 {
                lastGlob = iterator
            } else {
                m[uid] = iterator
            }
            iterator++;
        } else if reqType == 2 { // show last msg
            var lastUserNotification uint64
            lastUserNotification, ok := m[uid]; 
            if !ok {
                fmt.Println(lastGlob)
                continue
            }
            if lastGlob > lastUserNotification {
                fmt.Println(lastGlob)
            } else {
                fmt.Println(lastUserNotification)
            }
        }
    }

