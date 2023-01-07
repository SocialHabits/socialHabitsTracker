import Link from 'next/link';
import { ReactNode } from 'react';

export const LandingPageLayout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <header>
        <nav className='flex'>
          <div>Logo</div>

          <ul>
            <li>
              <Link href='/'>Home</Link>
            </li>
          </ul>
        </nav>
      </header>
      <main>{children}</main>
    </>
  );
};
