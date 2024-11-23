const modeAdd = "add"
const modeSet = "set"


class Segment {
    constructor(start, intensity) {
        this.start = start;
        this.intensity = intensity;
    }

    toString() {
        return `[${this.start}, ${this.intensity}]`;
    }
}

class IntensitySegments {

    constructor() {
        this.segments = []
    }

    firstSegment() {
        return this.segments[0].start
    }

    lastSegment() {
        return this.segments[this.segments.length - 1].start
    }

    add(from, to, amount) {
        if (amount == 0 || from >= to) return

        if (this.segments.length == 0 || from > this.lastSegment()) {
            this.segments.push(new Segment(from, amount), new Segment(to, 0))
            return
        }

        if (to < this.firstSegment()) {
            this.segments.unshift(new Segment(from, amount), new Segment(to, 0))
            return
        }
        this.handleOverlapping(from, to, amount, modeAdd)
    }

    set(from, to, amount) {
        if (from >= to) return

        if (this.segments.length == 0 || from > this.lastSegment() || to < this.firstSegment()) {
            this.add(from, to, amount)
            return
        }

        this.handleOverlapping(from, to, amount, modeSet)
    }

    handleOverlapping(from, to, amount, mode) {
        let left = this.findPosition(from)
        let right = this.findPosition(to)

        if (to > this.lastSegment()) {
            this.segments.push(new Segment(to, 0))
        } else if (this.segments[right].start != to) {
            this.segments.splice(right, 0, new Segment(to, this.segments[right - 1].intensity))
        }

        for (let index = left; index < right && index < this.segments.length; index++) {
            if (this.segments[index].start < from || this.segments[index].start >= to) {
                continue
            }
            if (mode === modeSet) {
                this.segments[i].intensity = amount
            } else if (mode === modeAdd) {
                this.segments[index].intensity += amount
            }
        }

        if (from < this.firstSegment()) {
            this.segments.unshift(new Segment(from, amount))
        } else if (this.segments[left].start != from) {
            let newSegment = new Segment(from, amount)
            if (mode === modeAdd) {
                newSegment.intensity += this.segments[left - 1].intensity
            }
            this.segments.splice(left, 0, newSegment)
        }
        this.merge()
    }

    merge() {
        if (this.segments.length == 0) return

        while (this.segments.length > 0 && this.segments[0].intensity == 0) {
            this.segments.shift()
        }

        let writeIndex = 0;

        for (let readIndex = 1; readIndex < this.segments.length; readIndex++) {
            if (this.segments[readIndex].intensity != this.segments[writeIndex].intensity) {
                writeIndex++
                this.segments[writeIndex] = this.segments[readIndex]
            }
        }
        this.segments.length = writeIndex + 1;
    }

    findPosition(target) {
        let left = 0
        let right = this.segments.length - 1
        while (left <= right) {
            let mid = left + Math.floor((right - left) / 2)
            if (this.segments[mid].start == target) {
                return mid
            } else if (this.segments[mid].start < target) {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return left
    }

    toString() {
        return '[' + this.segments.join(',') + ']'
    }
}

module.exports = IntensitySegments