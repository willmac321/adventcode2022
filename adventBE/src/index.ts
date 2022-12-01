import express from 'express';
const port = 4000;
const app = express();

app.get('/hey', (req, res) => res.send('hey'));

app.listen(port, () => {
    console.log(`⚡️[server]: Server is running at https://localhost:${port}`);
});

