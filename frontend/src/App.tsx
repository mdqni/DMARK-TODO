import {useEffect, useState} from "react";
import {ListTasks, ToggleDone, DeleteTask} from "../wailsjs/go/usecase/TaskUseCase";
import {dto} from "../wailsjs/go/models";
import TaskForm from "./features/tasks/components/TaskForm";
import TaskList from "./features/tasks/components/TaskList";

function App() {
    const [tasks, setTasks] = useState<dto.TaskDTO[]>([]);
    const [statusFilter, setStatusFilter] = useState("all");
    const [dateFilter, setDateFilter] = useState("all");

    const fetchTasks = async (status = statusFilter, date = dateFilter) => {
        try {
            const result = await ListTasks(status, date);
            setTasks(result || []);
        } catch (err) {
            console.error("Failed to fetch tasks:", err);
            setTasks([]);
        }
    };

    const handleFilterChange = (status: string, date: string) => {
        setStatusFilter(status);
        setDateFilter(date);
        fetchTasks(status, date);
    };

    useEffect(() => {
        fetchTasks();
    }, []);

    const handleToggle = async (id: number) => {
        try {
            await ToggleDone(id.toString());
            fetchTasks();
        } catch (err) {
            console.error("Toggle failed:", err);
        }
    };

    const handleDelete = async (id: number) => {
        try {
            await DeleteTask(id.toString());
            fetchTasks();
        } catch (err) {
            console.error("Delete failed:", err);
        }
    };

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col items-center py-10 px-4">
            <h1 className="text-3xl text-gray-900 font-bold mb-6">Task Manager</h1>
            <TaskForm onTaskAdded={fetchTasks}/>
            <TaskList tasks={tasks} onToggle={handleToggle} onDelete={handleDelete}
                      onFilterChange={handleFilterChange}/>
        </div>
    );
}

export default App;
