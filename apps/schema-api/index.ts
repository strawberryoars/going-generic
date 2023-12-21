const file = Bun.file('/schemas/plants.json');
console.log(typeof await Bun.file('/schemas/plants.json').text());
// const contents = await file.json();
// console.log(contents);

const server = Bun.serve({
 port: 3000,
 async fetch(request) {
  // return new Response("Welcome to Bun!");
  return new Response(JSON.stringify(await file.json()), {
    headers: { 'Content-Type': 'application/json' },
  });
 },
});

console.log(`Listening on localhost:${server.port}`);
