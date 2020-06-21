import React, { useState } from "react";
import { Button, TextField } from "@material-ui/core";
import axios from "axios";
import "./App.css";

function App() {
  const [name, setName] = useState("");
  const [lat, setLat] = useState(0.0);
  const [lon, setLon] = useState(0.0);
  const [type, setType] = useState("");

  const handleSubmit = (lat, lon, name, type) => {
    const response = axios.post(
      "http://localhost:8080/locations",
      {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      JSON.stringify({
        lat: lat,
        lon: lon,
        name: name,
        type: type
      })
    );

    fetch("http://localhost:8080/locations", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        lat: lat,
        lon: lon,
        name: name,
        type: type
      })
    }).then(response => console.log(response.json));
    console.log("-------");
    alert(`Submitted`);
  };

  return (
    <form
      onSubmit={(lat, lon, name, type) => handleSubmit(lat, lon, name, type)}
    >
      <div>
        <TextField
          type="input"
          id="lat"
          label="Latitude"
          variant="outlined"
          onChange={e => setLat(e.target.value)}
        />
      </div>
      <div>
        <TextField
          type="input"
          id="lon"
          label="Longitude"
          variant="outlined"
          onChange={e => setLon(e.target.value)}
        />
      </div>
      <div>
        <TextField
          type="input"
          id="type"
          label="Location Type"
          variant="outlined"
          onChange={e => setType(e.target.value)}
        />
      </div>
      <div>
        <TextField
          type="input"
          id="name"
          label="Location Name"
          variant="outlined"
          onChange={e => setName(e.target.value)}
        />
      </div>
      <div>
        <Button
          type="submit"
          value="Submit"
          variant="contained"
          color="primary"
        >
          Submit
        </Button>
      </div>
    </form>
  );
}

export default App;
