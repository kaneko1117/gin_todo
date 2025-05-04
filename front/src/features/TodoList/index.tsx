"use client";

import { gql, useQuery } from "@apollo/client";
import { GraphQLResponse } from "./type";
import { TodoListForm } from "./components/TodoListForm";
import { TodoCheckBox } from "./components/TodoCheckBox";

const TodosQuery = gql(`
  query GetTasks($id : ID!) {
    getTasks(id: $id) {
      ID
      task
      isChecked
    }
  }
`);

export const TodoList = () => {
  const { data, loading, error, refetch } = useQuery<GraphQLResponse>(
    TodosQuery,
    {
      variables: {
        id: 1,
      },
    }
  );
  if (loading || !data) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;
  return (
    <div className="flex flex-col items-center gap-5 h-screen mt-12">
      <div className="w-[500px]">
        <TodoListForm refetch={refetch} />
        <TodoCheckBox data={data} refetch={refetch} />
      </div>
    </div>
  );
};
