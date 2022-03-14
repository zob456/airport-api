DO
$$
    DECLARE
        renee_ID    uuid default '30f0bdb7-23b8-4c63-9f77-149b981e296d';
        Ohare_ID    uuid default '76da9c9e-2113-44b7-9d76-34b3451ee6e7';
        jfk_ID      uuid default '99287b16-0277-4177-93a9-aaeb01be2184';
        lax_ID      uuid default '47391fab-ad0a-4a12-b682-d483223353eb';
        jeanu_ID    uuid default 'f0a2724d-fc3e-4ab7-9634-22a3c16dfbaf';
        heathrow_ID uuid default '23e0edd1-4c16-413e-8610-6f6bd02b5e6d';
    BEGIN
        -- Create Users

        INSERT INTO "UserData"."user"("UserID", "UserName", "Password", "Active")
        VALUES (renee_ID, 'heyRenee', 'password', true);

        -- Create Airports

        INSERT INTO "AirportData"."airport"("AirportID", "Name", "City", "Country", "IATACode", "ICAOCode", "Latitude",
                                            "Longitude", "Altitude", "TimeZone")
        VALUES (Ohare_ID, 'Chicago O''Hare International Airport', 'Chicago', 'United States', 'ORD', 'KORD', 41.9786,
                -87.9048, 672, 'America/Chicago'),
               (jfk_ID, 'John F Kennedy International Airport', 'New York', 'United States', 'JFK', 'KJFK', 40.63980103,
                -73.77890015, 13, 'America/New_York'),
               (lax_ID, 'Los Angeles International Airport', 'Los Angeles', 'United States', 'LAX', 'KLAX', 33.94250107,
                -118.4079971, 125, 'America/Los_Angeles'),
               (jeanu_ID, 'Juneau International Airport', 'Juneau', 'United States', 'JNU', 'PAJN', 58.35499954223633,
                -134.5760040283203, 21, 'America/Anchorage'),
               (heathrow_ID, 'London Heathrow Airport', 'London', 'United Kingdom', 'LHR', 'EGLL', 51.4706, -0.461941,
                83, 'Europe/London');

    END
$$;
