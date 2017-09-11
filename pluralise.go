package pluralise

import (
	"regexp"
	"strconv"
	"strings"
)

// Pluralise pluralises a word
func Pluralise(word string) string {
	return replaceWord(irregularSingles, irregularPlurals, pluralRules, word)
}

// WithCount pluralises or singularises a word depending on the count supplied
func WithCount(word string, count int) string {
	if count == 1 {
		return Singularise(word)
	}

	return Pluralise(word)
}

// WithCountInclusive pluralises or singularises a word depending on the count supplied,
// prepending the count to the word
func WithCountInclusive(word string, count int) string {
	return strconv.Itoa(count) + " " + word
}

// Singularise singularises a word
func Singularise(word string) string {
	return replaceWord(irregularPlurals, irregularSingles, singularRules, word)
}

// AddPluralRule adds a pluralisation rule to the collection
func AddPluralRule(pattern string, replacement string) {
	pluralRules = append(pluralRules, rule{
		Rule:        regexp.MustCompile("(?i)" + pattern),
		Replacement: replacement,
	})
}

// AddSingularRule adds a singularisation rule to the collection
func AddSingularRule(pattern string, replacement string) {
	singularRules = append(singularRules, rule{
		Rule:        regexp.MustCompile("(?i)" + pattern),
		Replacement: replacement,
	})
}

// AddUncountableRule adds an uncountable word rule
func AddUncountableRule(word string) {
	AddIrregularRule(word, word)
}

// AddIrregularRule adds and irregular word definition
func AddIrregularRule(single string, plural string) {
	single = strings.ToLower(single)
	plural = strings.ToLower(plural)

	irregularSingles[single] = plural
	irregularPlurals[plural] = single
}

var (
	pluralRules      []rule
	singularRules    []rule
	irregularSingles map[string]string
	irregularPlurals map[string]string
)

