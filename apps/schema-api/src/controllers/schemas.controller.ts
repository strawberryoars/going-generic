// import { Elysia } from 'elysia'
// import { existsSync } from 'fs';


// export const schemas = new Elysia()
//   .group('/schemas', (app) =>
//     app
//       .get ('/:schema', async (context) => {
//         console.log('GET /schemas', context);
//         // Create a new URL object from the request URL
//         // const url: URL = new URL(context.request.url, `http://localhost:3000`);

//         // Get the path part of the URL
//         // const path: string = url.pathname;
//         const path: string = context.path;
//         // Remove leading slash and replace remaining slashes with directory separators
//         const filePath: string = import.meta.dir + '/' +  path.replace(/^\/|\/$/g, '') + '.json';

//         console.log('found', existsSync('../../../schemas/blog.json'))
//         // Check if the file exists../
//         if (!existsSync(filePath)) {
//           return new Response("File not found", { status: 404 });
//         }

//         // Read the file
//         const file = Bun.file(filePath);

//         // Return the file content as JSON
//         return new Response(JSON.stringify(await file.json()), {
//           headers: { 'Content-Type': 'application/json' },
//         });
//     })
//   );

import { Elysia } from 'elysia';
import { existsSync } from 'fs';

export const schemas = new Elysia()
  .group('/schemas', (app) =>
    app
      .get('/:schema', async (context) => {
        const path: string = context.path;
        const filePath: string = `${import.meta.dir.split('/').slice(0, -2).join('/')}/${path.replace(/^\/|\/$/g, '')}.json`;
        console.log('filePath', filePath);  
        if (!existsSync(filePath)) {
          return new Response("File not found", { status: 404 });
        }

        const file = Bun.file(filePath);
        return new Response(JSON.stringify(await file.json()), {
          headers: { 'Content-Type': 'application/json' },
        });
      })
  );
