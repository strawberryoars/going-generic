import { Elysia } from "elysia";
import { schemas } from "./controllers/schemas.controller";
import { faker } from "./controllers/faker.controller";

const app = new Elysia()
  .use(schemas)
  .use(faker)
  .listen(3000);
