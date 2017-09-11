# Pluralise for Golang

> Pluralise and singularise any word.

A port of [blakeembrey/pluralize](https://github.com/blakeembrey/pluralize) to golang

## Installation

```
go get -u github.com/UnwrittenFun/pluralise
```

## Usage

```golang
import "github.com/UnwrittenFun/pluralise"

pluralise.Pluralise("test") //=> "tests"
pluralise.WithCount("test", 1) //=> "test"
pluralise.WithCount("test", 5) //=> "tests"
pluralise.WithCountInclusive("test", 1) //=> "1 test"
pluralise.WithCountInclusive("test", 5) //=> "5 tests"
pluralise.WithCountInclusive("蘋果", 2) //=> "2 蘋果"

pluralise.Pluralise("regex") //=> "regexes"
pluralise.AddPluralRule("gex$", "gexii")
pluralise.Pluralise("regex") //=> "regexii"

pluralise.Singularise("singles") //=> "single"
pluralise.AddSingularRule("singles$", "singular")
pluralise.Singularise("singles") //=> "singular"

pluralise.Pluralise("irregular") //=> "irregulars"
pluralise.AddIrregularRule("irregular", "regular")
pluralise.Pluralise("irregular") //=> "regular"

pluralise.Pluralise("paper") //=> "papers"
pluralise.AddUncountableRule("paper")
pluralise.Pluralise("paper") //=> "paper"
```

## License

MIT
