package types

type StrArray []string

func (s StrArray) SortAlphabeticalAsc() StrArray {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				tmp := s[i]
				s[i] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}
