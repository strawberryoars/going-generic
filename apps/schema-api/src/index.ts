import { Elysia } from "elysia";
import { schemas } from "./controllers/schemas.controller";

const app = new Elysia()
  .use(schemas)
  .listen(3000);
