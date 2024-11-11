const toDoInput = document.getElementById('to-do-input');
const toDoListTable = document.getElementById('to-do-list');
const toDoSubmit = document.getElementById('add-to-do');

function formatDate(date) {
    if (!date) {
        return '';
    }

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

function toDoRow(toDo) {
    const toDoRow = `
        <tr>
            <td>${formatDate(toDo.create_at ?? null)}</td>
            <td class="to-do-item-content">
                <span class="to-do-item-text" data-id="${toDo.id}">
                    ${toDo.title ?? ''}
                </span>
                <div class="to-do-item-btn-container">
                    <a href="/to-do/check/${toDo.id}">
                        <button class="to-do-check-btn" data-id="${toDo.id}">
                            <i class="fa-solid fa-check"></i>
                        </button>
                    </a>
                    <a href="/to-do/delete/${toDo.id}">
                        <button class="to-do-delete-btn" data-id="${toDo.id}">
                            <i class="fa-solid fa-trash"></i>
                        </button>
                    </a>
                </div>
            </td>
        </tr>
    `;

    return toDoRow;
}

toDoSubmit.addEventListener('click', async () => {
    try {

        const toDoText = toDoInput.value;

        if (!toDoText) {
            return;
        }

        const toDoRes = await fetch(`/to-do`, {
            method: 'POST',
            body: JSON.stringify({ title: toDoText }),
            headers: {
                'Content-Type': 'application/json'
            }
        });

        const toDo = await toDoRes.json();

        toDoListTable.innerHTML += toDoRow(toDo);
        toDoInput.value = '';
        window.location.reload();
    } catch (error) {
        console.error(error);
    }
});

async function fillToDoList() {
    const toDoList = [];

    try {
        const toDoRes = await fetch(`/to-do`);
        const toDoListData = await toDoRes.json();
        toDoList.push(...toDoListData);
    } catch (error) {
        console.error(error);
    }

    if (!toDoList.length) {
        toDoListTable.innerHTML = `
            <tr>
                <td colspan="2">No To-Do Items</td>
            </tr>
        `;
    }

    toDoList.forEach((toDo) => {
        toDoListTable.innerHTML += toDoRow(toDo);
    });
}

fillToDoList();
