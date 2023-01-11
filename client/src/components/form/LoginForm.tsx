import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/router';
import * as React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { z } from 'zod';

import { useGraphQLClient } from '@/hooks';

import Button from '@/components/buttons/Button';
import { LoginFormSchema } from '@/components/form/validation';
import UnderlineLink from '@/components/links/UnderlineLink';

import { useLogin } from '@/graphql';

type FormValuesSchema = z.infer<typeof LoginFormSchema>;

const LoginForm = () => {
  const router = useRouter();
  const {
    handleSubmit,
    register,
    reset,
    formState: { isSubmitting, errors },
  } = useForm<FormValuesSchema>({
    resolver: zodResolver(LoginFormSchema),
  });
  const { graphQLClient } = useGraphQLClient();

  const { mutate, data } = useLogin(graphQLClient, {
    onSuccess: () => {
      router.push('/dashboard');
    },
  });

  const handleLogin: SubmitHandler<FormValuesSchema> = (data) => {
    mutate({
      email: data.email,
      password: data.password,
    });

    reset();
  };

  return (
    <form
      onSubmit={handleSubmit(handleLogin)}
      className='flex w-full flex-col rounded-t-3xl px-10 py-8 md:my-10 md:mx-auto md:w-[800px] md:flex-grow-0 md:rounded-3xl md:bg-gray-50 md:py-10 md:drop-shadow-2xl xl:w-[600px]'
    >
      <div className='mb-8 text-center'>
        <h1 className='mb-3 font-medium'>Welcome back!</h1>
        <p className='text-neutral-500'>Continue tracking your habits!</p>
      </div>

      <div className='mb-4 flex flex-col'>
        <div className='mb-4 flex w-full flex-col'>
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

        <div className='flex w-full flex-col'>
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

      <p>
        <span className='text-neutral-500'>
          Don't have an account? Sign up{' '}
        </span>
        <UnderlineLink href='/auth/sign-up' className='text-primary-500'>
          here
        </UnderlineLink>
      </p>

      <div className='flex justify-end'>
        <Button
          type='submit'
          className='mt-4 rounded-lg'
          disabled={isSubmitting}
        >
          Log in
        </Button>
      </div>
    </form>
  );
};

export default LoginForm;
