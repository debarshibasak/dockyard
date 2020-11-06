# Detailed Overview

### Building side menu

You can build hierarchical directory structure that automatically detects the files and generates the html files.

***For instance***

You can have a directory called  - `doc`. To add a menu item create a file `<sequence>_<menu_item>.md` under the directory `doc`. 
The menu items are ordered by sequence number and then lexicologically.
A `<sequence>` could be something like `001` or anything that could be used to sort the menu item.

The file could look something as follows - 

```
doc
doc/001_first_item.md
doc/002_second_item.md
```

### Generating static pages


### Options


```$xslt

  -branding string
        branding of documentation (default "Acme")
  -location string
        location of documentation
  -serve
        generate and serve the documentation
  -theme string
        only default supported at this point (default "default")
  -watch
        watch for document change and regenerate when changed

```
