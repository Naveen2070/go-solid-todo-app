import { createQuery, createMutation } from '@tanstack/solid-query';
import './App.css';
import AddTask from './components/add-task/AddTasks';
import TaskList from './components/task-list/TaskList';
import { Switch, Match } from 'solid-js';

export const BASE_URL =
  import.meta.env.MODE === 'development' ? 'http://localhost:3003/api' : '/api';

const fetchTasks = async () => {
  const res = await fetch(`${BASE_URL}/todos`);
  if (!res.ok) throw new Error('Failed to fetch tasks');
  return res.json();
};

const App = () => {
  const todoQuery = createQuery(() => ({
    queryKey: ['todos'],
    queryFn: fetchTasks,
  }));

  const addTaskMutation = createMutation(() => ({
    mutationKey: ['todos-add'],
    mutationFn: async (newTask: string) => {
      const res = await fetch(`${BASE_URL}/add-todos`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ body: newTask }),
      });
      if (!res.ok) throw new Error('Failed to add task');
    },
    onSuccess: () => todoQuery.refetch(),
  }));

  const updateTaskMutation = createMutation(() => ({
    mutationKey: ['todos-update'],
    mutationFn: async (id: number) => {
      const res = await fetch(`${BASE_URL}/todos/${id}/complete`, {
        method: 'GET',
      });
      if (!res.ok) throw new Error('Failed to update task');
    },
    onSuccess: () => todoQuery.refetch(),
  }));

  const deleteTaskMutation = createMutation(() => ({
    mutationKey: ['todos-delete'],
    mutationFn: async (id: number) => {
      const res = await fetch(`${BASE_URL}/delete-todos/${id}`, {
        method: 'DELETE',
      });
      if (!res.ok) throw new Error('Failed to delete task');
    },
    onSuccess: () => todoQuery.refetch(),
  }));

  // Handlers
  const addTask = (body: string) => {
    addTaskMutation.mutate(body);
  };

  const toggleStatus = (id: number) => {
    updateTaskMutation.mutate(id);
  };

  const deleteTask = (id: number) => {
    deleteTaskMutation.mutate(id);
  };

  return (
    <div class="app" id={todoQuery.isLoading ? 'loading' : ''}>
      <AddTask onAddTask={addTask} />
      <div class="header">
        <h1>Today's Tasks</h1>
      </div>
      <Switch>
        <Match when={todoQuery.isFetching || todoQuery.isLoading}>
          <div class="spinner"></div>
        </Match>
        <Match when={todoQuery.isError}>
          Error: {todoQuery.error?.message}
        </Match>
        <Match when={todoQuery.isSuccess}>
          <TaskList
            tasks={todoQuery.data}
            onToggleStatus={toggleStatus}
            onDeleteTask={deleteTask}
          />
        </Match>
      </Switch>
    </div>
  );
};

export default App;
