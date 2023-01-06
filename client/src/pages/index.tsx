import * as React from 'react';

import { useGraphQLClient, useQueryStatusLogger } from '@/hooks';

import Layout from '@/components/layout/Layout';
import ButtonLink from '@/components/links/ButtonLink';
import UnderlineLink from '@/components/links/UnderlineLink';
import Seo from '@/components/Seo';

import { GetUsersDocument, useGetUsers } from '@/graphql';
import { dehydrate, QueryClient } from '@tanstack/query-core';
import { GraphQLClient } from 'graphql-request';

// !STARTERCONF -> Select !STARTERCONF and CMD + SHIFT + F
// Before you begin editing, follow all comments with `STARTERCONF`,
// to customize the default configuration.

export default function HomePage() {
  const { graphQLClient } = useGraphQLClient();
  const { data, isFetching } = useGetUsers(
    graphQLClient,
    {},
    {
      useErrorBoundary: true,
      onSuccess: () => {
        // eslint-disable-next-line no-console
        console.log(Date.now(), `Fetching users types succeeded`);
      },
    }
  );

  useQueryStatusLogger({ isFetching }, 'users');

  return (
    <Layout>
      {/* <Seo templateTitle='Home' /> */}
      <Seo />

      <main>
        <section className='bg-white'>
          <div className='layout flex min-h-screen flex-col items-center justify-center text-center'>
            <ButtonLink className='mt-6' href='/components' variant='light'>
              See all components
            </ButtonLink>

            <div>
              {data?.getUsers?.map((user) => (
                <div key={user.id}>{user.firstName}</div>
              ))}
            </div>

            <footer className='absolute bottom-2 text-gray-700'>
              © {new Date().getFullYear()} By{' '}
              <UnderlineLink href='https://theodorusclarence.com?ref=tsnextstarter'>
                Theodorus Clarence
              </UnderlineLink>
            </footer>
          </div>
        </section>
      </main>
    </Layout>
  );
}

export async function getServerSideProps() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery(['GetUsers', {}], async () => {
    const graphQLClient = new GraphQLClient('http://localhost:8080/query');
    const { getUsers } = await graphQLClient.request(GetUsersDocument);

    return getUsers;
  });

  return {
    props: {
      dehydratedState: dehydrate(queryClient),
    },
  };
}