func init() {
	irregularSingles = make(map[string]string)
	irregularPlurals = make(map[string]string)

	// Pronouns.
	AddIrregularRule("I", "we")
	AddIrregularRule("me", "us")
	AddIrregularRule("he", "they")
	AddIrregularRule("she", "they")
	AddIrregularRule("myself", "ourselves")
	AddIrregularRule("yourself", "yourselves")
	AddIrregularRule("itself", "themselves")
	AddIrregularRule("herself", "themselves")
	AddIrregularRule("himself", "themselves")
	AddIrregularRule("themself", "themselves")
	AddIrregularRule("is", "are")
	AddIrregularRule("was", "were")
	AddIrregularRule("has", "have")
	AddIrregularRule("this", "these")
	AddIrregularRule("that", "those")
	// Words ending in with a consonant and `o`.
	AddIrregularRule("echo", "echoes")
	AddIrregularRule("dingo", "dingoes")
	AddIrregularRule("volcano", "volcanoes")
	AddIrregularRule("tornado", "tornadoes")
	AddIrregularRule("torpedo", "torpedoes")
	// Ends with `us`.
	AddIrregularRule("genus", "genera")
	AddIrregularRule("viscus", "viscera")
	// Ends with `ma`
	AddIrregularRule("stigma", "stigmata")
	AddIrregularRule("stoma", "stomata")
	AddIrregularRule("dogma", "dogmata")
	AddIrregularRule("lemma", "lemmata")
	AddIrregularRule("schema", "schemata")
	AddIrregularRule("anathema", "anathemata")
	// Other irregular rules.
	AddIrregularRule("ox", "oxen")
	AddIrregularRule("axe", "axes")
	AddIrregularRule("die", "dice")
	AddIrregularRule("yes", "yeses")
	AddIrregularRule("foot", "feet")
	AddIrregularRule("eave", "eaves")
	AddIrregularRule("goose", "geese")
	AddIrregularRule("tooth", "teeth")
	AddIrregularRule("quiz", "quizzes")
	AddIrregularRule("human", "humans")
	AddIrregularRule("proof", "proofs")
	AddIrregularRule("carve", "carves")
	AddIrregularRule("valve", "valves")
	AddIrregularRule("looey", "looies")
	AddIrregularRule("thief", "thieves")
	AddIrregularRule("groove", "grooves")
	AddIrregularRule("pickaxe", "pickaxes")
	AddIrregularRule("whiskey", "whiskies")

	AddPluralRule("^thou$", "you")
	AddPluralRule("m[ae]n$", "men")
	AddPluralRule("eaux$", "${0}")
	AddPluralRule("(child)(?:ren)?$", "${1}ren")
	AddPluralRule("(pe)(?:rson|ople)$", "${1}ople")
	AddPluralRule("(m|l)(?:ice|ouse)$", "${1}ice")
	AddPluralRule("(x|ch|ss|sh|zz)$", "${1}es")
	AddPluralRule("([^ch][ieo][ln])ey$", "${1}ies")
	AddPluralRule("([^aeiouy]|qu)y$", "${1}ies")
	AddPluralRule("(?:(kni|wi|li)fe|(ar|l|ea|eo|oa|hoo)f)$", "${1}${2}ves")
	AddPluralRule("sis$", "ses")
	AddPluralRule("(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)(?:a|on)$", "${1}a")
	AddPluralRule("(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|automat|quor)(?:a|um)$", "${1}a")
	AddPluralRule("(her|at|gr)o$", "${1}oes")
	AddPluralRule("(seraph|cherub)(?:im)?$", "${1}im")
	AddPluralRule("(alumn|alg|vertebr)(?:a|ae)$", "${1}ae")
	AddPluralRule("(alumn|syllab|octop|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$", "${1}i")
	AddPluralRule("([^l]ias|[aeiou]las|[emjzr]as|[iu]am)$", "${1}")
	AddPluralRule("(e[mn]u)s?$", "${1}s")
	AddPluralRule("(alias|[^aou]us|tlas|gas|ris)$", "${1}es")
	AddPluralRule("(ax|test)is$", "${1}es")
	AddPluralRule("([^aeiou]ese)$", "${1}")
	AddPluralRule("[^\u0000-\u007F]$", "${0}")
	AddPluralRule("s?$", "s")

	AddSingularRule("men$", "man")
	AddSingularRule("(eau)x?$", "${1}")
	AddSingularRule("(child)ren$", "${1}")
	AddSingularRule("(pe)(rson|ople)$", "${1}rson")
	AddSingularRule("(matr|append)ices$", "${1}ix")
	AddSingularRule("(cod|mur|sil|vert|ind)ices$", "${1}ex")
	AddSingularRule("(alumn|alg|vertebr)ae$", "${1}a")
	AddSingularRule("(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)a$", "${1}on")
	AddSingularRule("(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|quor)a$", "${1}um")
	AddSingularRule("(alumn|syllab|octop|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$", "${1}us")
	AddSingularRule("(test)(?:is|es)$", "${1}is")
	AddSingularRule("(movie|twelve|abuse|e[mn]u)s$", "${1}")
	AddSingularRule("(analy|ba|diagno|parenthe|progno|synop|the|empha|cri)(?:sis|ses)$", "${1}sis")
	AddSingularRule("(x|ch|ss|sh|zz|tto|go|cho|alias|[^aou]us|tlas|gas|(?:her|at|gr)o|ris)(?:es)?$", "${1}")
	AddSingularRule("(seraph|cherub)im$", "${1}")
	AddSingularRule("(m|l)ice$", "${1}ouse")
	AddSingularRule("\b(mon|smil)ies$", "${1}ey")
	AddSingularRule("\b([pl]|zomb|(?:neck|cross)?t|coll|faer|food|gen|goon|group|lass|talk|goal|cut)ies$", "${1}ie")
	AddSingularRule("ies$", "y")
	AddSingularRule("(ar|(?:wo|[ae])l|[eo][ao])ves$", "${1}f")
	AddSingularRule("(wi|kni|(?:after|half|high|low|mid|non|night|[^\\w]|^)li)ves$", "${1}fe")
	AddSingularRule("(ss)$", "${1}")
	AddSingularRule("s$", "")

	AddUncountableRule("adulthood")
	AddUncountableRule("advice")
	AddUncountableRule("agenda")
	AddUncountableRule("aid")
	AddUncountableRule("alcohol")
	AddUncountableRule("ammo")
	AddUncountableRule("anime")
	AddUncountableRule("athletics")
	AddUncountableRule("audio")
	AddUncountableRule("bison")
	AddUncountableRule("blood")
	AddUncountableRule("bream")
	AddUncountableRule("buffalo")
	AddUncountableRule("butter")
	AddUncountableRule("carp")
	AddUncountableRule("cash")
	AddUncountableRule("chassis")
	AddUncountableRule("chess")
	AddUncountableRule("clothing")
	AddUncountableRule("cod")
	AddUncountableRule("commerce")
	AddUncountableRule("cooperation")
	AddUncountableRule("corps")
	AddUncountableRule("debris")
	AddUncountableRule("diabetes")
	AddUncountableRule("digestion")
	AddUncountableRule("elk")
	AddUncountableRule("energy")
	AddUncountableRule("equipment")
	AddUncountableRule("excretion")
	AddUncountableRule("expertise")
	AddUncountableRule("flounder")
	AddUncountableRule("fun")
	AddUncountableRule("gallows")
	AddUncountableRule("garbage")
	AddUncountableRule("graffiti")
	AddUncountableRule("headquarters")
	AddUncountableRule("health")
	AddUncountableRule("herpes")
	AddUncountableRule("highjinks")
	AddUncountableRule("homework")
	AddUncountableRule("housework")
	AddUncountableRule("information")
	AddUncountableRule("jeans")
	AddUncountableRule("justice")
	AddUncountableRule("kudos")
	AddUncountableRule("labour")
	AddUncountableRule("literature")
	AddUncountableRule("machinery")
	AddUncountableRule("mackerel")
	AddUncountableRule("mail")
	AddUncountableRule("media")
	AddUncountableRule("mews")
	AddUncountableRule("moose")
	AddUncountableRule("music")
	AddUncountableRule("manga")
	AddUncountableRule("news")
	AddUncountableRule("pike")
	AddUncountableRule("plankton")
	AddUncountableRule("pliers")
	AddUncountableRule("pollution")
	AddUncountableRule("premises")
	AddUncountableRule("rain")
	AddUncountableRule("research")
	AddUncountableRule("rice")
	AddUncountableRule("salmon")
	AddUncountableRule("scissors")
	AddUncountableRule("series")
	AddUncountableRule("sewage")
	AddUncountableRule("shambles")
	AddUncountableRule("shrimp")
	AddUncountableRule("species")
	AddUncountableRule("staff")
	AddUncountableRule("swine")
	AddUncountableRule("tennis")
	AddUncountableRule("they")
	AddUncountableRule("them")
	AddUncountableRule("traffic")
	AddUncountableRule("transporation")
	AddUncountableRule("trout")
	AddUncountableRule("tuna")
	AddUncountableRule("wealth")
	AddUncountableRule("welfare")
	AddUncountableRule("whiting")
	AddUncountableRule("wildebeest")
	AddUncountableRule("wildlife")
	AddUncountableRule("you")
}

type rule struct {
	Rule        *regexp.Regexp
	Replacement string
}

func restoreCase(word, token string) string {
	if word == token {
		return token
	}

	if word == strings.ToUpper(word) {
		return strings.ToUpper(token)
	}

	firstChar := string([]rune(word)[0])
	if firstChar == strings.ToUpper(firstChar) {
		return strings.Title(token)
	}

	return strings.ToLower(token)
}

func replace(word string, r rule) string {
	return r.Rule.ReplaceAllString(word, r.Replacement)
}

func sanitiseWord(token, word string, rules []rule) string {
	if len(token) == 0 {
		return word
	}

	for _, r := range rules {
		if r.Rule.MatchString(word) {
			return replace(word, r)
		}
	}

	return word
}

func replaceWord(replaceMap, keepMap map[string]string, rules []rule, word string) string {
	token := strings.ToLower(word)

	if _, ok := keepMap[token]; ok {
		return restoreCase(word, token)
	}

	if rpToken, ok := replaceMap[token]; ok {
		return restoreCase(word, rpToken)
	}

	return sanitiseWord(token, word, rules)
}
