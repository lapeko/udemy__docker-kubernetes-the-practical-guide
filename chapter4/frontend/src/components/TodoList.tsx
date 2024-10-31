import {useEffect, useState} from "react";
import {TodoApiService} from "../services/todo-api-service.ts";
import {TodoResponse} from "../utils/TodoResponse.ts";
import {TodoListItem} from "./TodoListItem.tsx";

export const TodoList = () => {
  const [todos, setTodos] = useState<TodoResponse[]>([]);

  useEffect(() => {
    TodoApiService
      .fetchTodos()
      .then(todos => setTodos(todos ?? []))
      .catch(console.warn)
  }, []);

  return <>
    {/*<form>*/}
    {/*  */}
    {/*</form>*/}
    <TodoListItem todos={todos} />
  </>
}