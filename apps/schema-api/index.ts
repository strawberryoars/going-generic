
import { existsSync } from 'fs';

const server = Bun.serve({
  port: 3000,
  async fetch(request) {
    // Create a new URL object from the request URL
    const url: URL = new URL(request.url, `http://localhost:${server.port}`);

    // Get the path part of the URL
    const path: string = url.pathname;

    // Remove leading slash and replace remaining slashes with directory separators
    const filePath: string = import.meta.dir + '/' +  path.replace(/^\/|\/$/g, '') + '.json';

    // Check if the file exists
    if (!existsSync(filePath)) {
      return new Response("File not found", { status: 404 });
    }

    // Read the file
    const file = Bun.file(filePath);

    // Return the file content as JSON
    return new Response(JSON.stringify(await file.json()), {
      headers: { 'Content-Type': 'application/json' },
    });
  },
});

console.log(`Listening on localhost:${server.port}`);
