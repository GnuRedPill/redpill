package cut

import (
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/go-ego/gse"
)

const (
	NOUN     = "n" //名词
	NOUNTYPE = 1
	STR      = "x" //字符串
	STRTYPE  = 2
	DESC     = "l" //简称或者同义词
	DESCTYPE = 3
)

type Segmenter struct {
	seg    gse.Segmenter
	IsInit bool
}

func NewSegmenter() *Segmenter {
	seg := gse.Segmenter{}
	err := seg.LoadDict()
	if err != nil {
		panic(err)
	}
	log.SetFlags(log.Lshortfile & log.Ldate)
	return &Segmenter{seg: seg}
}

type GseToken struct {
	Text string
	Freq int
	Pos  string
}

func (s *Segmenter) AddToken(gs []GseToken) {
	for _, v := range gs {
		err := s.seg.AddToken(v.Text, float64(v.Freq), v.Pos)
		if err != nil {
			log.Println(err, v)
		}
	}
	err := s.seg.LoadDict()
	if err != nil {
		log.Println(err)
	}
}
func (s *Segmenter) UpdateToken(gs []GseToken) {
	for _, v := range gs {
		err := s.seg.ReAddToken(v.Text, float64(v.Freq), v.Pos)
		if err != nil {
			log.Println(err)
		}
	}
	err := s.seg.LoadDict()
	if err != nil {
		log.Println(err)
	}
}
func (s *Segmenter) AddStop(texts []string) {
	for _, v := range texts {
		s.seg.AddStop(v)
	}
	err := s.seg.LoadDict()
	if err != nil {
		log.Println(err)
	}
}
func (s *Segmenter) Reload() {
	s.seg = gse.Segmenter{}
	err := s.seg.LoadDict()
	if err != nil {
		log.Println(err)
	}
}

func (s *Segmenter) CutVisit(text string) map[string]int {
	a := s.seg.Pos(text)
	a = s.seg.TrimPos(a)
	visit := map[string]int{}
	for _, v := range a {
		if strings.Contains("nxl", v.Pos) && !regexp.MustCompile(`^\d+$`).MatchString(v.Text) {
			visit[v.Text] += 1
		}
	}
	return visit
}

func (s *Segmenter) Cut(text string) []string {
	m := s.CutVisit(text)
	pairs := make([]struct {
		Key   string
		Value int
	}, len(m))

	i := 0
	for k, v := range m {
		pairs[i] = struct {
			Key   string
			Value int
		}{k, v}
		i++
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	keys := make([]string, len(m))
	for i, pair := range pairs {
		keys[i] = pair.Key
	}
	return keys
}
