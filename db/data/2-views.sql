-- Creating a view for user data to be used by auth for performance
CREATE OR REPLACE VIEW "UserData"."vw_User" as
SELECT "UserID",
       "UserName",
       "Password",
       "Active",
       "JWT"
FROM "UserData"."user";

CREATE OR REPLACE VIEW "AirportData"."vw_Airport" as
SELECT "AirportID",
    "Name",
    "City",
    "Country",
    "IATACode",
    "ICAOCode",
    "Latitude",
    "Longitude",
    "Altitude",
    "TimeZone"
FROM "AirportData"."airport";