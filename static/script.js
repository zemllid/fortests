document.addEventListener("DOMContentLoaded", async () => {await loadFrogs();})
async function loadFrogs()
{
    try {
        const response = await fetch("/frogs"); // Запрашиваем список жаб с сервера
        const frogs = await response.json();
        const tableBody = document.querySelector("#frogsTableBody");

        if (frogs.length === 0) {
            tableBody.innerHTML = "<tr><td colspan='4'>Нет данных о жабах 🐸</td></tr>";
            return;
        }

        frogs.forEach(frog => {
            const row = tableBody.insertRow();
            row.insertCell(0).innerText = frog.name;
            row.insertCell(1).innerText = frog.species;
            row.insertCell(2).innerText = frog.habitat;
            row.insertCell(3).innerText = frog.age;
        });
    } catch (error) {
        console.error("Ошибка загрузки данных о жабах:", error);
    }
});

// Функция для добавления новой жабы (пример)
async function addFrog(name, species, habitat, age) {
    const frog = { name, species, habitat, age };
    try {
        const response = await fetch("/frogs", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(frog),
        });

        if (!response.ok) throw new Error("Ошибка добавления жабы");

        console.log("Жаба успешно добавлена:", frog);
        await loadFrogs(); // Обновляем страницу, чтобы отобразить новую жабу
    } catch (error) {
        console.error("Ошибка:", error);
    }
}
