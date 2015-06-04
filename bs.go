package bs

import (
	"fmt"
	"strconv"

	"github.com/golangplus/math"

	hf "github.com/gohtml/html"
)

type Bootstrap struct {
	*hf.Html
}

func Container(fluid bool, children ...hf.Node) *hf.Element {
	c := hf.DIV()

	if fluid {
		c.Attr("class", "container-fluid")
	} else {
		c.Attr("class", "container")
	}

	return c.Child(children...)
}

func Row(children ...hf.Node) *hf.Element {
	return hf.DIV().Attr("class", "row").Child(children...)
}

func Col(n int, sz string, children ...hf.Node) *hf.Element {
	return hf.DIV().Attr("class", fmt.Sprintf("col-%s-%d", sz, n)).Child(children...)
}

func Tabs(tp string, active int, tabs ...hf.Node) *hf.Element {
	for i, tab := range tabs {
		li := hf.LI(hf.A("#", tab))

		if i == active {
			li.Attr("class", "active")
		}

		tabs[i] = li
	}
	return hf.UL(tabs...).Attr("class", fmt.Sprintf("nav nav-%s", tp))
}

const (
	fixedNone = iota
	fixedTop
	fixedBottom
)

type Navbar struct {
	inverse bool
	fixed   int

	brand   *hf.Element
	left    []*hf.Element
	actives []bool
	right   []*hf.Element
}

func INPUT(tp, name, value string) *hf.Void {
	return hf.INPUT(tp, name, value).AddClass("form-control")
}

func TEXTAREA(name, value string, children ...hf.Node) *hf.Element {
	return hf.TEXTAREA(name, value, children...).AddClass("form-control")
}

func BUTTON(isDefault bool, children ...hf.Node) *hf.Element {
	btn := hf.BUTTON(children...).AddClass("btn")
	if isDefault {
		btn.AddClass("btn-default")
	}
	return btn
}

func Glyphicon(name string) *hf.Element {
	return hf.SPAN().AddClass("glyphicon", "glyphicon-"+name)
}

func InputGroupBtn(children ...hf.Node) *hf.Element {
	return hf.SPAN(children...).AddClass("input-group-btn")
}

func Alert(tp string, dismissible bool, children ...hf.Node) *hf.Element {
	t := hf.DIV().AddClass("alert", "alert-"+tp).Attr("role", "alert")

	if dismissible {
		t.AddClass("alert-dismissible")

		t.Child(
			hf.BUTTON().AddClass("close").Attr("data-dismiss", "alert").Attr("aria-label", "Close").Child(
				hf.SPAN(hf.TIMES).Attr("aria-hidden", "true"),
			),
		)
	}

	return t.Child(children...)
}

func FormGroup(children ...hf.Node) *hf.Element {
	return hf.DIV(children...).AddClass("form-group")
}

func HelpBlock(children ...hf.Node) *hf.Element {
	return hf.P(children...).AddClass("help-block")
}

func Panel(state string, children ...hf.Node) *hf.Element {
	return hf.DIV(children...).AddClass("panel", "panel-"+state)
}

func PanelHeading(children ...hf.Node) *hf.Element {
	return hf.DIV(children...).AddClass("panel-heading")
}

func PanelBody(children ...hf.Node) *hf.Element {
	return hf.DIV(children...).AddClass("panel-body")
}

func UListGroup(children ...hf.Node) *hf.Element {
	return hf.UL(children...).AddClass("list-group")
}

func OListGroup(children ...hf.Node) *hf.Element {
	return hf.OL(children...).AddClass("list-group")
}

func ListGroupItem(children ...hf.Node) *hf.Element {
	return hf.LI(children...).AddClass("list-group-item")
}

