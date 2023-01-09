import * as React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

import Button from '@/components/buttons/Button';
import UnderlineLink from '@/components/links/UnderlineLink';

type FormValues = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  role: string;
  city: string;
  country: string;
};

const SignUpForm = () => {
  const { handleSubmit, register } = useForm<FormValues>();

  const handleLogin: SubmitHandler<FormValues> = (data) => {
    // TODO: Handle login
    console.log(data);
  };

  return (
    <form
      onSubmit={handleSubmit(handleLogin)}
      className='flex w-full flex-col rounded-t-3xl px-10 py-8 md:my-10 md:mx-auto md:w-[800px] md:flex-grow-0 md:rounded-3xl md:bg-gray-50 md:py-10 md:drop-shadow-2xl xl:w-[600px]'
    >
      <div className='mb-8'>
        <h1 className='mb-3 font-medium'>Create an account</h1>
        <p className='text-neutral-500'>
          Start tracking your habits by creating an account!
        </p>
      </div>

      <div className='mb-4 flex flex-col md:flex-row'>
        <div className='mb-4 flex w-full flex-col md:mr-4 md:mb-0'>
          <label className='mb-1' htmlFor='firstName'>
            First name
          </label>
          <input {...register('firstName')} className='form-input rounded-lg' />
        </div>

        <div className='flex w-full flex-col md:ml-4'>
          <label className='mb-1' htmlFor='lastName'>
            Last name
          </label>
          <input {...register('lastName')} className='form-input rounded-lg' />
        </div>
      </div>

      <div className='mb-4 flex flex-col md:flex-row'>
        <div className='mb-4 flex w-full flex-col md:mb-0 md:mr-4'>
          <label className='mb-1' htmlFor='email'>
            Email
          </label>
          <input
            type='email'
            {...register('email')}
            className='form-input rounded-lg'
          />
        </div>

        <div className='flex w-full flex-col md:ml-4'>
          <label className='mb-1' htmlFor='password'>
            Password
          </label>
          <input
            type='password'
            {...register('password')}
            className='form-input rounded-lg'
          />
        </div>
      </div>

      <div className='mb-6 flex flex-col'>
        <label className='mb-1' htmlFor='role'>
          Role
        </label>
        <select className='form-input rounded-lg' {...register('role')}>
          <option value='REGULAR'>Regular</option>
          <option value='TRAINER'>Trainer</option>
        </select>
      </div>

      <div className='mb-4 flex flex-col md:flex-row'>
        <div className='mb-4 flex w-full flex-col md:mb-0 md:mr-4'>
          <label className='mb-1' htmlFor='street'>
            Country
          </label>
          <input {...register('country')} className='form-input rounded-lg' />
        </div>

        <div className='flex w-full flex-col md:ml-4'>
          <label className='mb-1' htmlFor='city'>
            City
          </label>
          <input
            type='city'
            {...register('city')}
            className='form-input rounded-lg'
          />
        </div>
      </div>

      <p>
        <span className='text-neutral-500'>
          Already have an account? Log in{' '}
        </span>
        <UnderlineLink href='/auth/login' className='text-primary-500'>
          here
        </UnderlineLink>
      </p>

      <div className='flex justify-end'>
        <Button type='submit' className='mt-4 rounded-lg'>
          Create Account
        </Button>
      </div>
    </form>
  );
};

export default SignUpForm;
