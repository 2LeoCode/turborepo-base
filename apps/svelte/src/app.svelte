<script lang="ts">
  import type { TodoList } from "@repo/libs-todo-ts";
  import { TodoDisplay } from "@repo/libs-todo-ui";

  let todoList: TodoList = {
    addTodo(name, description) {
      return fetch("http://localhost:3000", {
        mode: "cors",
        method: "POST",
        body: JSON.stringify(["AddTodo", name, description]),
      }).then(async (res) => {
        if (!res.ok)
          throw new Error("AddTodo error", { cause: res.statusText });
        const result = await res.json();
        if (result === false)
          throw new Error("AddTodo error", { cause: "todo already exists" });
        todoList = todoList; // trigger a re-render to components depending on todoList
        return result;
      });
    },
    removeTodo(name) {
      return fetch("http://localhost:3000", {
        mode: "cors",
        method: "POST",
        body: JSON.stringify(["RemoveTodo", name]),
      }).then(async (res) => {
        if (!res.ok)
          throw new Error(`RemoveTodo error`, { cause: res.statusText });
        const result = await res.json();
        if (result === false)
          throw new Error("AddTodo error", { cause: "todo does not exist" });
        todoList = todoList; // trigger a re-render to components depending on todoList
        return result;
      });
    },
    getTodos() {
      return fetch("http://localhost:3000", {
        mode: "cors",
        method: "POST",
        body: JSON.stringify(["GetTodos"]),
      }).then(async (res) => {
        if (!res.ok)
          throw new Error(`GetTodos error`, { cause: res.statusText });
        return res.json();
      });
    },
    saveTodos() {
      return fetch("http://localhost:3000", {
        mode: "cors",
        method: "POST",
        body: JSON.stringify(["SaveTodos"]),
      }).then((res) => {
        if (!res.ok)
          throw new Error(`SaveTodos error`, { cause: res.statusText });
        return res.json();
      });
    },
  };
</script>

<main>
  <TodoDisplay {todoList} />
</main>
