package main

import "bookutil/author"

func main() {
	//create an author instance.
	authorInstance := author.NewAuthor("erick ndeto", "erickndeto202@gmail.com")
	//write and review a chapter.
	chapterTitle := "Introduction to Go modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability."

	authorInstance.WriteChapter(chapterTitle, chapterContent)
	authorInstance.ReviewChapter(chapterTitle, "This chapter looks great, but let's add some more examples.")
	authorInstance.FinalizeChapter(chapterTitle)
}
