import { ReactNode } from 'react';

export const LandingPageLayout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <header>Ovo je header</header>
      <main>{children}</main>
    </>
  );
};