func Pagination(urlFunc func(p int) string, curPage, totalPages, maxLeft, maxRight int) *hf.Element {
	ul := hf.UL().AddClass("pagination")

	if curPage > 1 {
		ul.Child(hf.LI(
			hf.A(urlFunc(curPage-1), hf.LAQUO),
		))

		mn := mathp.MaxI(1, curPage-maxLeft)
		for i := mn; i < curPage; i++ {
			ul.Child(hf.LI(
				hf.A(urlFunc(i), hf.T(strconv.Itoa(i))),
			))
		}
	}
	ul.Child(hf.LI(hf.A("#",
		hf.T(strconv.Itoa(curPage)),
		hf.SPAN(hf.T("(current)")).AddClass("sr-only"),
	)).AddClass("active"))
	if curPage < totalPages {
		mx := mathp.MinI(totalPages, curPage+maxRight)
		for i := curPage + 1; i <= mx; i++ {
			ul.Child(hf.LI(
				hf.A(urlFunc(i), hf.T(strconv.Itoa(i))),
			))
		}

		ul.Child(hf.LI(
			hf.A(urlFunc(curPage+1), hf.RAQUO),
		))
	}

	return ul
}

func PageHeader(children ...hf.Node) *hf.Element {
	return hf.DIV(children...).AddClass("page-header")
}

/* Navbar */

func NewNavbar(inverse bool) *Navbar {
	return &Navbar{inverse: inverse}
}

func (nb *Navbar) FixedTop() *Navbar {
	nb.fixed = fixedTop
	return nb
}

func (nb *Navbar) FixedBottom() *Navbar {
	nb.fixed = fixedBottom
	return nb
}

func (nb *Navbar) Brand(b *hf.Element) *Navbar {
	nb.brand = b
	return nb
}

func (nb *Navbar) Left(active bool, tag *hf.Element) *Navbar {
	nb.left = append(nb.left, tag)
	nb.actives = append(nb.actives, active)
	return nb
}

func (nb *Navbar) LeftLink(active bool, href string, text string) *Navbar {
	return nb.Left(active, hf.A(href, hf.T(text)))
}

func (nb *Navbar) Right(tag *hf.Element) *Navbar {
	nb.right = append(nb.right, tag)
	return nb
}

func (nb *Navbar) RightLink(href string, text string) *Navbar {
	return nb.Right(hf.A(href, hf.T(text)))
}

func (nb *Navbar) RightText(text string) *Navbar {
	return nb.Right(hf.DIV(hf.T(text)).AddClass("navbar-text"))
}

func (nb *Navbar) FORM(method, action string, children ...hf.Node) *hf.Element {
	return hf.FORM(method, action, children...).AddClass("navbar-form")
}

func (nb *Navbar) AsTag() *hf.Element {
	tag := hf.NAV().AddClass("navbar")
	if nb.inverse {
		tag.AddClass("navbar-inverse")
	} else {
		tag.AddClass("navbar-default")
	}

	switch nb.fixed {
	case fixedTop:
		tag.AddClass("navbar-fixed-top")
	case fixedBottom:
		tag.AddClass("navbar-fixed-bottom")
	}

	c := Container(true).AddClass()

	if nb.brand != nil {
		c.Child(hf.DIV().AddClass("navbar-header").Child(nb.brand.AddClass("navbar-brand")))
	}

	if len(nb.left) > 0 || len(nb.right) > 0 {
		div := hf.DIV().AddClass("collapse", "navbar-collapse")

		if len(nb.left) > 0 {
			ul := hf.UL().AddClass("nav", "navbar-nav")
			for i, el := range nb.left {
				li := hf.LI(el)
				if nb.actives[i] {
					li.AddClass("active")
				}
				ul.Child(li)
			}
			div.Child(ul)
		}

		if len(nb.right) > 0 {
			ul := hf.UL().AddClass("nav", "navbar-nav", "navbar-right")
			for _, el := range nb.right {
				ul.Child(hf.LI(el))
			}
			div.Child(ul)
		}

		c.Child(div)
	}

	tag.Child(c)

	return tag
}

func New() *Bootstrap {
	bs := Bootstrap{
		Html: hf.HTML("en"),
	}

	bs.Css("http://netdna.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css")

	return &bs
}
