import { useEffect } from 'react';

const useQueryStatusLogging = (
  {
    isFetching,
  }: {
    isFetching: boolean;
  },
  text: string
) => {
  useEffect(() => {
    if (isFetching) {
      console.log(Date.now(), `Fetching ${text}...`);
    }
  }, [isFetching, text]);
};

export default useQueryStatusLogging;
