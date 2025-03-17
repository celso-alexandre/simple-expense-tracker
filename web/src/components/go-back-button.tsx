import { Button } from 'antd';
import { useNavigate } from 'react-router-dom';

export function GoBackButton() {
  const navigate = useNavigate();

  return (
    <Button
      type='text'
      onClick={() => navigate(-1)}
      className="text-black px-4 py-2 rounded-lg hover:text-gray-500"
    >
      {'<'}
    </Button>
  );
}