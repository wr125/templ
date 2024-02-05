#A s imple boilerplate template which inculdes htmx, templ, echo(go framework) and tailwind (css). Hosted with an image.


import "github.com/wr125/templ/templates/layout"

templ HomeIndex(
    title,
    username string,
    fromProtected bool,
    isError bool,
	errMsgs, sucMsgs []string,
    cmp templ.Component,
    ) {
	@layout.Base(title, username, fromProtected, isError, errMsgs, sucMsgs) {
		@cmp
	}
}
 