import {FormEventHandler, useEffect, useRef, useState} from "react";

import {TodoApiService} from "../services/todo-api-service.ts";
import {TodoResponse} from "../utils/TodoResponse.ts";
import {TodoListItem} from "./TodoListItem.tsx";

export const TodoList = () => {
  const [todos, setTodos] = useState<TodoResponse[]>([]);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    TodoApiService
      .fetchTodos()
      .then(todos => setTodos(todos ?? []))
      .catch(console.warn)
  }, []);

  const completeTodo = (id: string) => {
    TodoApiService.deleteTodo(id)
      .then(() => setTodos(todos => todos.filter(t => t.id !== id)))
      .catch(console.warn);
  }

  const submit: FormEventHandler = $event => {
    $event.preventDefault();
    const inputValue = inputRef.current?.value.trim();

    if (!inputValue)
      return;

    TodoApiService.createTodo({title: inputValue})
      .then((res) => res && setTodos(todos => [...todos, {id: res, title: inputValue}]))
      .catch(console.warn);

    inputRef.current!.value = "";
  }

  return <>
    <form onSubmit={submit} >
      <input type="text" ref={inputRef} />
      <input type="submit" value="Add" />
    </form>
    {todos.map(t => <TodoListItem key={t.id} todo={t} onClick={() => completeTodo(t.id)} />)}
  </>
}