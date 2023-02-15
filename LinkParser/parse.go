package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)



type Link struct {
    Href string
    Text string
}


func Parse(r io.Reader) ([]Link, error) {

    fmt.Println("my parser called")
    doc, err := html.Parse(r)

    if err != nil {
        return nil, err
    }

    nodes := linkNodes(doc)
    var links []Link
    for _, node := range nodes {
        links = append(links, buildLink(node))
    }

    fmt.Println(links)
    return links, nil
}

func buildLink(node *html.Node) Link {
    var link Link
    for _, attr := range node.Attr {
        if attr.Key == "href" {
            link.Href = attr.Val
            break
        }
    }

    link.Text = extractText(node)
    return link
}

func extractText (n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data
    }

    if n.Type != html.ElementNode {
        return ""
    }

    var ret string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret += extractText(c) + " "
    }
    return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {

    if n.Type == html.ElementNode && n.Data == "a" {
        return []*html.Node{n}
    }
    var ret []*html.Node

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret = append(ret, linkNodes(c)...)
    }
    return ret
}


