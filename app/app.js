
import Express from "express";
import {env} from 'process';   
import dotenv from 'dotenv';
import { dirname } from 'path';
import { fileURLToPath } from 'url';
import ToDoService from "./service.js";
import ToDoController from "./controller.js";

const scopes = env.SCOPES?.split(";") || [];

if (scopes.some((scope) => scope === 'dev')) {
    dotenv.config();
}

const app = Express();
const { APP_PORT = 3000 } = env;

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

app.use(Express.static(__dirname + "/public"));

const { API_HOST = 'localhost', API_PORT = 8080 } = env;

const toDoService = new ToDoService(API_HOST, API_PORT);
const toDoController = new ToDoController(toDoService);

app.get('/to-do', toDoController.getAllToDos.bind(toDoController));
app.post('/to-do', toDoController.createToDo.bind(toDoController));
app.get('/to-do/check/:id', toDoController.checkToDoById.bind(toDoController));
app.get('/to-do/delete/:id', toDoController.deleteToDoById.bind(toDoController));

app.listen(APP_PORT, () => {
    console.log(`server started at http://localhost:${APP_PORT}`);
});

