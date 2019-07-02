package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML doc.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML doc and return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil

	// 1 Find <a> nodes in doc
	// 2 for each link node...
	//		build a Link
	// 3 return the Links
}

// *Main parsing process, using DFS
// takes in a root(Node), returns all html sub-Nodes that are links (<a>)
func linkNodes(n *html.Node) []*html.Node {
	// base case; if node passed in is a link, then just return it
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// recursively append link nodes that are found to our return variable
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// Takes node and turns it into a Link for us
func buildLink(n *html.Node) Link {
	var ret Link
	// iterate over entire slice of attributes
	for _, attr := range n.Attr {
		// get the link
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)

	return ret
}

// returns a string representing all the text inside a specific node
func text(n *html.Node) string {
	// base case
	if n.Type == html.TextNode {
		return n.Data
	}

	// if not an element node we don't care about contents
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	// for every child, get text of that child, and add it to return val
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	// clean up return text by only grabbing text (no whitespace/returns)
	// and joining them into a string seperated by only one space
	return strings.Join(strings.Fields(ret), " ")
}

// takes in an html node and finds/prints sub nodes
func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
