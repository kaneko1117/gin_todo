"use client";

import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { TodoList } from "@/features/TodoList";
import { TodoData } from "@/features/TodoList/type";

const client = new ApolloClient({
  uri: "http://localhost:8080/query",
  cache: new InMemoryCache(),
});

export default function Todo() {
  const TEST_DATA: TodoData = [
    { id: "1", task: "test1", checked: false },
    { id: "2", task: "test2", checked: true },
    { id: "3", task: "test3", checked: false },
    { id: "4", task: "test4", checked: true },
    { id: "5", task: "test5", checked: false },
  ];

  return (
    <ApolloProvider client={client}>
      <TodoList data={TEST_DATA} />/
    </ApolloProvider>
  );
}
