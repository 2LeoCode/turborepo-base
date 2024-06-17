export type TodoEntry = {
  name: string;
  description: string;
};

export type TodoList = {
  getTodos(): Promise<TodoEntry[]>;
  addTodo(name: string, description: string): Promise<boolean>;
  removeTodo(name: string): Promise<boolean>;
  saveTodos(): Promise<void>;
};
