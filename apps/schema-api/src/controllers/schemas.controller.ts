import { Elysia } from 'elysia';
import { existsSync, readdirSync } from 'fs';
import { join } from 'path';

// schemas endpoint
export const schemas = new Elysia()
  .group('/schemas', (app) => 
    app
      // serves all schemas
      .get('/', async (context) => {
        const dirPath: string = `${import.meta.dir.split('/').slice(0, -2).join('/')}/schemas`;
        console.log('dirPath', dirPath); 

        if (!existsSync(dirPath)) {
          return new Response("Directory not found", { status: 404 });
        }

        const files = readdirSync(dirPath);
        const schemas = await Promise.all(files.map(async (file) => {
          const filePath = join(dirPath, file);
          const fileData = await Bun.file(filePath).json();
          return fileData;
        }));

        return new Response(JSON.stringify(schemas), {
          headers: { 'Content-Type': 'application/json' },
        });
      })
      // serves requested schema
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
