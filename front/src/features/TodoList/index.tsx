"use client";

import { gql, useQuery } from "@apollo/client";
import { TodoData } from "./type";
import { TodoListForm } from "./components/TodoListForm";
import { TodoCheckBox } from "./components/TodoCheckBox";

type Props = {
  data: TodoData;
};

const TodosQuery = gql(`
  query GetTasks($id : ID!) {
    getTasks(id: $id) {
      task
    }
  }
`);

export const TodoList = ({ data }: Props) => {
  const { data: graphql } = useQuery(TodosQuery, {
    variables: {
      id: 1,
    },
  });
  console.log("graphql", graphql);
  return (
    <div className="flex flex-col items-center gap-5 h-screen mt-12">
      <div className="w-[500px]">
        <TodoListForm />
        <TodoCheckBox data={data} />
      </div>
    </div>
  );
};
