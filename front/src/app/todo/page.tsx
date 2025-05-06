"use client";

import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { TodoList } from "@/features/TodoList";

const client = new ApolloClient({
  uri: "http://localhost/back/query",
  cache: new InMemoryCache(),
});

export default function Todo() {
  return (
    <ApolloProvider client={client}>
      <TodoList />
    </ApolloProvider>
  );
}
