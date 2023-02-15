package html

/*
	repackage goquery to make it more convenient to use
*/

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	*goquery.Document
}

type Selection struct {
	*goquery.Selection
}

func NewDocumentFromReader(reader io.Reader) (*Document, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	return &Document{doc}, nil
}

func (doc *Document) FindByClass(class string) *Selection {
	return &Selection{doc.Find("." + class)}
}

func (doc *Document) FindByTag(tag string) *Selection {
	return &Selection{doc.Find(tag)}
}

func (doc *Document) FindById(id string) *Selection {
	return &Selection{doc.Find("#" + id)}
}

func (doc *Document) FindTitle() *Selection {
	return &Selection{doc.Find("title")}
}

func (doc *Document) FindIcon() *Selection {
	return &Selection{doc.Find("link[rel='icon']")}
}

func (doc *Document) FindIconUrl() string {
	return doc.FindIcon().AttrOr("href", "")
}

func (doc *Document) FindTitleValue() string {
	return doc.FindTitle().Text()
}