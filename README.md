# snippetbox

Snippetbox is an application that allows users to create, read and delete Snippets.
The Directory layout is as follows:

## cmd/

The *cmd/* directory contains the application-specific code for the executable
applications in the project. 

### cmd/web/

Snippetbox has a web application, this is where most of the code to operate the
web app is found.

## internal/

The *internal/* directory contains the ancillary non-application-specific code
used in the project. Such as validation helpers and SQL database models. In Go,
any packages under this directory can only be imported by code inside of the
parent of the internal directory.

### internal/models

The *internal/models* directory contains the models for the MySQL database.

## ui/

The *ui/* directory contains the user-interface assets used by the web app.

### ui/html/
The *ui/html/* directory contains the HTML templates used by the web app.

- *base.tmpl* is the base template that is used for all web pages.
- *pages/* directory contains the templates for each individual page.
- *partials/* directory contains the re-usuable template partials.

### ui/static/
The *ui/static/* directory contains static files (CSS and images).

Note: This project uses the .tmpl extension for HTML templates, to make it clear
that the file contains a Go template. Alternatively .tmpl.html can be used.
