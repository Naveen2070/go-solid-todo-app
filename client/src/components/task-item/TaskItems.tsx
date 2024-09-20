import './TaskItems.css';

interface TaskItemProps {
  key: number;
  task: {
    id: number;
    body: string;
    isCompleted: boolean;
  };
  onToggleStatus: () => void;
  onDeleteTask: () => void;
}

const TaskItem = (props: TaskItemProps) => {
  return (
    <div class="task-item" id={props.task.id.toString()}>
      <span class={`task-body ${props.task.isCompleted ? 'done' : ''}`}>
        {props.task.body}
      </span>
      <span class={`status ${props.task.isCompleted ? 'done' : 'in-progress'}`}>
        {props.task.isCompleted ? 'DONE' : 'IN PROGRESS'}
      </span>
      <div class="btn-container">
        <button class="check-btn" onClick={props.onToggleStatus}>
          ✔️
        </button>
        <button class="delete-btn" onClick={props.onDeleteTask}>
          🗑️
        </button>
      </div>
    </div>
  );
};

export default TaskItem;
