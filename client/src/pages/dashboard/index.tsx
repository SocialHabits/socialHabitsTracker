import Calendar from '@/components/Calendar';
import PrimaryLink from '@/components/links/PrimaryLink';

const Dashboard = () => {
  return (
    <div className='flex'>
      <div className='w-2/3 pr-8'>
        <div className='flex flex-grow items-center justify-between'>
          <h2 className='text-lg font-medium capitalize'>Find your activity</h2>

          <PrimaryLink href='/dashboard/allHabits'>View More</PrimaryLink>
        </div>
      </div>

      <div className='w-1/3 pl-6 pr-8'>
        <div>
          <h2 className='text-lg font-medium capitalize'>Calendar</h2>
          <div>
            <Calendar />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
