package example

import "github.com/codecov/example-go/awesome"

var result string

func Setup(other bool) {

    // Comment

    if other {
        result = "hello world"
    }


    result = awesome.Smile()

}

func GetResult() string {

    /*
    Comment
    */

    return result

}
