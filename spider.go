package tinyspider

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Spider struct {
	URL string

	doc *goquery.Document

	tasks []func(doc *goquery.Document)

	beforeReqCallbacks []func(req *http.Request) error
	afterReqCallbacks  []func(res *http.Response) error
}

func NewSpider() *Spider {
	return &Spider{}
}

func (s *Spider) Get(url string) error {
	cli := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	for _, cb := range s.beforeReqCallbacks {
		if err := cb(req); err != nil {
			return err
		}
	}

	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	for _, cb := range s.afterReqCallbacks {
		if err := cb(res); err != nil {
			return err
		}
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	s.URL = url
	s.doc = doc

	s.run()

	return nil
}

func (s *Spider) run() {
	for _, task := range s.tasks {
		task(s.doc)
	}
}

func (s *Spider) AddTask(task func(doc *goquery.Document)) {
	s.tasks = append(s.tasks, task)
}

func (s *Spider) BeforeReq(cb func(req *http.Request) error) {
	s.beforeReqCallbacks = append(s.beforeReqCallbacks, cb)
}

func (s *Spider) AfterReq(cb func(res *http.Response) error) {
	s.afterReqCallbacks = append(s.afterReqCallbacks, cb)
}
