"use client";

import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { TodoList } from "@/features/TodoList";
import { TodoData } from "@/features/TodoList/type";

const client = new ApolloClient({
  uri: "http://localhost:8080/query",
  cache: new InMemoryCache(),
});

export default function Todo() {
  return (
    <ApolloProvider client={client}>
      <TodoList />
    </ApolloProvider>
  );
}
