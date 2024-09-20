import { For } from 'solid-js';
import TaskItem from '../task-item/TaskItems';
import './TaskList.css';

interface Task {
  id: number; // Added ID to track task
  body: string;
  isCompleted: boolean;
}

interface TaskListProps {
  tasks: Task[];
  onToggleStatus: (id: number) => void; // Updated to take id and status
  onDeleteTask: (id: number) => void;
}

const TaskList = ({ tasks, onToggleStatus, onDeleteTask }: TaskListProps) => {
  if (tasks.length === 0) {
    return (
      <div class="task-list">
        <p>No tasks found. Add a new task to get started.</p>
      </div>
    );
  }

  return (
    <For each={tasks}>
      {(task) => (
        <div class="task-list">
          <TaskItem
            key={task.id} // Added key for better performance
            task={task}
            onToggleStatus={() => onToggleStatus(task.id)}
            onDeleteTask={() => onDeleteTask(task.id)}
          />
        </div>
      )}
    </For>
  );
};

export default TaskList;
