// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

package main

import (
      _ "fmt"
      "io/ioutil"
)

// func main() {
//         fileContents, err := ioutil.ReadFile("first-post.txt")
//         if err != nil {
//             // A common use of `panic` is to abort if a function returns an error
//             // value that we don’t know how to (or want to) handle. This example
//             // panics if we get an unexpected error when creating a new file.
//             panic(err)
//         }
//         fmt.Print(string(fileContents))
// }

func main() {
        bytesToWrite := []byte(`Mega Bacon Ipsum Tempor irure proident, pork do porchetta exercitation minim est biltong rump corned beef ad t-bone. Chicken ullamco pastrami ham hock chislic shankle. Anim spare ribs chislic eu. Anim turducken in jowl enim hamburger cupim mollit.
		Porchetta tenderloin short ribs consequat sint, ut est bresaola ground round reprehenderit excepteur. Do pancetta est, chuck dolore nisi doner burgdoggen tempor occaecat brisket nulla hamburger. Chislic labore lorem officia tri-tip, et dolore salami laboris pastrami tenderloin ham hock short loin swine. Ipsum officia aute drumstick culpa. Minim labore salami nostrud spare ribs meatloaf. Chuck swine kielbasa, nulla laboris adipisicing beef ribs t-bone pig alcatra pork loin fatback sausage drumstick.
		Andouille occaecat ball tip, sunt ullamco culpa eu ut dolor. Shank meatball spare ribs, tri-tip andouille bresaola tongue. Ut consectetur t-bone sirloin salami strip steak short loin quis. Venison strip steak consequat pig dolore in quis ut laborum nisi nostrud aliquip pastrami elit. Frankfurter lorem shankle ut elit minim.
		Landjaeger chuck tenderloin ipsum bacon, porchetta officia laborum short ribs ad beef buffalo burgdoggen. Duis beef ribs pork chop qui pancetta minim sausage proident pork loin eiusmod. Picanha enim ipsum eiusmod. Corned beef dolor laborum ut pork exercitation ball tip buffalo swine tongue picanha chicken ribeye sunt jowl. Beef cillum chuck excepteur sunt. Beef ribs nisi magna excepteur aute.
		Cupim pork leberkas quis drumstick. Esse shoulder turducken doner sunt, nulla pastrami dolore. Laboris ex short ribs filet mignon. Ex anim swine ipsum non pancetta magna aliqua chuck pork chop rump nostrud veniam nulla shank. Sunt excepteur leberkas swine, tongue proident rump turducken sausage brisket drumstick turkey.
		`)
        err := ioutil.WriteFile("template.tmpl", bytesToWrite, 0644)
        if err != nil {
            panic(err)
        }
}


// Build Program 
// $ go build 

// Run latest Build 
// ./makesite