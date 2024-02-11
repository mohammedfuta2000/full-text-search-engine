# Full Text Search Engine

### Run the Application
prerequistie: download the following file into the current dir
https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
it is the dataset we will be searching through

```bash
$ go run . -q "any-string"

```
this will return all the docs with "any-string" present in them

