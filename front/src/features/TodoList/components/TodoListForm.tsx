"use client";
import { gql, useMutation } from "@apollo/client";
import { Button } from "@/components/button";
import { Input } from "@/components/input";

const REGISTER_TASK = gql(`
  mutation CreateTask($data: Todo!) {
    createTask(data: $data) {
      msg
    }
  }
`);

type Response = {
  msg: string;
};

export const TodoListForm = () => {
  const [registerTask, { loading, error }] =
    useMutation<Response>(REGISTER_TASK);
  const registerFormData = (formData: FormData) => {
    const task = formData.get("task");
    registerTask({
      variables: {
        data: {
          tasks: task,
          id: 2,
        },
      },
    });
  };

  return (
    <>
      <form className="flex gap-2" action={registerFormData}>
        <Input type="text" placeholder="タスクを追加" id="task" name="task" />
        <Button type="submit" className="cursor-pointer" disabled={loading}>
          追加
        </Button>
      </form>
      {error && (
        <div className="mt-2">
          <p className="text-red-500">登録に失敗しました</p>
        </div>
      )}
    </>
  );
};
