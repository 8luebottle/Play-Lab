package TestCode

import "strings"

func Iterator(character string, iterateTimes int) (result string) {
    result = character
    if iterateTimes > 1 {
        result = strings.Repeat(character, iterateTimes)
    }
    return  result
}