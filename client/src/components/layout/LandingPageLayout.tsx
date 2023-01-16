import { useRouter } from 'next/router';
import { ReactNode } from 'react';

import ButtonLink from '@/components/links/ButtonLink';
import UnderlineLink from '@/components/links/UnderlineLink';

export const LandingPageLayout = ({ children }: { children: ReactNode }) => {
  const router = useRouter();

  return (
    <>
      <header>
        <nav className='flex h-20 items-center justify-between px-8'>
          <div className='w-10'>Logo</div>

          <div className='hidden items-center justify-between md:flex'>
            <ul className='flex'>
              <li className='px-4'>
                <UnderlineLink href='/' className='text-lg font-medium'>
                  Home
                </UnderlineLink>
              </li>

              <li className='px-4'>
                <UnderlineLink href='/about' className='text-lg font-medium'>
                  About
                </UnderlineLink>
              </li>
            </ul>

            <div className='ml-40'>
              {router.route !== '/auth/sign-up' && (
                <ButtonLink
                  href='/auth/sign-up'
                  variant='outline'
                  className='mr-4'
                >
                  Sign Up
                </ButtonLink>
              )}

              {router.route !== '/auth/login' && (
                <ButtonLink href='/auth/login'>Log in</ButtonLink>
              )}
            </div>
          </div>
        </nav>
      </header>
      <main>{children}</main>
    </>
  );
};
