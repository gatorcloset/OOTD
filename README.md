# OOTD App

OOTD (Outfit Of The Day) is an application created to help users put together outfits quickly, without the hassle or the mess of trying them on! Users are able to create a digital inventory of their closets and organize their clothes, simply by uploading images from the Internet. OOTD makes it easy for users to generate outfit ideas and create the optimal outfit for the onset of their day and even plan and save outfits for the future.

## Features
<img width="1440" alt="mockup2" src="https://github.com/gatorcloset/OOTD/assets/92867456/47c6d93b-3c7f-4a1c-9013-f07895776aeb">



https://github.com/gatorcloset/OOTD/assets/92867456/79063b85-6451-49d6-9c98-dd14fedaeb5e



https://github.com/gatorcloset/OOTD/assets/92867456/ddac6e61-ae7f-4465-a2fa-8e03ad4affa7





## Prerequisites
Before you get started, make sure you have the following tools installed:

- Go (version 1.20.3)
- Angular CLI (version 15.2.6)
- Node.js (version 18.15.0)

** These versions are the optimal version types to install. Depracated versions may not behave as expected.

## General Installation
To install and run the project locally, follow these steps:

1. Clone the repository locally using the following command 
```
git clone https://github.com/gatorcloset/OOTD.git
```

### Backend Installation
2. Ensure Go is properly installed by running "go --version" in Terminal
3. Navigating to the backend directory (by performing the cd command) like so:
```
cd backend
cd go
cd src
cd github.com
```
4. Install Go dependencies by running the following command:
```
go mod download
```

** If certain packages are not being installed properly, run "go get -u ____", where ___ is the package name

### Frontend Installation
5. Verify Node.js, Angular CLI and NPM are installed correctly by running "node -v", "ng -v", and "npm -v"
6. Run the following command, to install all necessary dependencies:
```
npm install
```
7. Run the build command to construct the project:
```
ng build
```
8. Lastly, run this command to launch the project on http://localhost:4200
```
ng serve
```
9. Happy outfit building :)

## Further help

- To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
- To get more help about Go use `go help` or go check out the [Go Installation Steps](https://go.dev/doc/install) page.

