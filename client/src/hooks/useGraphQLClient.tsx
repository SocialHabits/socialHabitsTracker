import { useContext } from 'react';

import { GraphQLClientContext } from '@/providers/GraphQLClientProvider';

export const useGraphQLClient = () => {
  const contextValue = useContext(GraphQLClientContext);

  if (!contextValue) {
    throw new Error(
      'Wrap your components tree with a GraphQLClientProvider component'
    );
  }

  return contextValue;
};

export default useGraphQLClient;
