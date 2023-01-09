import * as React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import Button from '@/components/buttons/Button';

type FormValues = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  role: string;
  street: string;
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
      className='mx-8 flex h-[700px] w-96 flex-grow flex-col rounded-3xl bg-zinc-100 py-16 px-10 drop-shadow-2xl'
    >
      <div className='mb-12'>
        <h1 className='mb-4'>Create an account</h1>
        <p className='text-neutral-500'>Start tracking your habits!</p>
      </div>

      <h2 className='mb-6'>Basic information</h2>
      <div className='mb-4 flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label className='mb-1' htmlFor='firstName'>
            First name
          </label>
          <input {...register('firstName')} className='form-input rounded-lg' />
        </div>

        <div className='ml-4 flex w-full flex-col'>
          <label className='mb-1' htmlFor='lastName'>
            Last name
          </label>
          <input {...register('lastName')} className='form-input rounded-lg' />
        </div>
      </div>

      <div className='mb-4 flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label className='mb-1' htmlFor='email'>
            Email
          </label>
          <input
            type='email'
            {...register('email')}
            className='form-input rounded-lg'
          />
        </div>

        <div className='ml-4 flex w-full flex-col'>
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

      <div className='mb-4 flex flex-col'>
        <label className='mb-1' htmlFor='role'>
          Role
        </label>
        <select className='form-input rounded-lg' {...register('role')}>
          <option value='REGULAR'>Regular</option>
          <option value='TRAINER'>Trainer</option>
        </select>
      </div>

      <h2>Additional information</h2>

      <div className='flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label className='mb-1' htmlFor='street'>
            Street
          </label>
          <input
            type='street'
            {...register('street')}
            className='form-input rounded-lg'
          />
        </div>

        <div className='ml-4 flex w-full flex-col'>
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

      <label className='mb-1' htmlFor='country'>
        Country
      </label>
      <input
        type='country'
        {...register('country')}
        className='form-input rounded-lg'
      />

      <div className='flex justify-end'>
        <Button type='submit' className='mt-4'>
          Sign up
        </Button>
      </div>
    </form>
  );
};

export default SignUpForm;
