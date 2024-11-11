
export default class ToDoService {
    constructor(apiHost, apiPort) {
        this.apiHost = apiHost;
        this.apiPort = apiPort;
    }

    async getAllToDos() {
        const response = await fetch(`${this.apiHost}:${this.apiPort}/to-do`);
        return await response.json();
    }

    async createToDo(title) {
        const response = await fetch(`${this.apiHost}:${this.apiPort}/to-do`, {
            method: 'POST',
            body: JSON.stringify({ title }),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return await response.json();
    }

    async checkToDo(id) {
        const response = await fetch(`${this.apiHost}:${this.apiPort}/to-do/${id}`, {
            method: 'PATCH',
            body: JSON.stringify({ id, completed: true }),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return await response.json();
    }

    async deleteToDoById(id) {
        const response = await fetch(`${this.apiHost}:${this.apiPort}/to-do/${id}`, {
            method: 'DELETE'
        });
        return await response.json();
    }
}