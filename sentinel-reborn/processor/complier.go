//janai egeonu

package processor

func Complier(text string) string {

	text = Casing(text)

	text = Base(text)
	text = FixArticles(text)
	text = FixPunct(text)

	return fixQuote(text)
}
