package fake

import (
	"strings"
)

// Company generates company name
func Company() string {
	return lookup(lang, "company_names", true)
}

// JobTitle generates job title
func JobTitle() string {
	job := lookup(lang, "job_titles", true)
	return strings.Replace(job, "#{N}", jobTitleSuffix(), 1)
}

func jobTitleSuffix() string {
	return lookup(lang, "job_titles_suffixes", false)
}

// Industry generates industry name
func Industry() string {
	return lookup(lang, "industries", true)
}
