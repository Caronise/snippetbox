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
Specifically, the ui/html directory will contain HTML templates, and the
ui/static directory will contain static files (CSS and images).

