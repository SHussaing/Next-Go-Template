'use client';  // This directive marks the component as a client component

import React, { useState, useEffect } from 'react';

export default function HelloPage() {
  const [message, setMessage] = useState<string>('');

  // Fetch the message directly when the component is mounted
  useEffect(() => {
    const fetchMessage = async () => {
      const res = await fetch(`${process.env.NEXT_PUBLIC_GO_API}/hello`);
      const data = await res.json();
      setMessage(data.message);
    };

    fetchMessage();
  }, []);

  return (
    <div>
      <h1>API Response</h1>
      <p>{message || 'Loading...'}</p>
    </div>
  );
}
