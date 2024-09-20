import { createSignal } from 'solid-js';
import './AddTasks.css';

interface AddTaskProps {
  onAddTask: (taskBody: string) => void;
}

const AddTask = (props: AddTaskProps) => {
  const [taskInput, setTaskInput] = createSignal('');

  const handleAddTask = () => {
    if (taskInput()) {
      props.onAddTask(taskInput());
      setTaskInput(''); // Clear input field after adding
    }
  };

  return (
    <div class="add-task">
      <input
        type="text"
        placeholder="Add a new task"
        value={taskInput()}
        onInput={(e) => setTaskInput(e.currentTarget.value)}
      />
      <button class="add-btn" onClick={handleAddTask} disabled={!taskInput()}>
        +
      </button>
    </div>
  );
};

export default AddTask;
