import { AssignmentIndOutlined, Chat, Home, Mood } from '@mui/icons-material';
import Link from 'next/link';
import * as React from 'react';

export default function Layout({ children }: { children: React.ReactNode }) {
  // Put Header or Footer Here
  return (
    <div className='flex bg-primary-100'>
      <aside className='flex h-screen flex-col rounded-tr-[30px] bg-gradient-to-r from-primary-400 to-primary-500 md:w-[200px]'>
        <h1 className='py-4 px-6'>Logo</h1>
        <nav className='mt-4'>
          <ul className='p-6'>
            <li className='mb-4'>
              <Link href='/dashboard'>
                <span className='flex items-center'>
                  <Home className='mr-2' />
                  <p>Home</p>
                </span>
              </Link>
            </li>

            <li className='mb-4'>
              <Link href='/dashboard/allHabits'>
                <span className='flex items-center'>
                  <AssignmentIndOutlined className='mr-2' />
                  <p>Habits</p>
                </span>
              </Link>
            </li>

            <li className='mb-4'>
              <Link href='/dashboard/moods'>
                <span className='flex items-center'>
                  <Mood className='mr-2' />
                  <p>Moods</p>
                </span>
              </Link>
            </li>

            <li>
              <Link href='/dashboard/conversations'>
                <span className='flex items-center'>
                  <Chat className='mr-2' />
                  <p>Messages</p>
                </span>
              </Link>
            </li>
          </ul>
        </nav>
      </aside>

      <div className='flex-grow pl-4'>
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
