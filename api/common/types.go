package common

func (s *StringReply) StringReply() string {
	if s == nil {
		return ""
	}
	return s.Body
}
