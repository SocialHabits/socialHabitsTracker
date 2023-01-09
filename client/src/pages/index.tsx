import { dehydrate, QueryClient } from '@tanstack/query-core';
import { GraphQLClient } from 'graphql-request';
import * as React from 'react';
import { ReactElement } from 'react';

import { useGraphQLClient, useQueryStatusLogger } from '@/hooks';

import { LandingPageLayout } from '@/components/layout/LandingPageLayout';
import ButtonLink from '@/components/links/ButtonLink';
import UnderlineLink from '@/components/links/UnderlineLink';
import Seo from '@/components/Seo';

import { useGetUsers } from '@/graphql';

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
    <>
      <Seo />

      <main>
        <section className='bg-white'>
          <div className='layout flex min-h-screen flex-col items-center justify-center text-center'>
            <ButtonLink className='mt-6' href='/components' variant='light'>
              See all components
            </ButtonLink>

            {data?.getUsers && (
              <div>
                {data.getUsers.map((user) => (
                  <div key={user.id}>
                    <p>Name: {user.firstName}</p>
                    <p>Role: {user.role}</p>
                  </div>
                ))}
              </div>
            )}
          </div>
        </section>
      </main>
    </>
  );
}

HomePage.getLayout = function (page: ReactElement) {
  return <LandingPageLayout>{page}</LandingPageLayout>;
};

export async function getServerSideProps() {
  const queryClient = new QueryClient();
  const graphQLClient = new GraphQLClient('http://localhost:8080/query');

  await queryClient.prefetchQuery(
    useGetUsers.getKey({}),
    useGetUsers.fetcher(graphQLClient, {})
  );

  return {
    props: {
      dehydratedState: dehydrate(queryClient),
    },
  };
}
