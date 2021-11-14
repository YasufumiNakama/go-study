package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	/* forEachNodeの出力をパースできることを保証するためのテスト */
	var tests = []struct {
		html string
	}{
		{`
		<html lang="en">
			<head>
				<title>Sample</title>
			</head>
			<body>
				<!-- comment -->
				<h1>This is Sample</h1>
				<p><a href="https://golang.org/">Sample Link</a></p>
				<img></img>
			</body>
		</html>
		`},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Error(err)
		}
		writer = new(bytes.Buffer)
		forEachNode(doc, startElement, endElement)
		got := writer.(*bytes.Buffer).String()
		if _, err := html.Parse(strings.NewReader(got)); err != nil {
			t.Errorf("forEachNodeの出力をパースできません %v", err)
		}
	}
}
