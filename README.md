# To-Do List Desktop App (Wails + Go + React)

## Описание
Десктопное приложение для управления задачами (To-Do List) с возможностью:
- Создавать задачи
- Удалять задачи с подтверждением
- Отмечать задачи как выполненные
- Фильтровать и сортировать задачи
- Сохранять состояние задач между запусками приложения
- Использовать PostgreSQL для хранения данных (если реализовано)

## Установка и запуск

1. Установить Go >= 1.20 и Node.js >= 20
2. Установить Wails:
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
```bash
git clone https://github.com/mdqni/DMARK-TODO.git
export GOSUMDB=sum.golang.org
export PATH=$PATH:~/go/bin
cd DMARK-TODO
wails dev
```
