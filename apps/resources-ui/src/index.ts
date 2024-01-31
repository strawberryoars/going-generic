// const http = require('http');
const server = Bun.serve({
    port: 3000,
    fetch(request) {
        const filePath: string = './src/public/index.html';
        const file = Bun.file(filePath);
        return new Response(file);
    },
  });
  
  console.log(`Listening on localhost:${server.port}`);