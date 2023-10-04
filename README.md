# shop-organiser-api
Side project. This repo is meant to be only the back-end of the system.

Planned functionalities:
- fetch data from database
  - set db connection
  - query db
  - handle errors
- manipulate db data using internal models and services
- serve API for frontend client

## IntelliJ IDE configuration to work with project locally

#### Steps below are required to be able to run the app and its tests locally.

1. Install Go plugin
   - go to *Settings/Plugins*
   - search for *Go* plugin and install it
2. Enable Go modules integration so IntelliJ properly links local Go packages
   - go to *Settings/Languages & Frameworks/Go/Go modules*
   - select *Enable Go modules integration*
3. Download all dependencies
   - run *go mod tidy*