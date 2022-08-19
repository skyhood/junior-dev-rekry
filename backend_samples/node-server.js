// To run this, simply run node node-server.js and visit localhost:3333/api?name=Petri

const http = require('http');
const url = require('url');

const PORT = 3333;

const api = (res, urlObj) => {
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.write(JSON.stringify({
    'hello': urlObj.query?.name || "World",
  }));
  res.end();
}

const requestListener = (req, res) => {
  const urlObj = url.parse(req.url, true);
  
  switch (urlObj.pathname) {
    case '/api':
      return api(res, urlObj);
    default:
      res.writeHead(404);
      res.write(JSON.stringify({"oops": "This is not the page you are looking for"}));
      res.end();
  }
}

const server = http.createServer(requestListener);

console.log(`Listening on localhost:${PORT}`);
server.listen(PORT)