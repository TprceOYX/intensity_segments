# Intensity-Segments
## Introduction
This repository provides a solution for IntensitySegements in Golang and JavaScript.

I am a JavaScript newbee. The Javascript code is basically translated from Golang, so there are no unit tests.
## Example
### Golang
```go
func main() {
    s := segments.NewIntensitySegments()
    fmt.Println(s.ToString()) //  Should print: "[]"
    s.Add(10, 30, 1)
    fmt.Println(s.ToString()) //  Should print: "[[10,1],[30,0]]"
    s.Set(30, 40, 1)
    fmt.Println(s.ToString()) //  Should print: "[[10,1],[40,0]]"
    s.Add(30, 40, 1)
    fmt.Println(s.ToString()) //  Should print: "[[10,1],[30,2],[40,0]]"
    s.Set(30, 40, 0)
    fmt.Println(s.ToString()) //  Should print: "[[10,1],[30,0]]"
}
```
### JavaScript
```javascript
const IntensitySegments = require("./intensity-segments.js");

const s1 = new IntensitySegments();
console.log(s1.toString());; // Should be "[]"

s1.add(10, 30, 1);
console.log(s1.toString()); // Should be: "[[10,1],[30,0]]"
s1.add(20, 40, 1);
console.log(s1.toString()); // Should be: "[[10,1],[20,2],[30,1],[40,0]]"
s1.add(10, 40, -2);
console.log(s1.toString()); // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"

// Another example sequence:
const s2 = new IntensitySegments();
console.log(s2.toString()); // Should be "[]"
s2.add(10, 30, 1);
console.log(s2.toString()); // Should be "[[10,1],[30,0]]"
s2.add(20, 40, 1);
console.log(s2.toString()); // Should be "[[10,1],[20,2],[30,1],[40,0]]"
s2.add(10, 40, -1);
console.log(s2.toString()); // Should be "[[20,1],[30,0]]"
s2.add(10, 40, -1);
console.log(s2.toString()); // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
```
## Run Test in Golang
```bash
$ go test . -v
```