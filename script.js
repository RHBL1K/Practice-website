document.addEventListener("DOMContentLoaded", function() {
    const cards = document.querySelectorAll(".card");
    const popup = document.getElementById("popup");
    const popupBody = document.getElementById("popup-body");

    popup.style.display = "none"; // Убедимся, что pop-up скрыт при загрузке

    cards.forEach((card, index) => {
        const button = document.createElement("button");
        button.innerText = "Подробнее";
        button.classList.add("popup-button");
        button.addEventListener("click", function(event) {
            event.stopPropagation(); // Предотвращает всплытие клика
            
            if (index === 4) { // Для card5 (индекс 4)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#2D2374; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>5 л, 10 л</td></tr>
                        <tr><td>Цвет</td><td>Молочный</td></tr>
                        <tr><td>Разбавление водой</td><td>Готовая к применению</td></tr>
                        <tr><td>Расход</td><td>120 мл / м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>12 часов для последующего слоя</td></tr>
                    </table>
                `;
            } else if (index === 0) { // Для card1 (индекс 0)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#C2185B; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>1,5 кг, 3,5 кг, 7 кг, 15 кг, 25 кг</td></tr>
                        <tr><td>Цвет</td><td>Белый, матовый</td></tr>
                        <tr><td>База для колерования</td><td>А, в светлые оттенки</td></tr>
                        <tr><td>Расход</td><td>120 гр/м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>4-6 часов для последующего слоя</td></tr>
                        <tr><td>Стойкость к мытью</td><td>Выдерживает влажную уборку</td></tr>
                    </table>
                `;
            } else if (index === 1) { // Для card2 (индекс 1)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#FFB300; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>1,2 кг, 3 кг, 12 кг, 19 кг</td></tr>
                        <tr><td>Цвет</td><td>Серый, полуматовый</td></tr>
                        <tr><td>База для колерования</td><td>С, яркие и темные цвета</td></tr>
                        <tr><td>Расход</td><td>120 гр/м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>4-6 часов для последующего слоя</td></tr>
                        <tr><td>Стойкость к мытью</td><td>Высокая, с применением моющих средств</td></tr>
                    </table>
                `;
            } else if (index === 2) { // Для card3 (индекс 2)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#0097A7; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>1,2 кг, 3 кг, 5 кг, 13 кг, 22 кг</td></tr>
                        <tr><td>Цвет</td><td>Белый, полуматовый</td></tr>
                        <tr><td>База для колерования</td><td>А, в светлые оттенки</td></tr>
                        <tr><td>Расход</td><td>120 гр/м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>4-6 часов для последующего слоя</td></tr>
                        <tr><td>Стойкость к мытью</td><td>Высокая, с применением моющих средств</td></tr>
                    </table>
                `;
            } else if (index === 3) { // Для card4 (индекс 3)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#FFB300; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>0,5 л, 1 л, 4 л</td></tr>
                        <tr><td>Цвет</td><td>Прозрачный, полу-глянцевый</td></tr>
                        <tr><td>Растворитель</td><td>Готов к употреблению</td></tr>
                        <tr><td>Расход</td><td>100 мл/м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>72 часа для практического применения</td></tr>
                    </table>
                `;
            } else if (index === 5) { // Для card6 (индекс 5)
                popupBody.innerHTML = `
                    <table border="1" style="width:100%; text-align:left; border-collapse: collapse;">
                        <tr style="background-color:#C2185B; color:white;"><th colspan="2">Область применения</th></tr>
                        <tr><td>Фасовка</td><td>3 кг, 5 кг, 10 кг, 15 кг, 20 кг</td></tr>
                        <tr><td>Цвет</td><td>Розовый</td></tr>
                        <tr><td>Разбавление водой</td><td>При необходимости</td></tr>
                        <tr><td>Расход</td><td>250 гр/м² при однослойном нанесении</td></tr>
                        <tr><td>Время высыхания</td><td>24 часа для последующего нанесения</td></tr>
                    </table>
                `;
            } else {
                popupBody.innerHTML = "<p>Здесь будет описание товара.</p>";
            }
            
            popup.style.display = "flex";
        });
        card.appendChild(button);
    });
});

function closePopup() {
    document.getElementById("popup").style.display = "none";
}