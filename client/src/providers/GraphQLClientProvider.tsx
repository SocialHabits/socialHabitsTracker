import { GraphQLClient } from 'graphql-request';
import { createContext } from 'react';

type ClientProps = {
  children: React.ReactNode;
  defaultState?: GraphQLClientState;
};

export type GraphQLClientState = {
  graphQLClient: GraphQLClient;
};

export type GraphQLClientProviderState = GraphQLClientState | null;

export const GraphQLClientContext =
  createContext<GraphQLClientProviderState>(null);

const initialState: GraphQLClientState = {
  graphQLClient: new GraphQLClient('http://localhost:8080/query', {
    credentials: 'include',
  }),
};

export function GraphQLClientProvider({ children, defaultState }: ClientProps) {
  return (
    <GraphQLClientContext.Provider value={defaultState || initialState}>
      {children}
    </GraphQLClientContext.Provider>
  );
}
