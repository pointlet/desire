CREATE TABLE IF NOT EXISTS paths (
  id SERIAL PRIMARY KEY, 
  title VARCHAR(60) NOT NULL,
  description TEXT, 
  image_lookup_id SERIAL UNIQUE NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create a index for a combination of latitude and longitude
CREATE INDEX IF NOT EXISTS idx_lat_lon ON paths (latitude, longitude);


