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
