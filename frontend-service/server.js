const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');

const app = express();

app.use(bodyParser.json());

const PORT = 3001;

app.post('/api/ping-results', async (req, res) => {
    try {
        console.log("Received ping results: ", req.body);

        const response = await axios.post('http://localhost:8080/api/ping/add', req.body);
        console.log("Response from ping service: ", response.data);

        res.status(200).json({ message: "Ping results received" });
    } catch (error) {
        console.error("Error saving ping results: ", error);
        res.status(500).json({ message: "Error saving ping results" });
    }

});

app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});