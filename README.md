# Authedia Developers

Using the Authedia API in Go. You can find the documentation [here](https://github.com/Authedia/Developers)

### Example
```
package main

import "Developers/authedia"

func main() {

    your_api_key := "..."

    authedia.Wrap(
        your_api_key,
        "example_media/PNG.png",
        "test.png",
    )

    response := authedia.Verify(
        your_api_key,
        "test.png",
    )

}
```
