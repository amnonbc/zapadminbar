package htmlparse

import (
	"golang.org/x/net/html"
	"io"
)

// <li id="wp-admin-bar-my-account" class="menupop with-avatar"><a class="ab-item" aria-haspopup="true" href="https://shopstewards.net/wp-admin/profile.php">How are you, BenR?<img alt='' src='https://secure.gravatar.com/avatar/c6b616ada6b9f32b0bfbbc18bdcb2a24?s=26&#038;d=mm&#038;r=g' srcset='https://secure.gravatar.com/avatar/c6b616ada6b9f32b0bfbbc18bdcb2a24?s=52&amp;d=mm&amp;r=g 2x' class='avatar avatar-26 photo' height='26' width='26' /></a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-user-actions" class="ab-submenu">

func Process(r io.Reader) (hasToolbar bool, links []string, err error) {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return
			}
			return hasToolbar, links, z.Err()
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}
			if "li" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "id" && attr.Val == "wp-admin-bar-my-account" {
						hasToolbar = true
					}

				}
			}

		}
	}
	return
}
