

export default class ToDoController {
    constructor(service) {
        this.service = service;
    }

    async getAllToDos(req, res) {
        try {
            const toDos = await this.service.getAllToDos();
            res.status(200)
                .json(toDos);
        } catch (error) {
            res.status(500)
                .json({ message: error.message });
        }
    }

    async createToDo(req, res) {
        try {
            const { title } = req.body;
            const created = await this.service.createToDo(title);
            res.status(201)
                .json(created);
        } catch (error) {
            res.status(500)
                .json({ message: error.message });
        }  
    }

    async checkToDoById(req, res) {
        try {
            const { id } = req.query;
            await this.service.checkToDo(id);
            res.status(200)
                .json({ message: 'Checked', id });
        } catch (error) {
            res.status(500)
                .json({ message: error.message });
        }  
    }

    async deleteToDoById(req, res) {
        try {
            const { id } = req.query;
            await this.service.deleteToDoById(id);
            res.status(200)
                .json({ message: 'Deleted', id });
        } catch (error) {
            res.status(500)
                .json({ message: error.message });
        }
    }
}