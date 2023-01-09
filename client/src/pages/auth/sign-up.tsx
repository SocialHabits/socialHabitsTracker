import * as React from 'react';

import SignUpForm from '@/components/form/SignUpForm';
import AppIcon from '@/components/Icons/AppIcon';

export default function SignUp() {
  return (
    <section className='flex h-screen justify-center md:mt-0'>
      <div className='mt-10 w-full bg-gray-50 md:mt-0 md:flex md:items-center md:bg-transparent md:p-0 xl:w-1/2'>
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
