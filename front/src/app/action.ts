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
