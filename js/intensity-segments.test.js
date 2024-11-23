const IntensitySegments = require("./intensity-segments.js");

const s1 = new IntensitySegments();
test("empty", () => {
    expect(s1.toString()).toBe("[]")
})