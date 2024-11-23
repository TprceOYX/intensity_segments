package intensity_segments

import (
	"fmt"
	"strings"
)

const (
	modeAdd = "add"
	modeSet = "set"
)

type segment struct {
	start     int // segment start point
	intensity int
}

func (s *segment) ToString() string {
	return fmt.Sprintf("[%d,%d]", s.start, s.intensity)
}

type IntensitySegments struct {
	segments []*segment // "segment.start" is in ascending order
}

func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		segments: []*segment{},
	}
}

func (s *IntensitySegments) firstSegment() int {
	return s.segments[0].start
}

func (s *IntensitySegments) lastSegment() int {
	return s.segments[len(s.segments)-1].start
}

// Add specific intensity to segments in range [from, to), if segments does not exist, add new segments.
func (s *IntensitySegments) Add(from, to, amount int) {
	if amount == 0 || from >= to { // invalid args
		return
	}
	// init or 'from' greater than right border
	if len(s.segments) == 0 || from > s.lastSegment() {
		// push to tail
		s.segments = append(s.segments, &segment{start: from, intensity: amount})
		s.segments = append(s.segments, &segment{start: to})
		return
	}
	// 'to' less than left border
	if to < s.firstSegment() {
		// put to head
		segments := []*segment{
			{start: from, intensity: amount},
			{start: to},
		}
		s.segments = append(segments, s.segments...)
		return
	}

	s.handleOverlapping(from, to, amount, modeAdd)
}

// Set specific intensity to segments in range [from, to), if segments does not exist, add new segments.
func (s *IntensitySegments) Set(from, to, amount int) {
	if from >= to { // invalid args
		return
	}
	// init or no overlapping
	if len(s.segments) == 0 || from > s.lastSegment() || to < s.firstSegment() {
		s.Add(from, to, amount)
		return
	}

	// overlapping
	s.handleOverlapping(from, to, amount, modeSet)
}

// handleOverlapping handling scenarios where new segments may overlap with existing segments.
func (s *IntensitySegments) handleOverlapping(from, to, amount int, mode string) {
	left := s.findPosition(from)
	right := s.findPosition(to)

	// handle the 'to' first, because it may use the original data on its left
	// which may be changed in subsequent operations
	if to > s.lastSegment() { // out for right border, append a new segment in tail
		s.segments = append(s.segments, &segment{start: to})
	} else if s.segments[right].start != to {
		s.insertSegment(right, &segment{ // in border, insert a new segment
			start:     to,
			intensity: s.segments[right-1].intensity,
		})
	}

	// handle segments start in [from, to) next
	for i := left; i < right && i < len(s.segments); i++ {
		if s.segments[i].start < from || s.segments[i].start >= to {
			continue
		}
		if mode == modeSet {
			s.segments[i].intensity = amount // set same intensity to merge
		} else if mode == modeAdd {
			s.segments[i].intensity += amount // add intensity
		}
	}

	// handle the 'from' last
	if from < s.firstSegment() { // out for left border, put a new segment in head
		s.segments = append([]*segment{{start: from, intensity: amount}}, s.segments...)
	} else if s.segments[left].start != from { // in border, insert a new segment
		newSegment := &segment{
			start:     from,
			intensity: amount,
		}
		if mode == modeAdd {
			newSegment.intensity += s.segments[left-1].intensity
		}
		s.insertSegment(left, newSegment)
	}

	s.merge()
}

// merge removes prefix segments which intensity is zero and merges continuous segments with same intensity.
func (s *IntensitySegments) merge() {
	if len(s.segments) == 0 {
		return
	}
	// remove prefix zero
	zeroCount := 0
	for i := 0; i < len(s.segments) && s.segments[i].intensity == 0; i++ {
		zeroCount++
	}
	s.segments = s.segments[zeroCount:]
	// merge continuous segments with same intensity
	writeIndex := 0
	for i := 1; i < len(s.segments); i++ {
		if s.segments[i].intensity != s.segments[writeIndex].intensity {
			writeIndex++
			s.segments[writeIndex] = s.segments[i]
		}
	}
	s.segments = s.segments[:writeIndex+1]
}

// findPosition find a max index which satisfies: when x < index, s.segments[x].start < target.
func (s *IntensitySegments) findPosition(target int) int {
	left, right := 0, len(s.segments)-1
	for left <= right {
		mid := left + (right-left)/2
		if s.segments[mid].start == target {
			return mid
		} else if s.segments[mid].start < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// insertSegment inserts a new segment at specific index.
func (s *IntensitySegments) insertSegment(index int, newSegment *segment) {
	s.segments = append(s.segments, nil)           // add a empty slot
	copy(s.segments[index+1:], s.segments[index:]) // move segments after index one position back
	s.segments[index] = newSegment                 // insert segment at index position
}

func (s *IntensitySegments) ToString() string {
	sb := strings.Builder{}
	for index, segment := range s.segments {
		_, _ = sb.WriteString(segment.ToString())
		if index != len(s.segments)-1 {
			_ = sb.WriteByte(',')
		}
	}
	return fmt.Sprintf("[%s]", sb.String())
}
