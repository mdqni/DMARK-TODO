import {
    SwipeableList,
    SwipeableListItem,
    SwipeAction,
    TrailingActions,
    LeadingActions,
} from "react-swipeable-list";
import "react-swipeable-list/dist/styles.css";
import {dto} from "../../../../wailsjs/go/models";

interface TaskItemProps {
    task: dto.TaskDTO;
    onToggle: (id: number) => void;
    onDelete: (id: number) => void;
}

export default function TaskItem({task, onToggle, onDelete}: TaskItemProps) {

    const handleDelete = async () => {
        const confirmed = window.confirm(`Удалить задачу "${task.title}"?`);
        if (!confirmed) return;

        try {
            await onDelete(task.id);
        } catch (err) {
            console.error("Ошибка при удалении задачи:", err);
        }
    };

    const leadingActions = () => (
        <LeadingActions>
            <SwipeAction onClick={() => onToggle(task.id)}>
                <div
                    className={`flex items-center justify-center h-full w-full text-white font-bold ${
                        task.completed ? "bg-yellow-500" : "bg-green-500"
                    }`}
                >
                    {task.completed ? "Mark as Active" : "Mark as Done"}
                </div>
            </SwipeAction>
        </LeadingActions>
    );

    const trailingActions = () => (
        <TrailingActions>
            <SwipeAction destructive={false} onClick={() => handleDelete()}>
                <div className="flex items-center justify-center h-full w-full bg-red-500 text-white font-bold">
                    Delete
                </div>
            </SwipeAction>
        </TrailingActions>
    );

    return (
        <SwipeableList threshold={0.25}>
            <SwipeableListItem
                leadingActions={leadingActions()}
                trailingActions={trailingActions()}
            >
                <div className="bg-white p-4 rounded-lg shadow flex flex-col w-full">
                    <span
                        className={`font-semibold ${
                            task.completed ? "line-through text-gray-400" : ""
                        }`}
                    >
                        {task.title}
                    </span>
                    {task.description && (
                        <span className="text-sm text-gray-600">{task.description}</span>
                    )}
                    <span className="text-xs text-gray-500">
                        Priority: {task.priority} | Due:{" "}
                        {task.due_date
                            ? new Date(task.due_date).toLocaleDateString()
                            : "—"}
                    </span>
                </div>
            </SwipeableListItem>
        </SwipeableList>
    );
}
