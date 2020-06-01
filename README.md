# Open Crisis Line 2

## Overview
This an attempt of a rewrite of the Open Crisis Line. 
No code is shared between the projects.  A key point since they are licensed differently

The layout of this project follows:
- api holds the open api definition of the REST API
- app holds the main executables
- cicd is the build and test pipeline scripts
- pkg has all the guts of code down in it
- tests has mostly just l2 type tests
- third_party holds external tools and things like swagger ui
- web is the UI for the app


## Server Env Vars
- TWILLO_TOKEN  The auth token from Twillo
- TWILLO_SID The SID from Twillo
- DUTY_NUMBER  The number to SMS for the on duty call taker




# Design
When a user needs support, the enter their basic info:
- Phone number (required)
- Name optional playa or default
- Description of what is going on today

When the user submits a request:
 - A random code is generated and displayed to the user (4 digit friendly number)
 - an entry is saved in the DB (probably my sql)
 - a text is sent to the user's cell phone asking them to enter the code
 - if the same code is entered, SMS is sent to duty folks
 
 
 
and a verification is sent to their cell phone.