To-Do List Desktop App (Wails + Go + React)
Описание

Десктопное приложение для управления задачами (To-Do List).
Позволяет:

*Создавать задачи с указанием даты и времени выполнения, а также приоритета (низкий / средний / высокий)

*Удалять задачи с подтверждением

*Отмечать задачи как выполненные

*Фильтровать задачи по статусу (Все / Активные / Выполненные)

*Сортировать задачи по дате и приоритету

*Сохранять состояние задач между запусками приложения (PostgreSQL)

*Реализованные функции по чеклисту
| Функция                                                      | Статус | Примечание                             |
| ------------------------------------------------------------ | ------ | -------------------------------------  |
| Интерфейс: текстовое поле, кнопка, список задач, CSS, значки | ✅      | Основная часть выполнена / tailwind   |
| Адаптивная верстка                                           | ⚪      | Не реализована                        |
| Темная/светлая тема                                          | ⚪      | Не реализована                        |
| Добавление задач с проверкой на пустой ввод                  | ✅      | Основная часть выполнена              |
| Добавление даты и приоритета                                 | ✅      | Бонус реализован                      |
| Удаление задач                                               | ✅      | С подтверждением (модальное окно)     |
| Отметка задачи как выполненной                               | ✅      | Зачеркивание текста выполненных задач |
| Перемещение выполненных задач в отдельный раздел             | ⚪      | Не реализовано                        |
| Сохранение состояния                                         | ✅      | Через PostgreSQL                      |
| Фильтрация задач по статусу                                  | ✅      | Все / Активные / Выполненные          |
| Сортировка задач                                             | ✅      | По дате и приоритету                  |


**Video Link: https://drive.google.com/file/d/1tTJe-2oaIlfo_qY94ptTZLUSbu50uPy8/view?usp=sharing

<img  alt="image" src="https://github.com/user-attachments/assets/eb950b9f-8c1c-4876-981d-c8cf7cbb1a42" />
<img  alt="image" src="https://github.com/user-attachments/assets/c770f859-8307-4810-8939-be79a81171ba" />
<img  alt="image" src="https://github.com/user-attachments/assets/7ccbf50a-6a61-408f-8292-819b7969ebaa" />
<img  alt="image" src="https://github.com/user-attachments/assets/99c7f9ea-0af1-42e5-8e39-91fa1272257f" />
<img  alt="image" src="https://github.com/user-attachments/assets/ce2ee541-3a7c-4952-9b53-6c1fba1f6b50" />
<img  alt="image" src="https://github.com/user-attachments/assets/33931789-1044-45d4-8c23-3c122b05ece6" />
<img  alt="image" src="https://github.com/user-attachments/assets/75b24af0-a4d1-416f-b3eb-17135eb18e6c" />
<img  alt="image" src="https://github.com/user-attachments/assets/745fc6ae-798d-42d6-a440-8e2421347166" />



Требования:

Go >= 1.20

Node.js >= 20


*Запуск проекта

```bash
git clone https://github.com/mdqni/DMARK-TODO.git
export GOSUMDB=sum.golang.org
export PATH=$PATH:~/go/bin
cd DMARK-TODO
go install github.com/wailsapp/wails/v2/cmd/wails@latest
npm install prop-types
docker compose up --build
wails dev
```
