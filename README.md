`splitby -h`

# Overview
Reads content from `stdin` and split the content by given _spliter_ ( default `\n`)

# why?

I've been working with several files and I need to split the content
by a given separator so I can extract content from these files.
example:

```
cat queries.sql | splitby -spliter SELECT
```

So I can get all my queries ( I know there are single SELECT in each query)

Or

```
cat *.rb | splitby -spliter "def "
```

To find all the methods in a given folder of ruby files

# Usage

```
splitby -h
```
