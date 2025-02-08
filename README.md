# shop-organiser-api
Side project. This repo is meant to be only the back-end of the system.

Planned functionalities:
- fetch data from database
  - set db connection
  - query db
  - handle errors
- manipulate db data using internal models and services
- serve REST API for frontend client

## IntelliJ IDE configuration to work with project locally

#### Steps below are required to be able to run the app and its tests locally.

1. Install Go plugin
   - go to *Settings/Plugins*
   - search for *Go* plugin and install it
2. Enable Go modules integration so IntelliJ properly links local Go packages
   - go to *Settings/Languages & Frameworks/Go/Go modules*
   - select *Enable Go modules integration*
3. Download all dependencies
   - run *go mod tidy* in *./server* directory

#### Running app locally

Call *go run* in *./app* directory. Go to *localhost:5000* in the browser. Check out available endpoints in *./server* directory and explore them in the browser.

#### Running unit tests locally

Call *go test ./server* or use Intellij interface.

## Possible issues with local development

Combination of IntelliJ version 2022.2 (Build #IU-222.3345.118, built on July 26, 2022) and golang version 1.23.6 resulted in issues with debugging: code does not stop on debugging breakpoints but keeps processing forever. Working fix is to use golang 1.20 with this version of IntelliJ.