const file = Bun.file(import.meta.dir + '/schemas/plants.json');

const server = Bun.serve({
  port: 3000,
  async fetch(request) {
    return new Response(JSON.stringify(await file.json()), {
      headers: { 'Content-Type': 'application/json' },
    });
  },
});

console.log(`Listening on localhost:${server.port}`);
