# snippetbox

## cmd

The *cmd* directory contains the application-specific code for the executable
applications in the project. For now it is only the web application.

## internal

The *internal* directory contains the ancillary non-application-specific code
used in the project. Such as validation helpers and SQL database models. In Go,
any packages under this directory can only be imported by code inside of the
parent of the internal directory.

## ui

The *ui* directory contains the user-interface assets used by the web app.

### ui/html
The *ui/html* directory contains HTML templates

- *base.tmpl*: is the base template that is used in most pages.
- *ui/html/pages*: contains the templates for each individual page.
- *ui/html/partials*: contains the re-usuable template partials.

### ui/static
The *ui/static* directory contains static files (CSS and images).

Note: This project uses the .tmpl extension for HTML templates, to make it clear
that the file contains a Go template. Alternatively .tmpl.html can be used.
