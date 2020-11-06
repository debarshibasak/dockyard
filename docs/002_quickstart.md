# Quickstart

Dockyard is a static site generator that takes *Markdown* and uses predefined layouts to generate the static sites.
It is valuable for generating documentation, information sites, simple landing pages etc.

### Prerequisites

- Dockyard binary
- Understanding of Markdown language


### Instructions

Here we download the dockyard binary, then download some sample markdown files, run the generator on those files and serve them. 

#### Download dockyard

```$xslt
curl -o https://github.com
chmod +x dockyard
mv dockyard /usr/local/bin
dockyard -h
```

#### Download sample Markdown pages

```
curl -o data.tar.gz https://github.com
tar -xvzf data.tar.gz
```

#### Run it on .md files

```$xslt
dockyard --location docs --branding myowncompany --serve
```

To go, `http://localhost:10000`