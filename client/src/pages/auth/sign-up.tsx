import SignUpForm from '@/components/form/SignUpForm';
import { ReactElement } from 'react';
import { LandingPageLayout } from '@/components/layout/LandingPageLayout';
import * as React from 'react';

export default function SignUp() {
  return (
    <section className='flex h-[calc(100vh-80px)]'>
      <div className='flex w-1/2 items-center'>
        <SignUpForm />
      </div>

      <div className='w-1/2 bg-gradient-to-r from-red-50 to-blue-100'></div>
    </section>
  );
}

SignUp.getLayout = function (page: ReactElement) {
  return <LandingPageLayout>{page}</LandingPageLayout>;
};
