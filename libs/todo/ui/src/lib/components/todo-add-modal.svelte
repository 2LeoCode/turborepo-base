<script lang="ts">
  import type { TodoList } from "@repo/libs-todo-ts";

  export let todoList: TodoList;

  let name = "";
  let description = "";

  let dialog: HTMLDialogElement;

  function onSubmit() {
    return todoList
      .addTodo(name, description)
      .catch((err: Error) => alert(`${err.message}: ${err.cause}`));
  }
</script>

<button on:click={() => dialog.showModal()}>Add todo</button>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
  bind:this={dialog}
  on:click|self={() => dialog.close()}
  on:close|preventDefault
>
  <form method="dialog" on:submit={onSubmit}>
    <div>
      <label for="name">Todo name: </label>
      <input type="text" name="name" bind:value={name} required />
    </div>
    <div>
      <label for="description">Todo description: </label>
      <input type="text" name="description" bind:value={description} required />
    </div>
    <input type="submit" value="Add todo" />
  </form>
</dialog>
