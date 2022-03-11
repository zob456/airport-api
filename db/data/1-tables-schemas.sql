-- Creating schemas

CREATE SCHEMA "AirportData";
CREATE SCHEMA "UserData";

-- Creating tables

CREATE TABLE IF NOT EXISTS "UserData"."user"
(
    "UserID"   uuid PRIMARY KEY NOT NULL,
    "UserName" text             NOT NULL,
    "Password" text             not null,
    "Active"   boolean          NOT NULL,
    "JWT"      text
);

CREATE TABLE IF NOT EXISTS "AirportData"."airport"
(
    "AirportID" uuid PRIMARY KEY NOT NULL,
    "Name"      text             NOT NULL,
    "City"      text             NOT NULL,
    "Country"   text             NOT NULL,
    "IATACode"  text             NOT NULL,
    "ICAOCode"  text             NOT NULL,
    "Latitude"  float            NOT NULL,
    "Longitude" float            NOT NULL,
    "Altitude"  float            NOT NULL,
    "TimeZone"  text             NOT NULL
);
