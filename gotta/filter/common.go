package filter

import (
    "regexp"
    "strings"

    "gotta/model"
)

// MatchesRegexps filters out any Tree whose Task string does not match agains Regexps
func MatchesRegexps(regexps ...*regexp.Regexp) Predicate {
    return func(t model.Tree) bool {
        for _, re := range regexps {
            if re.FindStringIndex(t.Task().String()) != nil {
                return false
            }
        }

        return true
    }
}

// HasSubstrings filters out any Tree whose Task string does not contain substrings.
func HasSubstrings(substrings ...string) Predicate {
    return func(t model.Tree) bool {
        taskString := t.Task().String()
        for _, s := range substrings {
            if !strings.Contains(taskString, s) {
                return false
            }
        }

        return true
    }
}

// HasWordsWithPrefixSuffix filters out any Tree whose Task does not contain all
// the 'words' (prefixed with 'prefix' and suffixed with 'suffix')
func HasWordsWithPrefixSuffix(prefix string, suffix string, words ...string) Predicate {
    // TODO: Is there any way this compilation can fail?
    re, _ := regexp.Compile(
        `(?:^| )` +
        regexp.QuoteMeta(prefix) +
        `(\S+)` +
        regexp.QuoteMeta(suffix))

    // Create set of the words to match on
    requiredWords := map[string]bool{}
    for _, word := range words {
        requiredWords[word] = true
    }

    return func(t model.Tree) bool {
        // Create set of words in Task
        taskWords := map[string]bool{}
        for _, match := range re.FindAllStringSubmatch(t.Task().String(), -1) {
            taskWords[match[1]] = true
        }

        // requiredWords must be subset of taskWords
        for word, _ := range requiredWords {
            if !taskWords[word] {
                return false
            }
        }

        return true
    }
}

// HasWords filters out any Tree whose Task string does not contain all 'words'
func HasWords(words ...string) Predicate {
    return HasWordsWithPrefixSuffix("", "", words...)
}

// Leaves filters out any Tree that is not a leaf (no children)
func IsLeaf(t model.Tree) bool {
    _, ok := <- t.Children()
    return !ok
}
