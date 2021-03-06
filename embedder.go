// Tool for generating Go source files containing embedded resources,
// leveraging the library at http://github.com/cratonica/embed
//
// http://github.com/cratonica/embedder
package main

import (
	"fmt"
	"github.com/overlordtm/embed"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: embedder [package name] [path]")
		return
	}

	resourceMap, err := embed.CreateFromFiles(os.Args[2])
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}
	packed, err := embed.Pack(resourceMap)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}

	fmt.Print(embed.GenerateGoCode(os.Args[1], "Resources", packed))
}
