import { SubmitHandler, useForm } from 'react-hook-form';

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
    <form onSubmit={handleSubmit(handleLogin)} className='flex flex-col px-8'>
      <h2>Basic information</h2>
      <div className='flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label htmlFor='firstName'>First name</label>
          <input {...register('firstName')} className='form-input rounded-lg' />
        </div>

        <div className='ml-4 flex w-full flex-col'>
          <label htmlFor='lastName'>Last name</label>
          <input {...register('lastName')} className='form-input rounded-lg' />
        </div>
      </div>

      <div className='flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label htmlFor='email'>Email</label>
          <input
            type='email'
            {...register('email')}
            className='form-input rounded-lg'
          />
        </div>

        <div className='ml-4 flex w-full flex-col'>
          <label htmlFor='password'>Password</label>
          <input
            type='password'
            {...register('password')}
            className='form-input rounded-lg'
          />
        </div>
      </div>

      <label htmlFor='role'>Role</label>
      <select className='form-input rounded-lg' {...register('role')}>
        <option value='REGULAR'>Regular</option>
        <option value='TRAINER'>Trainer</option>
      </select>

      <h2>Additional information</h2>

      <div className='flex'>
        <div className='mr-4 flex w-full flex-col'>
          <label htmlFor='street'>Street</label>
          <input
            type='street'
            {...register('street')}
            className='form-input rounded-lg'
          />
        </div>

        <div className='ml-4 flex w-full flex-col'>
          <label htmlFor='city'>City</label>
          <input
            type='city'
            {...register('city')}
            className='form-input rounded-lg'
          />
        </div>
      </div>

      <label htmlFor='country'>Country</label>
      <input
        type='country'
        {...register('country')}
        className='form-input rounded-lg'
      />

      <button type='submit' className='btn btn-primary mt-4'>
        Sign up
      </button>
    </form>
  );
};

export default SignUpForm;
