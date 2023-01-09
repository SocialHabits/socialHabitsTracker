import { ReactElement } from 'react';
import * as React from 'react';

import SignUpForm from '@/components/form/SignUpForm';
import AppIcon from '@/components/Icons/AppIcon';
import { LandingPageLayout } from '@/components/layout/LandingPageLayout';

export default function SignUp() {
  return (
    <section className='flex h-[calc(100vh-80px)]'>
      <div className='flex w-full px-6 pt-10 md:items-center md:p-10 md:pt-0 xl:w-1/2'>
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
