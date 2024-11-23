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