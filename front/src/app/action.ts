"use server";
import { redirect } from "next/navigation";

export async function loginUser(_: unknown, formData: FormData) {
  const userName = formData.get("username");
  const password = formData.get("password");

  if (!userName || !password) {
    return { message: "入力してください" };
  }

  redirect("/todo");
}

export async function registerTodoList(
  _: { message: string },
  formData: FormData
) {
  const task = formData.get("task");
  if (!task) {
    return { message: "タスクを入力してください" };
  }

  return { message: "タスクを登録しました" };
}

export async function checkedTodoList(id: number, checked: boolean) {
  // ここでDBに登録する処理を実装する
  return { message: "タスクを登録しました" };
}
