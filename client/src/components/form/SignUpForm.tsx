import { zodResolver } from '@hookform/resolvers/zod';
import * as React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { z } from 'zod';

import { useGraphQLClient } from '@/hooks';

import Button from '@/components/buttons/Button';
import { SignUpFormSchema } from '@/components/form/validation';
import UnderlineLink from '@/components/links/UnderlineLink';

import { Role, useCreateUser } from '@/graphql';

type FormValuesSchema = z.infer<typeof SignUpFormSchema>;

const SignUpForm = () => {
  const {
    handleSubmit,
    register,
    reset,
    formState: { isSubmitting, errors },
  } = useForm<FormValuesSchema>({
    resolver: zodResolver(SignUpFormSchema),
  });
  const { graphQLClient } = useGraphQLClient();

  const { mutate, data } = useCreateUser(graphQLClient);

  console.log(data);

  const handleLogin: SubmitHandler<FormValuesSchema> = (data) => {
    // TODO: Handle login
    console.log(data);
    mutate({
      input: {
        firstName: data.firstName,
        lastName: data.lastName,
        email: data.email,
        password: data.password,
        role: data.role,
        address: [
          {
            city: data.address[0].city,
            country: data.address[0].country,
            street: '123 Main St',
          },
        ],
      },
    });

    reset();
  };

  return (
    <form
      onSubmit={handleSubmit(handleLogin)}
      className='flex w-full flex-col rounded-t-3xl px-10 py-8 md:my-10 md:mx-auto md:w-[800px] md:flex-grow-0 md:rounded-3xl md:bg-gray-50 md:py-10 md:drop-shadow-2xl xl:w-[600px]'
    >
      <div className='mb-8 text-center'>
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

          {errors.firstName && (
            <p className='mt-1 text-sm text-red-500'>
              {errors.firstName.message}
            </p>
          )}
        </div>

        <div className='flex w-full flex-col md:ml-4'>
          <label className='mb-1' htmlFor='lastName'>
            Last name
          </label>
          <input {...register('lastName')} className='form-input rounded-lg' />
          {errors.lastName && (
            <p className='mt-1 text-sm text-red-500'>
              {errors.lastName.message}
            </p>
          )}
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
          {errors.email && (
            <p className='mt-1 text-sm text-red-500'>{errors.email.message}</p>
          )}
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
          {errors.password && (
            <p className='mt-1 text-sm text-red-500'>
              {errors.password.message}
            </p>
          )}
        </div>
      </div>

      <div className='mb-6 flex flex-col'>
        <label className='mb-1' htmlFor='role'>
          Role
        </label>
        <select className='form-input rounded-lg' {...register('role')}>
          <option value={Role.Regular}>Regular</option>
          <option value={Role.Trainer}>Trainer</option>
        </select>
        {errors.role && (
          <p className='mt-1 text-sm text-red-500'>{errors.role.message}</p>
        )}
      </div>

      <div className='mb-4 flex flex-col md:flex-row'>
        <div className='mb-4 flex w-full flex-col md:mb-0 md:mr-4'>
          <label className='mb-1' htmlFor='street'>
            Country
          </label>
          <input
            {...register('address.0.country')}
            className='form-input rounded-lg'
          />
          {errors.address?.[0]?.country && (
            <p className='mt-1 text-sm text-red-500'>
              {errors.address[0].country.message}
            </p>
          )}
        </div>

        <div className='flex w-full flex-col md:ml-4'>
          <label className='mb-1' htmlFor='city'>
            City
          </label>
          <input
            type='city'
            {...register('address.0.city')}
            className='form-input rounded-lg'
          />

          {errors.address?.[0]?.city && (
            <p className='mt-1 text-sm text-red-500'>
              {errors.address[0].city.message}
            </p>
          )}
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
        <Button
          type='submit'
          className='mt-4 rounded-lg'
          disabled={isSubmitting}
        >
          Create Account
        </Button>
      </div>
    </form>
  );
};

export default SignUpForm;
