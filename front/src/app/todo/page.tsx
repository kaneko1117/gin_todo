import { TodoList } from "@/features/TodoList";
import { todoData } from "@/features/TodoList/type";

export default async function Todo() {
  const TEST_DATA: todoData = [
    { id: 1, task: "test1", checked: false },
    { id: 2, task: "test2", checked: true },
    { id: 3, task: "test3", checked: false },
    { id: 4, task: "test4", checked: true },
    { id: 5, task: "test5", checked: false },
  ];
  return <TodoList data={TEST_DATA} />;
}
