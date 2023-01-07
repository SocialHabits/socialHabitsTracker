import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import {
  useMutation,
  useQuery,
  useInfiniteQuery,
  UseMutationOptions,
  UseQueryOptions,
  UseInfiniteQueryOptions,
} from '@tanstack/react-query';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};

function fetcher<TData, TVariables extends { [key: string]: any }>(
  client: GraphQLClient,
  query: string,
  variables?: TVariables,
  requestHeaders?: RequestInit['headers']
) {
  return async (): Promise<TData> =>
    client.request({
      document: query,
      variables,
      requestHeaders,
    });
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Any: any;
};

export type Address = {
  __typename?: 'Address';
  city: Scalars['String'];
  country: Scalars['String'];
  id: Scalars['ID'];
  street: Scalars['String'];
  userId: Scalars['ID'];
};

export type AddressInput = {
  city: Scalars['String'];
  country: Scalars['String'];
  street: Scalars['String'];
};

export type Book = {
  __typename?: 'Book';
  author: Scalars['String'];
  id: Scalars['Int'];
  publisher: Scalars['String'];
  title: Scalars['String'];
};

export type BookInput = {
  author: Scalars['String'];
  publisher: Scalars['String'];
  title: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  CreateBook: Book;
  DeleteBook: Scalars['String'];
  UpdateBook: Scalars['String'];
  createTodo: Todo;
  createUser: User;
  deleteTodo: Todo;
  deleteUser: User;
  login: Scalars['Any'];
  updateTodo: Todo;
  updateUser: User;
};

export type MutationCreateBookArgs = {
  input: BookInput;
};

export type MutationDeleteBookArgs = {
  id: Scalars['Int'];
};

export type MutationUpdateBookArgs = {
  id: Scalars['Int'];
  input: BookInput;
};

export type MutationCreateTodoArgs = {
  text: Scalars['String'];
};

export type MutationCreateUserArgs = {
  input: UserInput;
};

export type MutationDeleteTodoArgs = {
  todoId: Scalars['ID'];
};

export type MutationDeleteUserArgs = {
  id: Scalars['ID'];
};

export type MutationLoginArgs = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type MutationUpdateTodoArgs = {
  input: TodoInput;
};

export type MutationUpdateUserArgs = {
  id: Scalars['ID'];
  input: UserInput;
};

export type Query = {
  __typename?: 'Query';
  GetAllBooks: Array<Book>;
  GetOneBook: Book;
  getRole: Role;
  getTodo: Todo;
  getTodos: Array<Todo>;
  getUser: User;
  getUsers: Array<User>;
};

export type QueryGetOneBookArgs = {
  id: Scalars['Int'];
};

export type QueryGetRoleArgs = {
  id: Scalars['Int'];
};

export type QueryGetTodoArgs = {
  todoId: Scalars['ID'];
};

export type QueryGetUserArgs = {
  id: Scalars['ID'];
};

export enum Role {
  Admin = 'ADMIN',
  Premium = 'PREMIUM',
  Regular = 'REGULAR',
  Trainer = 'TRAINER',
}

export type RoleInput = {
  name: Role;
};

export type Todo = {
  __typename?: 'Todo';
  done: Scalars['Boolean'];
  id: Scalars['ID'];
  text: Scalars['String'];
};

export type TodoInput = {
  done: Scalars['Boolean'];
  id: Scalars['ID'];
  text: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  address: Array<Address>;
  email: Scalars['String'];
  firstName: Scalars['String'];
  id: Scalars['ID'];
  lastName: Scalars['String'];
  password: Scalars['String'];
  role: Role;
};

export type UserInput = {
  address: Array<AddressInput>;
  email: Scalars['String'];
  firstName: Scalars['String'];
  lastName: Scalars['String'];
  password: Scalars['String'];
  role: Role;
};

export type LoginVariables = Exact<{
  email: Scalars['String'];
  password: Scalars['String'];
}>;

export type Login = { __typename?: 'Mutation'; login: any };

export type GetUsersVariables = Exact<{ [key: string]: never }>;

export type GetUsers = {
  __typename?: 'Query';
  getUsers: Array<{
    __typename?: 'User';
    id: string;
    email: string;
    firstName: string;
    role: Role;
  }>;
};

export const LoginDocument = /*#__PURE__*/ `
    mutation Login($email: String!, $password: String!) {
  login(email: $email, password: $password)
}
    `;
export const useLogin = <TError = unknown, TContext = unknown>(
  client: GraphQLClient,
  options?: UseMutationOptions<Login, TError, LoginVariables, TContext>,
  headers?: RequestInit['headers']
) =>
  useMutation<Login, TError, LoginVariables, TContext>(
    ['Login'],
    (variables?: LoginVariables) =>
      fetcher<Login, LoginVariables>(
        client,
        LoginDocument,
        variables,
        headers
      )(),
    options
  );
useLogin.fetcher = (
  client: GraphQLClient,
  variables: LoginVariables,
  headers?: RequestInit['headers']
) => fetcher<Login, LoginVariables>(client, LoginDocument, variables, headers);
export const GetUsersDocument = /*#__PURE__*/ `
    query GetUsers {
  getUsers {
    id
    email
    firstName
    role
  }
}
    `;
export const useGetUsers = <TData = GetUsers, TError = unknown>(
  client: GraphQLClient,
  variables?: GetUsersVariables,
  options?: UseQueryOptions<GetUsers, TError, TData>,
  headers?: RequestInit['headers']
) =>
  useQuery<GetUsers, TError, TData>(
    variables === undefined ? ['GetUsers'] : ['GetUsers', variables],
    fetcher<GetUsers, GetUsersVariables>(
      client,
      GetUsersDocument,
      variables,
      headers
    ),
    options
  );

useGetUsers.getKey = (variables?: GetUsersVariables) =>
  variables === undefined ? ['GetUsers'] : ['GetUsers', variables];
export const useInfiniteGetUsers = <TData = GetUsers, TError = unknown>(
  pageParamKey: keyof GetUsersVariables,
  client: GraphQLClient,
  variables?: GetUsersVariables,
  options?: UseInfiniteQueryOptions<GetUsers, TError, TData>,
  headers?: RequestInit['headers']
) =>
  useInfiniteQuery<GetUsers, TError, TData>(
    variables === undefined
      ? ['GetUsers.infinite']
      : ['GetUsers.infinite', variables],
    (metaData) =>
      fetcher<GetUsers, GetUsersVariables>(
        client,
        GetUsersDocument,
        {
          ...variables,
          ...(metaData.pageParam ? { [pageParamKey]: metaData.pageParam } : {}),
        },
        headers
      )(),
    options
  );

useInfiniteGetUsers.getKey = (variables?: GetUsersVariables) =>
  variables === undefined
    ? ['GetUsers.infinite']
    : ['GetUsers.infinite', variables];
useGetUsers.fetcher = (
  client: GraphQLClient,
  variables?: GetUsersVariables,
  headers?: RequestInit['headers']
) =>
  fetcher<GetUsers, GetUsersVariables>(
    client,
    GetUsersDocument,
    variables,
    headers
  );
