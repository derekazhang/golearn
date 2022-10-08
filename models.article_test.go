package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()
	if len(alist) != len(articleList) {
		t.Fail()
	}
	for i, v := range alist {
		articleAtI := articleList[i]
		if v.Content != articleAtI.Content || v.ID != articleAtI.ID || v.Title != articleAtI.Title {
			t.Fail()
			break
		}
	}
}
