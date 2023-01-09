import * as React from 'react';

import SignUpForm from '@/components/form/SignUpForm';
import AppIcon from '@/components/Icons/AppIcon';
import ButtonLink from '@/components/links/ButtonLink';

export default function SignUp() {
  return (
    <section className='relative flex h-screen justify-center md:mt-0'>
      <div className='absolute top-4 left-4'>LOGO</div>

      <ButtonLink href='/auth/login' className='absolute top-4 right-4'>
        Log in
      </ButtonLink>

      <div className='mt-10 w-full p-8 md:flex md:items-center md:bg-transparent xl:w-1/2'>
        <SignUpForm />
      </div>

      <div className='hidden w-1/2 items-center justify-center bg-gradient-to-r from-cyan-100 to-primary-300 xl:flex'>
        <div>
          <AppIcon />
        </div>
      </div>
    </section>
  );
}
