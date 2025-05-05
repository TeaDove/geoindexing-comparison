import React, { useState, useEffect } from 'react';
import './Notification.css';

export interface NotificationMessage {
    status: number;
    endpoint: string;
    method: string;
    error?: string;
    timestamp: number;
    durationMs?: number;
}

interface NotificationProps {
    message: NotificationMessage | null;
}

const Notification: React.FC<NotificationProps> = ({ message }) => {
    const [isVisible, setIsVisible] = useState(false);

    useEffect(() => {
        if (message) {
            setIsVisible(true);
            const timer = setTimeout(() => {
                setIsVisible(false);
            }, 5000); // Hide after 5 seconds
            return () => clearTimeout(timer);
        }
    }, [message]);

    if (!message || !isVisible) {
        return null;
    }

    const isError = message.status >= 400;
    const formattedTime = new Date(message.timestamp).toLocaleTimeString();
    const durationText = message.durationMs ? ` (${message.durationMs.toFixed(1)} ms)` : ''; // Format duration

    return (
        <div className={`notification ${isError ? 'error' : 'success'} ${isVisible ? 'visible' : ''}`}>
            <span className="timestamp">[{formattedTime}]</span>
            <span className="method">{message.method}</span>
            <span className="endpoint">{message.endpoint}</span>
            <span className="status"> -&gt; {message.status}</span>
            <span className="duration">{durationText}</span>
            {message.error && <span className="error-message">: {message.error}</span>}
        </div>
    );
};

export default Notification;