/*
	https://leetcode.com/problems/design-browser-history/description/

		You have a browser of one tab where you start on the homepage and you can
			visit another url, get back
	in the history number of steps or move forward in the history number of steps.

	Implement the BrowserHistory class:

		BrowserHistory(string homepage) Initializes the object with the homepage of
			the browser.
		void visit(string url) Visits url from the current page. It clears up all the
			forward history.
		string back(int steps) Move steps back in history. If you can only return x
			steps in the history
	and steps > x, you will return only x steps. Return the current url after
		moving back in history at most steps.
		string forward(int steps) Move steps forward in history. If you can only
			forward x steps in the history
	and steps > x, you will forward only x steps. Return the current url after
		forwarding in history at most steps.

*/

//nolint:revive //it's ok
package solutions

type DDL struct {
	Val     string
	Back    *DDL
	Forward *DDL
}

type BrowserHistory struct {
	curr *DDL
}

// NewBrowserHistory should call Constructor to pass LeetCode tests
func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{curr: &DDL{Val: homepage}}
}

func (h *BrowserHistory) Visit(url string) {
	oldCurr := h.curr
	curr := &DDL{Val: url, Back: oldCurr}
	oldCurr.Forward = curr

	h.curr = curr
}

func (h *BrowserHistory) Back(steps int) string {
	for h.curr.Back != nil && steps > 0 {
		h.curr = h.curr.Back
		steps--
	}

	if h.curr != nil {
		return h.curr.Val
	}
	return ""
}

func (h *BrowserHistory) Forward(steps int) string {
	for h.curr.Forward != nil && steps > 0 {
		h.curr = h.curr.Forward
		steps--
	}

	if h.curr != nil {
		return h.curr.Val
	}
	return ""
}
