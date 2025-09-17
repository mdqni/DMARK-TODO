import {useState} from "react";
import TaskItem from "./TaskItem";
import {dto} from "../../../../wailsjs/go/models";

interface TaskListProps {
    tasks: dto.TaskDTO[];
    onToggle: (id: number) => void;
    onDelete: (id: number) => void;
    onFilterChange: (status: string, dateFilter: string, sortBy: string) => void;
}

export default function TaskList({tasks, onToggle, onDelete, onFilterChange}: TaskListProps) {
    const [statusFilter, setStatusFilter] = useState("all");
    const [dateFilter, setDateFilter] = useState("all");
    const [sortBy, setSortBy] = useState("created");

    const handleFilterSortChange = (newStatus: string, newDate: string, newSort: string) => {
        setStatusFilter(newStatus);
        setDateFilter(newDate);
        setSortBy(newSort);
        onFilterChange(newStatus, newDate, newSort);
    };

    return (
        <div className="w-full max-w-md mt-6">
            <div className="flex justify-between mb-4 space-x-2">
                <select
                    className="border px-2 py-1 rounded"
                    value={statusFilter}
                    onChange={(e) => handleFilterSortChange(e.target.value, dateFilter, sortBy)}
                >
                    <option value="all">All</option>
                    <option value="active">Active</option>
                    <option value="done">Done</option>
                </select>

                <select
                    className="border px-2 py-1 rounded"
                    value={dateFilter}
                    onChange={(e) => handleFilterSortChange(statusFilter, e.target.value, sortBy)}
                >
                    <option value="all">All</option>
                    <option value="today">Today</option>
                    <option value="week">This Week</option>
                    <option value="overdue">Overdue</option>
                </select>

                <select
                    className="border px-2 py-1 rounded"
                    value={sortBy}
                    onChange={(e) => handleFilterSortChange(statusFilter, dateFilter, e.target.value)}
                >
                    <option value="created">Created</option>
                    <option value="due">Due Date</option>
                    <option value="priority">Priority</option>
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
