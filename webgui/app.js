// Import required modules
const express = require("express");
const bodyParser = require("body-parser");
const biometricAuth = require("biometric-authentication-module");

// Initialize the app
const app = express();

// Use body-parser to parse JSON data
app.use(bodyParser.json());

// Use biometricAuth middleware to authenticate users
app.use(biometricAuth.authenticate);

// Define the '/' route
app.get("/", (req, res) => {
    res.send("Welcome to the bioinformatic security suite!");
});

// Define the '/secure-route' route
app.get("/secure-route", (req, res) => {
    res.send("You have successfully accessed a secure route!");
});

// Start the app
app.listen(3000, () => {
    console.log("Bioinformatic security suite app listening on port 3000!");
});
