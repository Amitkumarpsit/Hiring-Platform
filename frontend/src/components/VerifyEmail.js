import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { verifyEmail } from '../api/api';  // Import from api.js

function VerifyEmail() {
  const [message, setMessage] = useState('');
  const location = useLocation();

  useEffect(() => {
    const verifyToken = async () => {
      const params = new URLSearchParams(location.search);
      const token = params.get('token');

      if (!token) {
        setMessage('Invalid verification token.');
        return;
      }

      try {
        const response = await verifyEmail(token);  // Use the imported function
        setMessage(response.data.message);
      } catch (error) {
        setMessage(error.response?.data?.message || 'Verification failed.');
      }
    };

    verifyToken();
  }, [location.search]);

  return (
    <div>
      <h2>Email Verification</h2>
      <p>{message}</p>
    </div>
  );
}

export default VerifyEmail;
