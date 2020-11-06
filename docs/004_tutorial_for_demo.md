# Github pages tutorials

### Setup your Github Page

 - For a given project, go to settings.
 - Go to the section, *GitHub Pages*
 - Choose your branch (default is `master`)
 - Choose the folder (default is `/`)
 
Note the location of published site.
 
### Setting up Markdown files for Dockyard

Create a directory structure as mentioned in section [Detailed Overview](#detailed_overview) with the markdown files inside the project directory.

Run -

```$xslt
dockyard --location <document-location> --branding myowncompany
```

Push the changes to remote Github repository.

Go to `https://<user-or-orgname>.github.io/<project>/<documentation-directory>`

The newly deployed pages should be visible here.