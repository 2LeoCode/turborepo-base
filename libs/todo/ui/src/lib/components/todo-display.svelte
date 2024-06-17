<script lang="ts">
  import type { TodoEntry, TodoList } from "@repo/libs-todo-ts";
  import TodoAddModal from "./todo-add-modal.svelte";
  import TodoRemoveModal from "./todo-remove-modal.svelte";
  import { onMount } from "svelte";

  export let todoList: TodoList;

  let mounted = false;
  onMount(() => (mounted = true));

  $: todos = mounted ? todoList.getTodos() : new Promise<TodoEntry[]>(() => {});
</script>

{#await todos}
  <div>Loading...</div>
{:then todos}
  <ul>
    {#each todos as { name, description }}
      <li id={name}>
        {name}: {description}
      </li>
    {/each}
  </ul>
{:catch err}
  <div>
    {(() => {
      console.error(err);
      return "An error occured while loading todo-list";
    })()}
  </div>
{/await}

<TodoAddModal bind:todoList />
<TodoRemoveModal bind:todoList />
<button
  on:click={async () => {
    await todoList.saveTodos();
    alert("todo-list saved");
  }}>Save todos</button
>
