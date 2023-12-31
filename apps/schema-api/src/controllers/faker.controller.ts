import { Elysia } from 'elysia';
import { existsSync } from 'fs';
// import jsf from 'json-schema-faker';
import { JSONSchemaFaker } from "json-schema-faker";

const getSchema = async (schemaName: string) => {
  const filePath: string = `${import.meta.dir.split('/').slice(0, -2).join('/')}/schemas/${schemaName}.json`;
  console.log('filePath', filePath);  
  if (!existsSync(filePath)) {
    return  null;
  }

  const file = Bun.file(filePath);
  return await file.json();
};

export const faker = new Elysia()
  .group('/faker', (app) =>
    app.get('/:schema', async (context) => {
      const schemaName: string = context.params.schema;
      // Assuming you have a function getSchema(schemaName) that returns the schema
      const schema = await getSchema(schemaName);

      if (!schema) {
        return new Response("File not found", { status: 404 });
      }

      try {
        // const result = await jsf.resolve(schema);
        const result = await JSONSchemaFaker.generate(schema);
        return new Response(JSON.stringify(result), {
          headers: { 'Content-Type': 'application/json' },
        });
      } catch (error) {
        console.error(`Error generating data for ${schemaName}:`, error);
        return new Response("Error generating data", { status: 500 });
      }
      })
  );
