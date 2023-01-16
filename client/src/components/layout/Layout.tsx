import { AssignmentIndOutlined, Chat, Home, Mood } from '@mui/icons-material';
import Link from 'next/link';
import * as React from 'react';

export default function Layout({ children }: { children: React.ReactNode }) {
  // Put Header or Footer Here
  return (
    <div className='flex bg-primary-50'>
      <aside className='flex h-screen flex-col rounded-tr-[60px] bg-gradient-to-r from-primary-100 to-primary-200 md:w-[250px]'>
        <h1 className='py-4 px-6'>Logo</h1>
        <nav className='mt-4'>
          <ul className='p-6'>
            <li className='mb-4'>
              <Link href='/dashboard' className='group'>
                <span className='flex items-center rounded px-4 py-2 group-hover:bg-primary-300'>
                  <Home className='mr-2 fill-primary-900' />
                  <p className='text-primary-900'>Home</p>
                </span>
              </Link>
            </li>

            <li className='mb-4'>
              <Link href='/dashboard/allHabits' className='group'>
                <span className='flex items-center rounded px-4 py-2 group-hover:bg-primary-300'>
                  <AssignmentIndOutlined className='mr-2 fill-primary-900' />
                  <p className='text-primary-900'>Habits</p>
                </span>
              </Link>
            </li>

            <li className='mb-4'>
              <Link href='/dashboard/moods' className='group'>
                <span className='flex items-center rounded px-4 py-2 group-hover:bg-primary-300'>
                  <Mood className='mr-2 fill-primary-900' />
                  <p className='text-primary-900'>Moods</p>
                </span>
              </Link>
            </li>

            <li>
              <Link href='/dashboard/conversations' className='group'>
                <span className='flex items-center rounded px-4 py-2 group-hover:bg-primary-300'>
                  <Chat className='mr-2 fill-primary-900' />
                  <p className='text-primary-900'>Conversations</p>
                </span>
              </Link>
            </li>

            {/*  create underline component with tailwind*/}

            {/*<hr className='my-4 border border-primary-500' />*/}
          </ul>
        </nav>
      </aside>

      <div className='flex-grow pl-12'>
        <div className='flex h-16 w-full items-center justify-between'>
          <p className='flex'>
            <b className='mr-1'>Hello,</b> user name
          </p>

          <p className='pr-6'>User icon</p>
        </div>

        {children}
      </div>
    </div>
  );
}
