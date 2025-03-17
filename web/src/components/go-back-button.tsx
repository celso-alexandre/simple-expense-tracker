import { Button } from 'antd';
import { useNavigate } from 'react-router-dom';

export function GoBackButton() {
  const navigate = useNavigate();

  return (
    <Button
      type='text'
      onClick={() => navigate(-1)}
      className="text-black hover:text-gray-500 w-[2px]"
    >
      <span className='text-xs'>{'<'}</span>
    </Button>
  );
}