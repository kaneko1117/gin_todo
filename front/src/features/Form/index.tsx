"use client";

import { useRouter } from "next/navigation";
import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { gql, useMutation } from "@apollo/client";

const LOGIN = gql(
  `
  mutation Login($data: LoginData!) {
    login(data: $data) {
      msg
    }
  }
`
);

type Response = {
  msg: string;
};

export const Form = () => {
  const router = useRouter();
  const [login, { loading, error }] = useMutation<Response>(LOGIN);

  const loginFormData = async (formData: FormData) => {
    const username = formData.get("username");
    const password = formData.get("password");
    await login({
      variables: {
        data: {
          userName: username,
          password: password,
        },
      },
    });
    router.push("/todo");
  };
  return (
    <form
      className="flex flex-col items-center justify-center gap-5 h-screen"
      action={loginFormData}
    >
      <Input
        type="text"
        placeholder="ユーザー名"
        className="w-1/2"
        id="username"
        name="username"
      />
      <Input
        type="password"
        placeholder="パスワード"
        className="w-1/2"
        id="password"
        name="password"
      />
      <Button type="submit" disabled={loading} className="cursor-pointer">
        ログイン
      </Button>
      {error && (
        <div className="mt-2">
          <p className="text-red-500">ログインに失敗しました</p>
        </div>
      )}
    </form>
  );
};
