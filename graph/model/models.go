package model

func (s *AggragateStats) UpdateCalculatedFields() {
	s.Total = s.Won + s.Lost
	s.Ratio = float64(s.Won) / float64(s.Total)
}
