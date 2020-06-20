CREATE TABLE public.locations (
    id uuid PRIMARY KEY NOT NULL,
    latitude FLOAT,
    longitude FLOAT,
    location_name VARCHAR NULL,
    location_type VARCHAR NULL
    );
