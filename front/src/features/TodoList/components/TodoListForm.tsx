"use client";
import { useActionState } from "react";
import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { registerTodoList } from "@/app/action";

const initialState = {
  message: "",
};
export const TodoListForm = () => {
  const [state, formAction, pending] = useActionState(
    registerTodoList,
    initialState
  );
  return (
    <>
      <form className="flex gap-2" action={formAction}>
        <Input type="text" placeholder="タスクを追加" id="task" name="task" />
        <Button type="submit" className="cursor-pointer" disabled={pending}>
          追加
        </Button>
      </form>
      <p className="h-6">{state.message}</p>
    </>
  );
};
