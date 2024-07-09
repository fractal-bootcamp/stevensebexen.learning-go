import http from 'http';
import { execFile } from 'child_process';

const PORT = 80;

const server = http.createServer((req,res) => {
  if (req.url === '/') {
    if (req.method !== 'GET') {
      // Bad method
      res.writeHead(405).end();
      return;
    }

    // Run child process
    const process = execFile('./child', {encoding: 'utf-8'}, (error, stdout, stderr) => {
      if (error) {
        console.error('Child process error:', stderr);
        res.writeHead(500).end();
        return;
      }

      res.setHeader('Content-Type', 'application/json').writeHead(200).end(stdout);
    });

  } else { 
    // Bad route
    res.writeHead(404).end();
    return;
  }
});

server.listen(PORT, () => console.log(`Server listening on port ${PORT}!`));