import {useState} from "react";
import TaskItem from "./TaskItem";
import {dto} from "../../../../wailsjs/go/models";

interface TaskListProps {
    tasks: dto.TaskDTO[];
    onToggle: (id: number) => void;
    onDelete: (id: number) => void;
    onFilterChange: (status: string, dateFilter: string) => void;
}

export default function TaskList({tasks, onToggle, onDelete, onFilterChange}: TaskListProps) {
    const [statusFilter, setStatusFilter] = useState("all");
    const [dateFilter, setDateFilter] = useState("all");

    const handleFilterChange = (newStatus: string, newDate: string) => {
        setStatusFilter(newStatus);
        setDateFilter(newDate);
        onFilterChange(newStatus, newDate);
    };

    return (
        <div className="w-full max-w-md mt-6">
            <div className="flex justify-between mb-4 space-x-2">
                <select
                    className="border px-2 py-1 rounded"
                    value={statusFilter}
                    onChange={(e) => handleFilterChange(e.target.value, dateFilter)}
                >
                    <option value="all">All</option>
                    <option value="active">Active</option>
                    <option value="done">Done</option>
                </select>

                <select
                    className="border px-2 py-1 rounded"
                    value={dateFilter}
                    onChange={(e) => handleFilterChange(statusFilter, e.target.value)}
                >
                    <option value="all">All</option>
                    <option value="today">Today</option>
                    <option value="week">This Week</option>
                    <option value="overdue">Overdue</option>
                </select>
            </div>

            {(!tasks || tasks.length === 0) ? (
                <p className="text-gray-500">No tasks yet</p>
            ) : (
                <ul className="space-y-2">
                    {tasks.map((t) => (
                        <TaskItem key={t.id} task={t} onToggle={onToggle} onDelete={onDelete}/>
                    ))}
                </ul>
            )}
        </div>
    );
}
