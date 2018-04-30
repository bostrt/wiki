package main

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestCamelCaseText(t *testing.T) {
    assert := assert.New(t)

    body := []byte("WikiLink")
    output := WikiPageFilter(body, "example.com/view/")
    assert.Equal("[WikiLink](example.com/view/WikiLink)", string(output))
}

func TestCamelLinkText(t *testing.T) {
    assert := assert.New(t)

    body := []byte("[AnExampleLink](http://example.com)")
    output := WikiPageFilter(body, "example.com/view/")
    assert.Equal("[AnExampleLink](http://example.com)", string(output))
}
