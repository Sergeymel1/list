package slice

import "fmt"

type Slice struct {
	sl []int64
}

func (s *Slice) Add(data int64) (index int64) {
	s.sl = append(s.sl, data)
	return int64(len(s.sl) - 1)
}

func (s *Slice) Delete(index int64) (ok bool) {
	if index > int64(len(s.sl)-1) {
		fmt.Println("Invalid index")
		return false
	}

	copy(s.sl[index:], s.sl[index+1:])
	s.sl[len(s.sl)-1] = 0
	s.sl = s.sl[:len(s.sl)-1]
	return true
}

func (s *Slice) Print() {
	fmt.Println(s.sl)
}

func (s *Slice) SortInc() {
	for i := 0; i < len(s.sl); i++ {
		max := s.sl[0]
		iMax := 0
		for j := 0; j < len(s.sl)-i; j++ {
			if s.sl[j] > max {
				max = s.sl[j]
				iMax = j
			}
		}
		s.sl[len(s.sl)-i-1], s.sl[iMax] = s.sl[iMax], s.sl[len(s.sl)-i-1]
	}
}

func (s *Slice) SortDec() {
	for i := 0; i < len(s.sl); i++ {
		min := s.sl[0]
		iMin := 0
		for j := 0; j < len(s.sl)-i; j++ {
			if s.sl[j] < min {
				min = s.sl[j]
				iMin = j
			}
		}
		s.sl[len(s.sl)-i-1], s.sl[iMin] = s.sl[iMin], s.sl[len(s.sl)-i-1]
	}
}
