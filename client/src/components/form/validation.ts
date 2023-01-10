import { z } from 'zod';

enum Role {
  Regular = 'REGULAR',
  Trainer = 'TRAINER',
}

export const SignUpFormSchema = z.object({
  firstName: z
    .string({
      required_error: 'First name is required',
      invalid_type_error: 'First name must be a string',
    })
    .min(1, { message: 'First name is required' })
    .max(255, { message: 'First name can be atmost 255 characters long' }),
  lastName: z
    .string({
      required_error: 'Last name is required',
      invalid_type_error: 'Last name must be a string',
    })
    .min(1, { message: 'Last name is required' })
    .max(255, { message: 'Last name can be atmost 255 characters long' }),
  email: z
    .string({
      required_error: 'Email is required',
      invalid_type_error: 'Email must be a string',
    })
    .email({ message: 'Email is required' })
    .min(5, { message: 'Email must be atleast 5 characters long' }),
  password: z
    .string({
      required_error: 'Password is required',
      invalid_type_error: 'Password must be a string',
    })
    .min(8, { message: 'Password must be atleast 8 characters long' })
    .max(255, {
      message: 'Password can be atmost 255 characters long',
    }),
  role: z.nativeEnum(Role),
  address: z.tuple([
    z.object({
      city: z
        .string({
          required_error: 'City is required',
          invalid_type_error: 'City must be a string',
        })
        .min(1, { message: 'City is required' }),
      country: z
        .string({
          required_error: 'Country is required',
          invalid_type_error: 'Country must be a string',
        })
        .min(1, { message: 'Country is required' }),
      street: z
        .string({
          required_error: 'Street is required',
          invalid_type_error: 'Street must be a string',
        })
        .min(1, { message: 'Street is required' }),
    }),
  ]),
});
