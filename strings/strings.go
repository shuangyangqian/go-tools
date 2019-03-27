package strings

type ByPiYin []string

func (s ByPiYin) Len() int      { return len(s) }
func (s ByPiYin) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPiYin) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}
