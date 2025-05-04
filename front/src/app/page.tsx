"use client";

import { Form } from "@/features/Form";
import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";

const client = new ApolloClient({
  uri: "http://localhost/back/query",
  cache: new InMemoryCache(),
});

export default function Login() {
  return (
    <ApolloProvider client={client}>
      <Form />;
    </ApolloProvider>
  );
}
