import React, {
  Fragment,
  useState,
  useEffect,
  useCallback,
  useRef
} from "react";
import { Button, TextField } from "@material-ui/core";
import axios from "axios";
import "./App.css";

function App() {
  const [name, setName] = useState("");
  const [lat, setLat] = useState(0.0);
  const [lon, setLon] = useState(0.0);
  const [type, setType] = useState("");
  const [data, setData] = useState({});
  const [isSending, setIsSending] = useState(false);

  const isMounted = useRef(true);

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios("http://localhost:8080/locations");
      setData(result.data);
    };

    fetchData();
    return () => {
      isMounted.current = false;
    };
  }, []);

  function makePost() {
    axios("http://localhost:8080/locations", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        lat: lat,
        lon: lon,
        name: name,
        type: type
      })
    }).then(error => console.log(error));
  }

  const sendRequest = useCallback(async () => {
    // don't send again while we are sending
    if (isSending) return;
    // update state
    setIsSending(true);
    // send the actual request
    await makePost();
    // once the request is sent, update state again
    setIsSending(false);
    if (isMounted.current)
      // only update if we are still mounted
      setIsSending(false);
  }, [isSending]); // update the callback if the state changes

  return (
    <div>
      <div>
        <p>hi</p>
        <form>
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
              onClick={sendRequest}
              disabled={isSending}
            >
              Submit
            </Button>
          </div>
        </form>
      </div>
      <Fragment>
        <ul>
          {Array.from(data).map((item, num) =>
            item.hasOwnProperty("metadata") ? (
              <li key={num}>
                {item.metadata.hasOwnProperty("name") ? item.metadata.name : {}}
              </li>
            ) : (
              {}
            )
          )}
        </ul>
      </Fragment>
    </div>
  );
}

export default App;
