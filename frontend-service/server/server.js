const express = require('express');
const axios = require('axios');
const cors = require('cors');
const app = express();


app.use(cors());
app.use(express.json());

app.post('/ping-results', async (req, res) => {
  try {
    const response = await axios.post('http://backend:8080/api/ping/add', req.body);

    res.status(response.status).send(response.data);
  } catch (error) {
    console.error('Error redirecting request to Backend:', error.message);
    res.status(500).send('Error forwarding request to Backend');
  }
});


app.get('/data', async (req, res) => {
  try {
    const response = await axios.get('http://backend:8080/api/pings');
    res.send(response.data);
  } catch (error) {
    console.error('Error fetching data from Backend:', error.message);
    res.status(500).send('Error fetching data from Backend');
  }
});

app.listen(3001, '0.0.0.0', () => {
  console.log('Express server running on http://localhost:3001');
});